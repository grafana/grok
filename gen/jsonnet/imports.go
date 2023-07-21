package jsonnet

import (
	"fmt"
	"strings"

	"github.com/grafana/codejen"
	"github.com/grafana/kindsys"
)

// Hack to make the Generate function re-entrant
var tempGlobalImportsBuf strings.Builder

type JsonnetCoreImportsJenny struct {
}

func (j JsonnetCoreImportsJenny) JennyName() string {
	return "JsonnetCoreImportsJenny"
}

func (j *JsonnetCoreImportsJenny) Generate(k ...kindsys.Kind) (*codejen.File, error) {
	tempGlobalImportsBuf.WriteString("[\n")
	for _, kind := range k {
		generateImports(kind)
	}
	return nil, nil
}

type JsonnetComposableImportsJenny struct {
	GrafanaVersion string
}

func (j JsonnetComposableImportsJenny) JennyName() string {
	return "JsonnetImportsCoreJenny"
}

func (j *JsonnetComposableImportsJenny) Generate(k ...kindsys.Composable) (*codejen.File, error) {
	for _, kind := range k {
		generateImports(kind)
	}
	tempGlobalImportsBuf.WriteString("]\n")
	filename := fmt.Sprintf("%s/imports.libsonnet", j.GrafanaVersion)
	return codejen.NewFile(filename, []byte(tempGlobalImportsBuf.String()), j), nil

	//return generateImports(k)
}

func generateImports(kind kindsys.Kind) {
	var schemaVersion string
	if kind.Maturity().Less(kindsys.MaturityStable) {
		schemaVersion = "x"
	} else {
		schemaVersion = kind.Lineage().Latest().Version().String()
	}
	// @TODO we should be receiving a name without schema interface type so that we don't
	// need to strip it with a hack like this:
	name := fixKindName(kind.MachineName())
	comp, isComposable := kind.(kindsys.Composable)
	var importStatement string
	if isComposable {
		schemaInterface := strings.ToLower(comp.Def().Properties.SchemaInterface)
		importStatement = fmt.Sprintf("  import \"./kinds/composable/%s/%s/%s/%s_types_gen.json\",\n",
			name, schemaInterface, schemaVersion, name)
	} else {
		importStatement = fmt.Sprintf("  import \"./kinds/core/%s/%s/%s_types_gen.json\",\n",
			name, schemaVersion, name)
	}
	tempGlobalImportsBuf.WriteString(importStatement)

}
