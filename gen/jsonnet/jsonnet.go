package jsonnet

import (
	"fmt"
	"os"
	"strings"

	"github.com/grafana/codejen"
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

func (j *JsonnetImportsCoreJenny) Generate(k ...kindsys.Kind) (*codejen.File, error) {
	grafanaVersion := os.Getenv("GRAFANA_VERSION")
	if grafanaVersion == "" {
		return nil, nil
	}
	buf := strings.Builder{}
	buf.WriteString("[\n")
	for _, kind := range k {
		var schemaVersion string
		if kind.Maturity().Less(kindsys.MaturityStable) {
			schemaVersion = "x"
		} else {
			schemaVersion = kind.Lineage().Latest().Version().String()
		}
		importStatement := fmt.Sprintf("  import \"github.com/grafana/grok/jsonschema/kinds/core/%s/%s/%s_types_gen.json\",\n",
			kind.Lineage().Name(), schemaVersion, kind.Lineage().Name())
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

func (j *JsonnetImportsComposableJenny) Generate(k ...kindsys.Composable) (*codejen.File, error) {
	grafanaVersion := os.Getenv("GRAFANA_VERSION")
	if grafanaVersion == "" {
		return nil, nil
	}
	buf := strings.Builder{}
	buf.WriteString("[\n")
	buf.WriteString(tempGlobalImports.String())
	for _, kind := range k {
		var schemaVersion string
		if kind.Maturity().Less(kindsys.MaturityStable) {
			schemaVersion = "x"
		} else {
			schemaVersion = kind.Lineage().Latest().Version().String()
		}
		importStatement := fmt.Sprintf("  import \"github.com/grafana/grok/jsonschema/kinds/composable/%s/%s/%s_types_gen.json\",\n",
			kind.Lineage().Name()+"/"+strings.ToLower(kind.Def().Properties.SchemaInterface), schemaVersion, kind.Lineage().Name())
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

func (j *JsonnetFileCoreJenny) Generate(k ...kindsys.Kind) (*codejen.File, error) {
	grafanaVersion := os.Getenv("GRAFANA_VERSION")
	if grafanaVersion == "" {
		return nil, nil
	}
	for _, kind := range k {
		var schemaVersion string
		if kind.Maturity().Less(kindsys.MaturityStable) {
			schemaVersion = "x"
		} else {
			schemaVersion = kind.Lineage().Latest().Version().String()
		}
		entry := fmt.Sprintf(jsonnetfileEntry, "core", kind.Lineage().Name(), schemaVersion)
		tempGlobalJsonnetFileEntries = append(tempGlobalJsonnetFileEntries, entry)
	}
	return nil, nil
}

type JsonnetFileComposableJenny struct{}

func (j JsonnetFileComposableJenny) JennyName() string {
	return "JsonnetFileComposableJenny"
}

func (j *JsonnetFileComposableJenny) Generate(k ...kindsys.Composable) (*codejen.File, error) {
	grafanaVersion := os.Getenv("GRAFANA_VERSION")
	if grafanaVersion == "" {
		return nil, nil
	}
	buf := strings.Builder{}
	buf.WriteString(jsonnetfilePre)
	entries := []string{}
	entries = append(entries, tempGlobalJsonnetFileEntries...)
	for _, kind := range k {
		var schemaVersion string
		if kind.Maturity().Less(kindsys.MaturityStable) {
			schemaVersion = "x"
		} else {
			schemaVersion = kind.Lineage().Latest().Version().String()
		}
		entry := fmt.Sprintf(jsonnetfileEntry, "composable", kind.Lineage().Name()+"/"+strings.ToLower(kind.Def().Properties.SchemaInterface), schemaVersion)
		entries = append(entries, entry)
	}
	buf.WriteString(strings.Join(entries, ",\n"))
	buf.WriteString(jsonnetfilePost)
	filename := fmt.Sprintf("%s/jsonnetfile.json", grafanaVersion)
	return codejen.NewFile(filename, []byte(buf.String()), j), nil
}
