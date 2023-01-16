package jsonschema

import (
	"encoding/json"
	"fmt"

	"cuelang.org/go/cue/cuecontext"
	"github.com/grafana/codejen"
	"github.com/grafana/grafana/pkg/codegen"
	"github.com/grafana/thema/encoding/jsonschema"
)

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

	filename := fmt.Sprintf("%s_types_gen.json", sfg.Schema.Lineage().Name())
	return codejen.NewFile(filename, []byte(str), j), nil
}
