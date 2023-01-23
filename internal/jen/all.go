package jen

import (
	"bytes"
	"fmt"
	"path/filepath"

	"github.com/grafana/codejen"
	"github.com/grafana/grafana/pkg/codegen"
	"github.com/grafana/grafana/pkg/kindsys"
)

// TargetJennies is a set of jennies for a particular target language or
// tool that perform all necessary code generation steps.
type TargetJennies struct {
	Core       *codejen.JennyList[*codegen.DeclForGen]
	Composable *codejen.JennyList[kindsys.Composable]
}

// NewTargetJennies initializes a new TargetJennies with appropriate namers for
// each JennyList.
func NewTargetJennies() TargetJennies {
	return TargetJennies{
		Core: codejen.JennyListWithNamer[*codegen.DeclForGen](func(decl *codegen.DeclForGen) string {
			return decl.Properties.Common().MachineName
		}),
		Composable: codejen.JennyListWithNamer[kindsys.Composable](func(k kindsys.Composable) string {
			return k.Name()
		}),
	}
}

// Prefixer returns a FileMapper that injects the provided path prefix to files
// passed through it.
func Prefixer(prefix string) codejen.FileMapper {
	return func(f codejen.File) (codejen.File, error) {
		f.RelativePath = filepath.Join(prefix, f.RelativePath)
		return f, nil
	}
}

// SlashHeaderMapper produces a FileMapper that injects a comment header onto
// a [codejen.File] indicating the main generator that produced it (via the provided
// maingen, which should be a path) and the jenny or jennies that constructed the
// file.
func SlashHeaderMapper(maingen string) codejen.FileMapper {
	return func(f codejen.File) (codejen.File, error) {
		// Never inject on certain filetypes, it's never valid
		switch filepath.Ext(f.RelativePath) {
		case ".json", ".yml", ".yaml", ".md":
			return f, nil
		default:
			buf := new(bytes.Buffer)
			if err := tmpls.Lookup("gen_header.tmpl").Execute(buf, tvars_gen_header{
				MainGenerator: maingen,
				Using:         f.From,
			}); err != nil {
				return codejen.File{}, fmt.Errorf("failed executing gen header template: %w", err)
			}
			fmt.Fprint(buf, string(f.Data))
			f.Data = buf.Bytes()
		}
		return f, nil
	}
}

var gen_header = `// Code generated - EDITING IS FUTILE. DO NOT EDIT.
//
// Generated by pipeline:
//     {{ .MainGenerator }}
// Using jennies:
{{- range .Using }}
//     {{ .JennyName }}
{{- end }}
//
// Run 'go generate ./' from repository root to regenerate.

`
