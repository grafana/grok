package jsonschema

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/grafana/codejen"
	"github.com/grafana/grafana/pkg/codegen"
	"github.com/grafana/grafana/pkg/kindsys"
	"github.com/grafana/grok/internal/jen"
)

type JsonSchemaCoreIndexJenny struct {
	parentdir string
}

func (j *JsonSchemaCoreIndexJenny) JennyName() string {
	return "JsonSchemaIndexJenny"
}

func (j *JsonSchemaCoreIndexJenny) Generate(decls ...*codegen.DeclForGen) (*codejen.File, error) {

	type kind map[string]string
	kinds := map[string]kind{}

	jsonSchemaFilename := func(parentdir string, version, name string) string {
		return fmt.Sprintf("%s/%s/%s/%s_types_gen.json", parentdir, name, version, name)
	}
	for _, decl := range decls {
		versions := map[string]string{}

		if decl.Properties.Common().Maturity.Less(kindsys.MaturityStable) {
			versions["x"] = jsonSchemaFilename(j.parentdir, "x", decl.Lineage().Name())
		} else {
			for sch := decl.Lineage().First(); sch != nil; sch = sch.Successor() {
				versions[sch.Version().String()] = jsonSchemaFilename(j.parentdir, sch.Version().String(), decl.Lineage().Name())
			}
		}
		kinds[decl.Lineage().Name()] = versions
	}
	str, err := json.MarshalIndent(kinds, "", "  ")
	if err != nil {
		return nil, err
	}
	return codejen.NewFile("core-index.json", []byte(str), j), nil
}

type JsonSchemaComposableIndexJenny struct {
	parentdir string
}

func (j *JsonSchemaComposableIndexJenny) JennyName() string {
	return "JsonSchemaIndexJenny"
}

func (j *JsonSchemaComposableIndexJenny) Generate(comps ...*jen.ComposableForGen) (*codejen.File, error) {

	type kind map[string]string
	kinds := map[string]kind{}

	jsonSchemaFilename := func(parentdir string, schemaInterface, version, name string) string {
		return fmt.Sprintf("%s/%s/%s/%s/%s_types_gen.json", parentdir, name, strings.ToLower(schemaInterface), version, name)
	}

	for _, comp := range comps {
		versions := map[string]string{}
		// Hardwiring maturity, as composables/plugins don't fully support kind system yet
		// see also internal/jen/jenny_eachmajorcomposable.go:
		maturity := kindsys.MaturityExperimental
		if maturity.Less(kindsys.MaturityStable) {
			versions["x"] = jsonSchemaFilename(j.parentdir, comp.Slot.Name(), "x", comp.Lineage.Name())
		} else {
			for sch := comp.Lineage.First(); sch != nil; sch = sch.Successor() {
				versions[sch.Version().String()] = jsonSchemaFilename(j.parentdir, comp.Slot.Name(), sch.Version().String(), comp.Lineage.Name())
			}
		}
		kinds[comp.Lineage.Name()] = versions
	}
	str, err := json.MarshalIndent(kinds, "", "  ")
	if err != nil {
		return nil, err
	}
	return codejen.NewFile("composable-index.json", []byte(str), j), nil
}
