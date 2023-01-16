package registry

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/grafana/codejen"
	"github.com/grafana/grafana/pkg/codegen"
	"github.com/grafana/grafana/pkg/kindsys"
	"github.com/grafana/grok/internal/jen"
)

func JenniesForRegistry() jen.TargetJennies {
	tgt := jen.NewTargetJennies()

	tgt.Core.Append(&CoreRegistryJenny{})
	tgt.Composable.Append(&ComposableRegistryJenny{})

	return tgt
}

type CoreRegistryJenny struct{}

func (j *CoreRegistryJenny) JennyName() string {
	return "CoreRegistryJenny"
}

func (j *CoreRegistryJenny) Generate(decls ...*codegen.DeclForGen) (*codejen.File, error) {

	type versionMapping struct {
		Kind   string `json:"kind"`
		Schema string `json:"version"`
	}

	versions := []versionMapping{}

	grafanaVersion := os.Getenv("GRAFANA_VERSION")
	if grafanaVersion == "" {
		return nil, nil
	}
	for _, decl := range decls {
		var schemaVersion string
		if decl.Properties.Common().Maturity.Less(kindsys.MaturityStable) {
			schemaVersion = "x"
		} else {
			schemaVersion = decl.ForLatestSchema().Schema.Version().String()
		}
		versions = append(versions, versionMapping{
			Kind:   decl.Lineage().Name(),
			Schema: schemaVersion,
		})
	}
	str, err := json.MarshalIndent(versions, "", "  ")
	if err != nil {
		return nil, err
	}
	filename := fmt.Sprintf("%s/registry-core.json", grafanaVersion)
	return codejen.NewFile(filename, []byte(str), j), nil
}

type ComposableRegistryJenny struct{}

func (j *ComposableRegistryJenny) JennyName() string {
	return "JsonSchemaIndexJenny"
}

func (j *ComposableRegistryJenny) Generate(comps ...*jen.ComposableForGen) (*codejen.File, error) {

	type versionMapping struct {
		Kind            string `json:"kind"`
		SchemaInterface string `json:"schemaInterface"`
		Schema          string `json:"version"`
	}

	versions := []versionMapping{}

	grafanaVersion := os.Getenv("GRAFANA_VERSION")
	if grafanaVersion == "" {
		return nil, nil
	}

	for _, comp := range comps {
		var schemaVersion string
		// Hardwiring maturity, as composables/plugins don't fully support kind system yet
		// see also internal/jen/jenny_eachmajorcomposable.go:
		maturity := kindsys.MaturityExperimental
		if maturity.Less(kindsys.MaturityStable) {
			schemaVersion = "x"
		} else {
			schemaVersion = comp.Lineage.Latest().Version().String()
		}
		versions = append(versions, versionMapping{
			Kind:            comp.Lineage.Name(),
			SchemaInterface: comp.Slot.Name(),
			Schema:          schemaVersion,
		})
	}
	str, err := json.MarshalIndent(versions, "", "  ")
	if err != nil {
		return nil, err
	}
	filename := fmt.Sprintf("%s/registry-composable.json", grafanaVersion)
	return codejen.NewFile(filename, []byte(str), j), nil
}
