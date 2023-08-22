package jennies

import (
	"github.com/grafana/codejen"
	"github.com/grafana/grok/internal/jen"
	"github.com/grafana/grok/internal/sandbox/gen/ast"
	"github.com/grafana/grok/internal/sandbox/gen/jennies/golang"
	"github.com/grafana/grok/internal/sandbox/gen/jennies/typescript"
)

func All(outputPrefix string) *codejen.JennyList[*ast.File] {
	// FIXME: instead of mixing jennies for every languages, we should return a map
	// associating a language with the jennies for it
	generationTargets := codejen.JennyListWithNamer[*ast.File](func(f *ast.File) string {
		return f.Package
	})
	generationTargets.AppendOneToOne(
		// Golang
		golang.GoRawTypes{},

		// Typescript
		typescript.TypescriptRawTypes{},
	)
	generationTargets.AppendOneToMany(
		// Golang
		&golang.GoBuilder{},
	)
	generationTargets.AddPostprocessors(
		golang.PostProcessFile,
		jen.Prefixer(outputPrefix),
	)

	return generationTargets
}
