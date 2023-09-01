package golang

import (
	"github.com/grafana/codejen"
	"github.com/grafana/grok/internal/jen"
	"github.com/grafana/grok/internal/sandbox/gen/ast"
)

func Jennies(pkg string) *codejen.JennyList[*ast.File] {
	targets := codejen.JennyListWithNamer[*ast.File](func(f *ast.File) string {
		return f.Package
	})
	targets.AppendOneToOne(GoRawTypes{})
	targets.AppendOneToMany(&GoBuilder{})
	targets.AddPostprocessors(
		PostProcessFile,
		jen.Prefixer(pkg),
	)

	return targets
}
