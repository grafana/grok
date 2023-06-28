package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing/fstest"

	"cuelang.org/go/cue/cuecontext"
	"github.com/grafana/codejen"
	_go "github.com/grafana/grok/gen/go"
	"github.com/grafana/grok/gen/jsonnet"
	"github.com/grafana/grok/gen/jsonschema"
	"github.com/grafana/grok/internal/jen"
	"github.com/grafana/thema"
)

const kindRegistryRoot = "/home/kevin/sandbox/work/kind-registry/grafana"
const targetVersion = "10.0.0"

const outputRoot = "/home/kevin/sandbox/work/grok/output"

var corePath = filepath.Join(kindRegistryRoot, targetVersion, "core")
var composablePath = filepath.Join(kindRegistryRoot, targetVersion, "composable")
var commonLibPath = filepath.Join(kindRegistryRoot, targetVersion, "common")

// Line up all the jennies from all the language targets, prefixing them with
// their lang target subpaths.
func lineUpJennies(targetGrafanaVersion string) jen.TargetJennies {
	targets := jen.NewTargetJennies()

	targetMap := map[string]jen.TargetJennies{
		"go":         _go.JenniesForGo(targetGrafanaVersion), // This is not ready yet
		"jsonschema": jsonschema.JenniesForJsonSchema(targetGrafanaVersion),
		"jsonnet":    jsonnet.JenniesForJsonnet(targetGrafanaVersion),
	}

	for path, target := range targetMap {
		target.Core.AddPostprocessors(jen.Prefixer(path), jen.SlashHeaderMapper(path))
		target.Composable.AddPostprocessors(jen.Prefixer(path), jen.SlashHeaderMapper(path))

		targets.Core.AppendManyToMany(target.Core)
		targets.Composable.AppendManyToMany(target.Composable)
	}

	return targets
}

func main() {
	themaRuntime := thema.NewRuntime(cuecontext.New())

	commonHandle, err := os.ReadDir(commonLibPath)
	if err != nil {
		panic(err)
	}

	commonCueImportPrefix := "cue.mod/pkg/github.com/grafana/grafana/packages/grafana-schema/src/common"
	commonFS := fstest.MapFS{}
	for _, file := range commonHandle {
		if file.IsDir() {
			continue
		}

		content, err := os.ReadFile(filepath.Join(commonLibPath, file.Name()))
		if err != nil {
			panic(err)
		}

		commonFS[filepath.Join(commonCueImportPrefix, file.Name())] = &fstest.MapFile{
			Data: content,
		}
	}

	fmt.Printf("Building core Kinds from CUE files in '%s'\n", corePath)
	coreKinds, err := loadCoreKinds(themaRuntime, corePath)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Building composable Kinds from CUE files in '%s'\n", composablePath)
	composableKinds, err := loadComposableKinds(themaRuntime, commonFS, composablePath)
	if err != nil {
		panic(err)
	}

	// Here begins the code generation setup
	rootCodeJenFS := codejen.NewFS()
	targetJennies := lineUpJennies("v" + targetVersion)

	fmt.Printf("Got %d core kinds\n", len(coreKinds))
	fmt.Printf("Got %d composable kinds\n", len(composableKinds))

	coreKindFS, err := targetJennies.Core.GenerateFS(coreKinds...)
	if err != nil {
		panic(fmt.Errorf("could not generate FS for core kind: %w", err))
	}

	if err = rootCodeJenFS.Merge(coreKindFS); err != nil {
		panic(fmt.Errorf("could not merge coreKindFS into rootCodeJenFS: %w", err))
	}

	composableKindFS, err := targetJennies.Composable.GenerateFS(composableKinds...)
	if err != nil {
		panic(fmt.Errorf("could not generate FS for composable kind: %w", err))
	}

	if err = rootCodeJenFS.Merge(composableKindFS); err != nil {
		panic(fmt.Errorf("could not merge composableKindFS into rootCodeJenFS: %w", err))
	}

	if err = rootCodeJenFS.Write(context.Background(), outputRoot); err != nil {
		panic(fmt.Errorf("could not write rootCodeJenFS to disk: %w", err))
	}
}
