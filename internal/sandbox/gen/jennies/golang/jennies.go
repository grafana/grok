package golang

import (
	"github.com/grafana/codejen"
	"github.com/grafana/grok/internal/sandbox/gen/ast"
	"github.com/grafana/grok/internal/sandbox/gen/jennies/tools"
)

func Jennies() *codejen.JennyList[[]*ast.File] {
	targets := codejen.JennyListWithNamer[[]*ast.File](func(files []*ast.File) string {
		return "golang"
	})
	targets.AppendManyToMany(
		tools.Foreach[*ast.File](GoRawTypes{}),
		tools.Foreach[*ast.File](&GoBuilder{}),
	)
	targets.AddPostprocessors(
		PostProcessFile,
		//jen.Prefixer(pkg),
	)

	return targets
}
