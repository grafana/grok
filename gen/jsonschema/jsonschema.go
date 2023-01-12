package jsonschema

import (
	"encoding/json"
	"path/filepath"

	"cuelang.org/go/cue/cuecontext"
	"github.com/grafana/codejen"
	"github.com/grafana/grafana/pkg/codegen"
	"github.com/grafana/grok/internal/jen"
	"github.com/grafana/thema"
	"github.com/grafana/thema/encoding/jsonschema"
)

func JenniesForJsonSchema() jen.TargetJennies {
	tgt := jen.NewTargetJennies()

	tgt.Core.Append(
		codegen.LatestMajorsOrXJenny(filepath.Join("kinds", "core"), JsonSchemaJenny{}),
		&JsonSchemaCoreIndexJenny{},
	)

	tgt.Composable.Append(
		// oooonly need to inject the proper path interstitial to make this right
		jen.ComposableLatestMajorsOrXJenny(filepath.Join("kinds", "composable"), JsonSchemaJenny{}),
		&JsonSchemaComposableIndexJenny{},
	)

	return tgt
}

func jsonSchemaFilename(sch thema.Schema) string {
	return sch.Lineage().Name() + "_types_gen.json"
}

type JsonSchemaJenny struct{}

func (j JsonSchemaJenny) JennyName() string {
	return "JsonSchemaJenny"
}

func (j JsonSchemaJenny) Generate(sfg codegen.SchemaForGen) (*codejen.File, error) {
	// TODO allow using name instead of machine name in thema generator
	ast, err := jsonschema.GenerateSchema(sfg.Schema)
	if err != nil {
		return nil, err
	}
	ctx := cuecontext.New()
	str, err := json.MarshalIndent(ctx.BuildFile(ast), "", "  ")
	if err != nil {
		return nil, err
	}

	return codejen.NewFile(jsonSchemaFilename(sfg.Schema), []byte(str), j), nil
}

type JsonSchemaCoreIndexJenny struct{}

func (j *JsonSchemaCoreIndexJenny) JennyName() string {
	return "JsonSchemaIndexJenny"
}

func (j *JsonSchemaCoreIndexJenny) Generate(decls ...*codegen.DeclForGen) (*codejen.File, error) {

	type version struct {
		Name     string
		Filename string
	}
	type schema struct {
		Name     string
		Versions []version
	}

	schemas := []schema{}

	for _, decl := range decls {
		versions := []version{}
		for sch := decl.Lineage().First(); sch != nil; sch = sch.Successor() {
			versions = append(versions, version{
				Name:     sch.Version().String(),
				Filename: jsonSchemaFilename(sch),
			})
		}
		schemas = append(schemas, schema{
			Name:     decl.Lineage().Name(),
			Versions: versions,
		})
	}
	str, err := json.MarshalIndent(schemas, "", "  ")
	if err != nil {
		return nil, err
	}
	return codejen.NewFile("core.json", []byte(str), j), nil
}

type JsonSchemaComposableIndexJenny struct{}

func (j *JsonSchemaComposableIndexJenny) JennyName() string {
	return "JsonSchemaIndexJenny"
}

func (j *JsonSchemaComposableIndexJenny) Generate(comps ...*jen.ComposableForGen) (*codejen.File, error) {

	type version struct {
		Name     string
		Filename string
	}
	type schema struct {
		Name      string
		Interface string
		Versions  []version
	}

	schemas := []schema{}

	for _, comp := range comps {
		versions := []version{}
		for sch := comp.Lineage.First(); sch != nil; sch = sch.Successor() {
			versions = append(versions, version{
				Name:     sch.Version().String(),
				Filename: jsonSchemaFilename(sch),
			})
		}
		schemas = append(schemas, schema{
			Name:      comp.Lineage.Name(),
			Interface: comp.Slot.Name(),
			Versions:  versions,
		})
	}
	str, err := json.MarshalIndent(schemas, "", "  ")
	if err != nil {
		return nil, err
	}
	return codejen.NewFile("composables.json", []byte(str), j), nil
}
