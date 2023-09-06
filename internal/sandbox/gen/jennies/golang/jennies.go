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
	)
	targets.AppendOneToMany(
		codejen.AdaptOneToMany[[]ast.Builder, []*ast.File](
			&GoBuilder{},
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
