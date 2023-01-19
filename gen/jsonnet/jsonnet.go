package jsonnet

import (
	"fmt"
	"os"
	"strings"

	"github.com/grafana/codejen"
	"github.com/grafana/grafana/pkg/codegen"
	"github.com/grafana/grafana/pkg/kindsys"
	"github.com/grafana/grok/internal/jen"
)

func JenniesForJsonnet() jen.TargetJennies {
	tgt := jen.NewTargetJennies()

	tgt.Core.Append(
		&JsonnetImportsCoreJenny{},
		&JsonnetFileCoreJenny{},
	)

	tgt.Composable.Append(
		&JsonnetImportsComposableJenny{},
		&JsonnetFileComposableJenny{},
	)

	return tgt
}

type JsonnetImportsCoreJenny struct{}

func (j JsonnetImportsCoreJenny) JennyName() string {
	return "JsonnetImportsCoreJenny"
}

var tempGlobalImports strings.Builder

func (j *JsonnetImportsCoreJenny) Generate(decls ...*codegen.DeclForGen) (*codejen.File, error) {
	grafanaVersion := os.Getenv("GRAFANA_VERSION")
	if grafanaVersion == "" {
		return nil, nil
	}
	buf := strings.Builder{}
	buf.WriteString("[\n")
	for _, decl := range decls {
		var schemaVersion string
		if decl.Properties.Common().Maturity.Less(kindsys.MaturityStable) {
			schemaVersion = "x"
		} else {
			schemaVersion = decl.ForLatestSchema().Schema.Version().String()
		}
		importStatement := fmt.Sprintf("  import \"github.com/grafana/grok/jsonschema/kinds/core/%s/%s/%s_types_gen.json\",\n",
			decl.Lineage().Name(), schemaVersion, decl.Lineage().Name())
		buf.WriteString(importStatement)
		tempGlobalImports.WriteString(importStatement)
	}
	buf.WriteString("]\n")
	return nil, nil
}

type JsonnetImportsComposableJenny struct{}

func (j JsonnetImportsComposableJenny) JennyName() string {
	return "JsonnetImportsComposableJenny"
}

func (j *JsonnetImportsComposableJenny) Generate(comps ...*jen.ComposableForGen) (*codejen.File, error) {
	grafanaVersion := os.Getenv("GRAFANA_VERSION")
	if grafanaVersion == "" {
		return nil, nil
	}
	buf := strings.Builder{}
	buf.WriteString("[\n")
	buf.WriteString(tempGlobalImports.String())
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
		importStatement := fmt.Sprintf("  import \"github.com/grafana/grok/jsonschema/kinds/composable/%s/%s/%s_types_gen.json\",\n",
			comp.Lineage.Name()+"/"+strings.ToLower(comp.Slot.Name()), schemaVersion, comp.Lineage.Name())
		buf.WriteString(importStatement)
	}
	buf.WriteString("]\n")
	filename := fmt.Sprintf("%s/imports.libsonnet", grafanaVersion)
	return codejen.NewFile(filename, []byte(buf.String()), j), nil
}

const (
	jsonnetfilePre = `
{
	"version": 1,
	"dependencies": [
`
	jsonnetfilePost = `
	],
	"legacyImports": false
  }
`
	jsonnetfileEntry = `
	{
		"source": {
		  "git": {
			"remote": "https://github.com/grafana/grok.git",
			"subdir": "jsonschema/kinds/%s/%s/%s"
		  }
		},
		"version": "main"
	}`
)

type JsonnetFileCoreJenny struct{}

func (j JsonnetFileCoreJenny) JennyName() string {
	return "JsonnetFileCoreJenny"
}

var tempGlobalJsonnetFileEntries []string

func (j *JsonnetFileCoreJenny) Generate(decls ...*codegen.DeclForGen) (*codejen.File, error) {
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
		entry := fmt.Sprintf(jsonnetfileEntry, "core", decl.Lineage().Name(), schemaVersion)
		tempGlobalJsonnetFileEntries = append(tempGlobalJsonnetFileEntries, entry)
	}
	return nil, nil
}

type JsonnetFileComposableJenny struct{}

func (j JsonnetFileComposableJenny) JennyName() string {
	return "JsonnetFileComposableJenny"
}

func (j *JsonnetFileComposableJenny) Generate(comps ...*jen.ComposableForGen) (*codejen.File, error) {
	grafanaVersion := os.Getenv("GRAFANA_VERSION")
	if grafanaVersion == "" {
		return nil, nil
	}
	buf := strings.Builder{}
	buf.WriteString(jsonnetfilePre)
	entries := []string{}
	entries = append(entries, tempGlobalJsonnetFileEntries...)
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
		entry := fmt.Sprintf(jsonnetfileEntry, "composable", comp.Lineage.Name()+"/"+strings.ToLower(comp.Slot.Name()), schemaVersion)
		entries = append(entries, entry)
	}
	buf.WriteString(strings.Join(entries, ",\n"))
	buf.WriteString(jsonnetfilePost)
	filename := fmt.Sprintf("%s/jsonnetfile.json", grafanaVersion)
	return codejen.NewFile(filename, []byte(buf.String()), j), nil
}
