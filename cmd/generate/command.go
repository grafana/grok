package generate

import (
	"context"
	"fmt"
	"io/fs"
	"path/filepath"

	"cuelang.org/go/cue/cuecontext"
	"github.com/grafana/codejen"
	_go "github.com/grafana/grok/gen/go"
	"github.com/grafana/grok/gen/jsonnet"
	"github.com/grafana/grok/gen/jsonschema"
	"github.com/grafana/grok/internal/jen"
	"github.com/grafana/thema"
	"github.com/spf13/cobra"
)

type kindGenerator func(opts options, themaRuntime *thema.Runtime, commonFS fs.FS, targetJennies jen.TargetJennies) (*codejen.FS, error)

type options struct {
	kindRegistryRoot string
	targetVersion    string

	outputDir string

	// Do not run the generation process for these kinds
	excludeKinds []string
}

func (opts options) grafanaRegistryRoot() string {
	return filepath.Join(opts.kindRegistryRoot, "grafana")
}

func (opts options) corePath() string {
	return filepath.Join(opts.grafanaRegistryRoot(), opts.targetVersion, "core")
}

func (opts options) composablePath() string {
	return filepath.Join(opts.grafanaRegistryRoot(), opts.targetVersion, "composable")
}

func (opts options) commonLibPath() string {
	return filepath.Join(opts.grafanaRegistryRoot(), opts.targetVersion, "common")
}

func Command() *cobra.Command {
	opts := options{}

	cmd := &cobra.Command{
		Use:  "generate REGISTRY_PATH SCHEMA_VERSION",
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.kindRegistryRoot = args[0]
			opts.targetVersion = args[1]

			return doGenerate(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.outputDir, "output", "o", ".", "Output directory")
	cmd.Flags().StringSliceVar(&opts.excludeKinds, "exclude-kind", nil, "Excludes a kind from the code generation process.")

	return cmd
}

func doGenerate(opts options) error {
	themaRuntime := thema.NewRuntime(cuecontext.New())
	kindGenerators := map[string]kindGenerator{
		"core":       generateCoreKinds,
		"composable": generateComposableKinds,
	}

	fmt.Printf("Loading grafana-schema/common module from '%s'\n", opts.commonLibPath())
	commonCueImportPrefix := "cue.mod/pkg/github.com/grafana/grafana/packages/grafana-schema/src/common"
	commonFS, err := dirToPrefixedFS(opts.commonLibPath(), commonCueImportPrefix)
	if err != nil {
		return err
	}

	// Here begins the code generation setup
	rootCodeJenFS := codejen.NewFS()
	targetJennies := lineUpJennies("v" + opts.targetVersion)

	// Run the generation process for the desired kinds
	for kind, generator := range kindGenerators {
		if contains(opts.excludeKinds, kind) {
			continue
		}

		fmt.Printf("Building '%s' kinds from CUE files in '%s'\n", kind, opts.corePath())

		kindFS, err := generator(opts, themaRuntime, commonFS, targetJennies)
		if err != nil {
			return err
		}

		if err = rootCodeJenFS.Merge(kindFS); err != nil {
			return fmt.Errorf("could not merge kind FS into root FS: %w", err)
		}
	}

	// Write the output to disk
	if err = rootCodeJenFS.Write(context.Background(), opts.outputDir); err != nil {
		return fmt.Errorf("could not write rootCodeJenFS to disk: %w", err)
	}

	return nil
}

func generateCoreKinds(opts options, themaRuntime *thema.Runtime, _ fs.FS, targetJennies jen.TargetJennies) (*codejen.FS, error) {
	coreKinds, err := loadCoreKinds(themaRuntime, opts.corePath())
	if err != nil {
		return nil, err
	}

	fmt.Printf("Got %d core kinds\n", len(coreKinds))

	coreKindFS, err := targetJennies.Core.GenerateFS(coreKinds...)
	if err != nil {
		return nil, fmt.Errorf("could not generate FS for core kind: %w", err)
	}

	return coreKindFS, nil
}

func generateComposableKinds(opts options, themaRuntime *thema.Runtime, commonFS fs.FS, targetJennies jen.TargetJennies) (*codejen.FS, error) {
	composableKinds, err := loadComposableKinds(themaRuntime, commonFS, opts.composablePath())
	if err != nil {
		return nil, err
	}

	fmt.Printf("Got %d composable kinds\n", len(composableKinds))

	composableKindFS, err := targetJennies.Composable.GenerateFS(composableKinds...)
	if err != nil {
		return nil, fmt.Errorf("could not generate FS for composable kind: %w", err)
	}

	return composableKindFS, nil
}

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

func contains(haystack []string, needle string) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}

	return false
}
