package golang

import (
	"github.com/grafana/codejen"
	"github.com/grafana/grok/internal/sandbox/gen/ast"
	"github.com/grafana/grok/internal/sandbox/gen/jennies/tools"
	"github.com/grafana/grok/internal/sandbox/gen/veneers"
)

func Jennies() *codejen.JennyList[[]*ast.File] {
	targets := codejen.JennyListWithNamer[[]*ast.File](func(files []*ast.File) string {
		return "golang"
	})
	targets.AppendManyToMany(
		tools.Foreach[*ast.File](GoRawTypes{}),
		codejen.AdaptManyToMany[[]ast.Builder, []*ast.File](
			tools.Foreach[ast.Builder](&GoBuilder{}),
			func(files []*ast.File) []ast.Builder {
				generator := &ast.BuilderGenerator{}
				builders := generator.FromAST(files)

				return veneers.Engine().ApplyTo(builders)
			},
		),
	)
	targets.AddPostprocessors(PostProcessFile)

	return targets
}
