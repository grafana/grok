package generate

import (
	"context"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"cuelang.org/go/cue/cuecontext"
	"github.com/grafana/codejen"
	_go "github.com/grafana/grok/gen/go"
	"github.com/grafana/grok/gen/jsonnet"
	"github.com/grafana/grok/gen/jsonschema"
	"github.com/grafana/grok/internal/jen"
	"github.com/grafana/kindsys"
	"github.com/grafana/thema"
	"github.com/spf13/cobra"
	"github.com/yalue/merged_fs"
)

type kindGenerator func(opts options, themaRuntime *thema.Runtime, commonFS fs.FS, targetJennies jen.TargetJennies) (*codejen.FS, error)

type includeImport struct {
	fsPath     string // path of the library on the filesystem
	importPath string // path used in CUE files to import that library
}

type options struct {
	kindRegistryRoot string
	targetVersion    string
	minimumMaturity  string

	outputDir string

	// Do not run the generation process for these kinds
	excludeKinds []string

	excludeTargets []string

	imports []string
}

func (opts options) versionString() string {
	if opts.targetVersion == "next" {
		return "next"
	}

	return "v" + opts.targetVersion
}

func (opts options) maturity() kindsys.Maturity {
	switch opts.minimumMaturity {
	case "merged":
		return kindsys.MaturityMerged
	case "experimental":
		return kindsys.MaturityExperimental
	case "stable":
		return kindsys.MaturityStable
	case "mature":
		return kindsys.MaturityMature
	default:
		return kindsys.MaturityExperimental
	}
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

func (opts options) includeImports() ([]includeImport, error) {
	if len(opts.imports) == 0 {
		return nil, nil
	}

	imports := make([]includeImport, len(opts.imports))
	for i, importDefinition := range opts.imports {
		parts := strings.Split(importDefinition, ":")
		if len(parts) != 2 {
			return nil, fmt.Errorf("'%s' is not a valid import definition", importDefinition)
		}

		imports[i].fsPath = parts[0]
		imports[i].importPath = parts[1]
	}

	return imports, nil
}

func Command() *cobra.Command {
	opts := options{}

	cmd := &cobra.Command{
		Use:   "generate REGISTRY_PATH",
		Short: "Generates code from kinds stored in the registry",
		Long: `Generates code from kinds stored in the registry.

"core" and "composable" kinds defined in the registry can be generated as:

 * go
 * jsonnet
 * jsonschema

By default, each of these output targets is enabled.
`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.kindRegistryRoot = args[0]

			// If a version was specified
			if len(args) == 2 {
				opts.targetVersion = args[1]
			}

			return doGenerate(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.outputDir, "output", "o", ".", "Output directory.")
	cmd.Flags().StringVar(&opts.targetVersion, "version", "next", "Version to generate.")
	cmd.Flags().StringVar(&opts.minimumMaturity, "min-maturity", "experimental", "Minimum maturity of the kinds to generate.")
	cmd.Flags().StringSliceVar(&opts.excludeKinds, "exclude-kind", nil, "Excludes a kind from the code generation process.")
	cmd.Flags().StringSliceVar(&opts.excludeTargets, "exclude-target", nil, "Excludes a target from the code generation process.")
	cmd.Flags().StringArrayVarP(&opts.imports, "include-import", "I", nil, "Specify an additional library import directory. Format: [path]:[import]. Example: '../grafana/common-library:github.com/grafana/grafana/packages/grafana-schema/src/common")

	return cmd
}

func doGenerate(opts options) error {
	themaRuntime := thema.NewRuntime(cuecontext.New())
	kindGenerators := map[string]kindGenerator{
		"core":       generateCoreKinds,
		"composable": generateComposableKinds,
	}

	importDefinitions, err := opts.includeImports()
	if err != nil {
		return err
	}

	var librariesFS []fs.FS
	for _, importDefinition := range importDefinitions {
		fmt.Printf("Loading '%s' module from '%s'\n", importDefinition.importPath, importDefinition.fsPath)

		libraryFS, err := dirToPrefixedFS(importDefinition.fsPath, "cue.mod/pkg/"+importDefinition.importPath)
		if err != nil {
			return err
		}

		librariesFS = append(librariesFS, libraryFS)
	}

	commonFS := merged_fs.MergeMultiple(librariesFS...)

	// Here begins the code generation setup
	rootCodeJenFS := codejen.NewFS()
	targetJennies := lineUpJennies(opts.versionString(), opts.excludeTargets)

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

	coreKinds = filterByMaturity(coreKinds, opts.maturity())

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

	composableKinds = filterByMaturity(composableKinds, opts.maturity())

	fmt.Printf("Got %d composable kinds\n", len(composableKinds))

	composableKindFS, err := targetJennies.Composable.GenerateFS(composableKinds...)
	if err != nil {
		return nil, fmt.Errorf("could not generate FS for composable kind: %w", err)
	}

	return composableKindFS, nil
}

// Line up all the jennies from all the language targets, prefixing them with
// their lang target subpaths.
func lineUpJennies(targetVersion string, excludedTargets []string) jen.TargetJennies {
	targets := jen.NewTargetJennies()

	targetMap := map[string]jen.TargetJennies{
		"go":         _go.JenniesForGo(targetVersion), // This is not ready yet
		"jsonschema": jsonschema.JenniesForJsonSchema(targetVersion),
		"jsonnet":    jsonnet.JenniesForJsonnet(targetVersion),
	}

	for path, target := range targetMap {
		if contains(excludedTargets, path) {
			continue
		}

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

func filterByMaturity[T kindsys.Kind](kinds []T, minMaturity kindsys.Maturity) []T {
	var filtered []T

	for _, kind := range kinds {
		if kind.Maturity().Less(minMaturity) {
			continue
		}

		filtered = append(filtered, kind)
	}

	return filtered
}
