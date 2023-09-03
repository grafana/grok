package typescript

import (
	"github.com/grafana/codejen"
	"github.com/grafana/grok/internal/sandbox/gen/ast"
	"github.com/grafana/grok/internal/sandbox/gen/jennies/tools"
)

func Jennies() *codejen.JennyList[[]*ast.File] {
	targets := codejen.JennyListWithNamer[[]*ast.File](func(f []*ast.File) string {
		return "typescript"
	})
	targets.AppendOneToOne(
		TypescriptOptionsBuilder{},
	)
	targets.AppendManyToMany(
		tools.Foreach[*ast.File](TypescriptRawTypes{}),
		tools.Foreach[*ast.File](&TypescriptBuilder{}),
	)
	//targets.AddPostprocessors(jen.Prefixer(pkg))

	return targets
}
