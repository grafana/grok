package jsonnet

import (
	"encoding/json"

	"cuelang.org/go/cue/cuecontext"
	"github.com/grafana/codejen"
	"github.com/grafana/kindsys/pkg/codegen"
	"github.com/grafana/thema/encoding/jsonschema"
)

type JsonnetSchemaJenny struct{}

func (j JsonnetSchemaJenny) JennyName() string {
	return "JsonnetImportsCoreJenny"
}

func (j JsonnetSchemaJenny) Generate(sfg codegen.SchemaForGen) (*codejen.File, error) {
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

	// @TODO we should be receiving a name without schema interface type so that we don't
	// need to strip it with a hack like this:
	name := fixKindName(sfg.Schema.Lineage().Name())
	return codejen.NewFile(name+"_types_gen.json", str, j), nil
}
