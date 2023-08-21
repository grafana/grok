package main

import (
	"context"
	"os"

	"github.com/grafana/codejen"
	"github.com/grafana/grok/internal/jen"
	"github.com/grafana/grok/internal/sandbox/gen/ast"
	"github.com/grafana/grok/internal/sandbox/gen/jennies/golang"
	"github.com/grafana/grok/internal/sandbox/gen/jsonschema"
)

func main() {
	entrypoint := "./schemas/jsonschema/core/playlist/playlist.json"
	pkg := "playlist"

	reader, err := os.Open(entrypoint)
	if err != nil {
		panic(err)
	}

	schemaAst, err := jsonschema.GenerateAST(reader, jsonschema.Config{
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
		golang.GoRawTypes{},
		golang.GoBuilder{},
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
