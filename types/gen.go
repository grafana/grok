// go:build ignore
//go:build ignore
// +build ignore

package main

import (
	"bytes"
	"fmt"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/pkg/encoding/yaml"
	ocg "github.com/deepmap/oapi-codegen/pkg/codegen"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/grafana/grafana/pkg/codegen"
	"github.com/grafana/grafana/pkg/framework/coremodel"
	"github.com/grafana/grafana/pkg/framework/coremodel/registry"
	"github.com/grafana/thema"
	"github.com/grafana/thema/encoding/openapi"
	"golang.org/x/tools/imports"
)

var lib = thema.NewLibrary(cuecontext.New())

func main() {
	reg, err := registry.ProvideGeneric()
	if err != nil {
		// Unreachable in any actual released version of Grafana
		panic(err)
	}

	wd := codegen.NewWriteDiffer()
	for _, cm := range reg.List() {
		iwd, err := generateCoremodel(cm)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error while generating from %q coremodel: %s\n", cm.Lineage().Name(), err)
			os.Exit(1)
		}

		wd.Merge(iwd)
	}

	if _, set := os.LookupEnv("CODEGEN_VERIFY"); set {
		err = wd.Verify()
		if err != nil {
			fmt.Fprintf(os.Stderr, "generated code is not up to date:\n%s\nrun `go generate ./...` to regenerate\n\n", err)
			os.Exit(1)
		}
	} else {
		err = wd.Write()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error while writing generated code to disk:\n%s\n", err)
			os.Exit(1)
		}
	}
}

type tplVars struct {
	Name        string
	PackageName string
	TitleName   string
	Seqv, Schv  uint
	GenPath     string
	ModelFile   string
}

func schemaToTplVars(cm coremodel.Interface, sch thema.Schema) tplVars {
	l := cm.Lineage()
	vars := tplVars{
		Name:        l.Name(),
		PackageName: fmt.Sprintf("goschema%s", l.Name()),
		TitleName:   strings.Title(l.Name()),
		Seqv:        sch.Version()[0],
		Schv:        sch.Version()[1],
		ModelFile:   fmt.Sprintf("%s_model_gen.go", l.Name()),
	}

	pathv := fmt.Sprintf("v%v", vars.Seqv)
	if !isCanonical(cm) {
		pathv = "x"
	}

	vars.GenPath = filepath.Join(l.Name(), pathv, "goschema")
	return vars
}

// FIXME specifying coremodel canonicality DOES NOT belong here - it should be part of the coremodel declaration.
// this is just hoisting the problem forward from grafana/grafana itself
func isCanonical(cm coremodel.Interface) bool {
	return false
}

func generateCoremodel(cm coremodel.Interface) (codegen.WriteDiffer, error) {
	lin := cm.Lineage()
	wd := codegen.NewWriteDiffer()

	var seqv uint
	for {
		sv, err := thema.LatestVersionInSequence(cm.Lineage(), seqv)
		if err != nil {
			break
		}

		sch := thema.SchemaP(lin, sv)
		vars := schemaToTplVars(cm, sch)
		vstr := fmt.Sprintf("%s(v%v.%v)", lin.Name(), sv[0], sv[1])

		f, err := openapi.GenerateSchema(sch, nil)
		if err != nil {
			return nil, fmt.Errorf("%s: thema openapi generation failed: %w", vstr, err)
		}

		str, err := yaml.Marshal(lib.Context().BuildFile(f))
		if err != nil {
			return nil, fmt.Errorf("%s: cue-yaml marshaling failed: %w", vstr, err)
		}

		loader := openapi3.NewLoader()
		oT, err := loader.LoadFromData([]byte(str))
		if err != nil {
			return nil, fmt.Errorf("%s: loading generated openapi failed; %w", vstr, err)
		}

		gostr, err := ocg.Generate(oT, vars.PackageName, ocg.Options{
			GenerateTypes: true,
			SkipPrune:     true,
			SkipFmt:       true,
			UserTemplates: map[string]string{
				// "imports.tmpl": fmt.Sprintf(tmplImports, ls.RelativePath),
				// "typedef.tmpl": tmplTypedef,
			},
		})
		if err != nil {
			return nil, fmt.Errorf("%s: openapi generation failed: %w", vstr, err)
		}

		path := filepath.Join(vars.GenPath, vars.ModelFile)
		byt, err := postprocessGoFile([]byte(gostr), path)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", vstr, err)
		}
		wd[path] = byt

		// Generate the schema version-reporting file
		var buf bytes.Buffer
		err = tmplVersionFile.Execute(&buf, vars)
		if err != nil {
			return nil, fmt.Errorf("%s: version file gen failed: %w", vstr, err)
		}
		vpath := filepath.Join(vars.GenPath, "version_gen.go")
		byt, err = postprocessGoFile(buf.Bytes(), vpath)
		if err != nil {
			return nil, fmt.Errorf("%s, %w", vstr, err)
		}
		wd[vpath] = buf.Bytes()

		seqv++
	}

	return wd, nil
}

func postprocessGoFile(src []byte, path string) ([]byte, error) {
	fset := token.NewFileSet()
	gf, err := parser.ParseFile(fset, path, src, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("parsing generated %q failed: %w", path, err)
	}

	var buf bytes.Buffer
	err = format.Node(&buf, fset, gf)
	if err != nil {
		return nil, fmt.Errorf("ast printing failed: %w", err)
	}

	byt, err := imports.Process(path, buf.Bytes(), nil)
	if err != nil {
		return nil, fmt.Errorf("goimports postprocessing of %q failed: %w", path, err)
	}

	return byt, nil
}

var genHeader = `// This file is autogenerated. DO NOT EDIT.
//
// Run "go generate ./..." from the grok repository root to regenerate.
`

var tmplImports = genHeader + `package {{ .PackageName }}

import (
	"embed"
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/grafana/thema"
)
`

var tmplVersionFile = template.Must(template.New("versionfile").Parse(genHeader + `
package {{ .PackageName }}

import (
	"github.com/grafana/grafana/pkg/coremodel/{{ .Name }}"
	"github.com/grafana/grafana/pkg/framework/coremodel/registry"
	"github.com/grafana/thema"
)

// Version is the syntactic version of the schema from which the Go
// code in this package was generated.
//
// This value will always be the same as calling Schema().Version().
var Version thema.SyntacticVersion = [2]uint{ {{- .Seqv }}, {{ .Schv -}} }

// Schema returns the schema from which the Go code in this package was generated.
//
// This uses the central thema.Library and cue.Context provided in github.com/grafana/grafana/pkg/cuectx.
// If you must provide your own, call github.com/grafana/grafana/pkg/coremodel/{{ .Name }}.Lineage() directly.
func Schema() thema.Schema {
	reg, err := registry.ProvideStatic()
	if err != nil {
		panic(err)
	}
	return thema.SchemaP(reg.{{ .TitleName }}().Lineage(), thema.SV({{ .Seqv }}, {{ .Schv }}))
}
`))
