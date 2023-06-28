package generate

import (
	"context"
	"fmt"
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

type options struct {
	kindRegistryRoot string
	targetVersion    string

	outputDir string
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
		Use:   "generate REGISTRY_PATH SCHEMA_VERSION",
		Short: "TBD",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.kindRegistryRoot = args[0]
			opts.targetVersion = args[1]

			return doGenerate(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.outputDir, "output", "o", ".", "Output directory")

	return cmd
}

func doGenerate(opts options) error {
	themaRuntime := thema.NewRuntime(cuecontext.New())

	fmt.Printf("Loading grafana-schema/common module from '%s'\n", opts.commonLibPath())
	commonCueImportPrefix := "cue.mod/pkg/github.com/grafana/grafana/packages/grafana-schema/src/common"
	commonFS, err := dirToPrefixedFS(opts.commonLibPath(), commonCueImportPrefix)
	if err != nil {
		return err
	}

	fmt.Printf("Building core Kinds from CUE files in '%s'\n", opts.corePath())
	coreKinds, err := loadCoreKinds(themaRuntime, opts.corePath())
	if err != nil {
		return err
	}

	fmt.Printf("Building composable Kinds from CUE files in '%s'\n", opts.composablePath())
	composableKinds, err := loadComposableKinds(themaRuntime, commonFS, opts.composablePath())
	if err != nil {
		return err
	}

	// Here begins the code generation setup
	rootCodeJenFS := codejen.NewFS()
	targetJennies := lineUpJennies("v" + opts.targetVersion)

	fmt.Printf("Got %d core kinds\n", len(coreKinds))
	fmt.Printf("Got %d composable kinds\n", len(composableKinds))

	coreKindFS, err := targetJennies.Core.GenerateFS(coreKinds...)
	if err != nil {
		return fmt.Errorf("could not generate FS for core kind: %w", err)
	}

	if err = rootCodeJenFS.Merge(coreKindFS); err != nil {
		return fmt.Errorf("could not merge coreKindFS into rootCodeJenFS: %w", err)
	}

	composableKindFS, err := targetJennies.Composable.GenerateFS(composableKinds...)
	if err != nil {
		panic(fmt.Errorf("could not generate FS for composable kind: %w", err))
	}

	if err = rootCodeJenFS.Merge(composableKindFS); err != nil {
		return fmt.Errorf("could not merge composableKindFS into rootCodeJenFS: %w", err)
	}

	if err = rootCodeJenFS.Write(context.Background(), opts.outputDir); err != nil {
		return fmt.Errorf("could not write rootCodeJenFS to disk: %w", err)
	}

	return nil
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
