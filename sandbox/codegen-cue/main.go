package main

import (
	"context"
	"fmt"
	"path/filepath"

	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/load"
	"github.com/grafana/codejen"
	"github.com/grafana/grok/internal/sandbox/gen/ast"
	"github.com/grafana/grok/internal/sandbox/gen/jennies"
	"github.com/grafana/grok/internal/sandbox/gen/simplecue"
)

func main() {
	entrypoints := []string{
		"./schemas/cue/core/dashboard/",
		"./schemas/cue/core/playlist/",
	}

	allSchemas := make([]*ast.File, 0, len(entrypoints))
	for _, entrypoint := range entrypoints {
		pkg := filepath.Base(entrypoint)

		// Load Cue files into Cue build.Instances slice
		// the second arg is a configuration object, we'll see this later
		bis := load.Instances([]string{entrypoint}, nil)

		values, err := cuecontext.New().BuildInstances(bis)
		if err != nil {
			panic(err)
		}

		schemaAst, err := simplecue.GenerateAST(values[0], simplecue.Config{
			Package: pkg, // TODO: extract from input schema/?
		})
		if err != nil {
			panic(err)
		}

		allSchemas = append(allSchemas, schemaAst)
	}

	// Here begins the code generation setup
	targetsByLanguage := jennies.All()
	rootCodeJenFS := codejen.NewFS()

	for language, target := range targetsByLanguage {
		fmt.Printf("Running '%s' jennies...\n", language)

		var err error
		processedAsts := allSchemas

		for _, compilerPass := range target.CompilerPasses {
			processedAsts, err = compilerPass.Process(processedAsts)
			if err != nil {
				panic(err)
			}
		}

		fs, err := target.Jennies.GenerateFS(processedAsts)
		if err != nil {
			panic(err)
		}

		err = rootCodeJenFS.Merge(fs)
		if err != nil {
			panic(err)
		}
	}

	err := rootCodeJenFS.Write(context.Background(), "generated")
	if err != nil {
		panic(err)
	}
}
