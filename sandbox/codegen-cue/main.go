package main

import (
	"context"

	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/load"
	"github.com/grafana/codejen"
	"github.com/grafana/grok/internal/jen"
	"github.com/grafana/grok/internal/sandbox/gen/ast"
	"github.com/grafana/grok/internal/sandbox/gen/jennies/golang"
	"github.com/grafana/grok/internal/sandbox/gen/jennies/typescript"
	"github.com/grafana/grok/internal/sandbox/gen/simplecue"
)

func main() {
	entrypoints := []string{"./schemas/cue/core/dashboard/dashboard.cue"}
	pkg := "dashboard"

	//entrypoints := []string{"./schemas/cue/core/sandbox/sandbox.cue"}
	//pkg := "sandbox"

	//entrypoints := []string{"./schemas/cue/core/playlist/playlist.cue"}
	//pkg := "playlist"

	// Load Cue files into Cue build.Instances slice
	// the second arg is a configuration object, we'll see this later
	bis := load.Instances(entrypoints, nil)

	values, err := cuecontext.New().BuildInstances(bis)
	if err != nil {
		panic(err)
	}

	schemaAst, err := simplecue.GenerateAST(values[0], simplecue.Config{
		Package: pkg, // TODO: extract from input schema/folder?
	})
	if err != nil {
		panic(err)
	}

	// Here begins the code generation setup
	generationTargets := codejen.JennyListWithNamer[*ast.File](func(f *ast.File) string {
		return f.Package
	})
	generationTargets.AppendOneToOne(
		// Golang
		golang.GoRawTypes{},
		golang.GoBuilder{},

		// Typescript
		typescript.TypescriptRawTypes{},
	)
	generationTargets.AddPostprocessors(
		golang.PostProcessFile,
		jen.Prefixer(pkg),
	)

	rootCodeJenFS := codejen.NewFS()

	fs, err := generationTargets.GenerateFS(schemaAst)
	if err != nil {
		panic(err)
	}

	err = rootCodeJenFS.Merge(fs)
	if err != nil {
		panic(err)
	}

	err = rootCodeJenFS.Write(context.Background(), "newgen")
	if err != nil {
		panic(err)
	}
}
