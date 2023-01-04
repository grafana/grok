// //go:build ignore
// // +build ignore

//go:generate go run gen.go

package main

import (
	"context"
	"fmt"
	"os"

	"cuelang.org/go/cue/cuecontext"
	"github.com/grafana/codejen"
	"github.com/grafana/grafana/pkg/codegen"
	"github.com/grafana/grafana/pkg/plugins/pfs/corelist"
	"github.com/grafana/grafana/pkg/registry/corekind"
	_go "github.com/grafana/grok/gen/go"
	"github.com/grafana/grok/gen/jsonschema"
	"github.com/grafana/grok/internal/jen"
	"github.com/grafana/thema"
)

var rt = thema.NewRuntime(cuecontext.New())

func main() {
	jfs := codejen.NewFS()

	// i've got a lovely bunch of coconuts
	// there they are all standing in a row
	coco := lineUpJennies()

	var corek []*codegen.DeclForGen
	var compok []*jen.ComposableForGen
	for _, kind := range corekind.NewBase(nil).AllStructured() {
		corek = append(corek, codegen.StructuredForGen(kind))
	}
	for _, ptree := range corelist.New(nil) {
		compok = append(compok, jen.ComposablesFromTree(ptree)...)
	}

	ckfs, err := coco.Core.GenerateFS(corek...)
	die(err)
	die(jfs.Merge(ckfs))
	ckfs, err = coco.Composable.GenerateFS(compok...)
	die(err)
	die(jfs.Merge(ckfs))

	if _, set := os.LookupEnv("CODEGEN_VERIFY"); set {
		if err = jfs.Verify(context.Background(), ""); err != nil {
			die(fmt.Errorf("generated code is out of sync with inputs:\n%s\nrun `make gen-cue` to regenerate", err))
		}
	} else if err = jfs.Write(context.Background(), ""); err != nil {
		die(fmt.Errorf("error while writing generated code to disk:\n%s", err))
	}
}

// Line up all the jennies from all the language targets, prefixing them with
// their lang target subpaths.
func lineUpJennies() jen.TargetJennies {
	tgt := jen.NewTargetJennies()

	tgtmap := map[string]jen.TargetJennies{
		"go":         _go.JenniesForGo(),
		"jsonschema": jsonschema.JenniesForJsonSchema(),
	}

	for path, tj := range tgtmap {
		tj.Core.AddPostprocessors(jen.Prefixer(path))
		tj.Composable.AddPostprocessors(jen.Prefixer(path))
		tgt.Core.AppendManyToMany(tj.Core)
		tgt.Composable.AppendManyToMany(tj.Composable)
	}

	return tgt
}

func die(err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, err, "\n")
		os.Exit(1)
	}
}
