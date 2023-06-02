package jsonschema

import (
	"encoding/json"
	"os"
	"path/filepath"

	"cuelang.org/go/cue/cuecontext"
	"github.com/grafana/codejen"
	"github.com/grafana/grafana/pkg/codegen"
	"github.com/grafana/grok/internal/jen"
	"github.com/grafana/thema/encoding/jsonschema"
)

func JenniesForJsonSchema() jen.TargetJennies {
	tgt := jen.NewTargetJennies()
	grafanaVersion := os.Getenv("GRAFANA_VERSION")
	if grafanaVersion == "" {
		return tgt
	}

	tgt.Core.Append(
		codegen.LatestMajorsOrXJenny(filepath.Join(grafanaVersion, "kinds", "core"), JsonSchemaJenny{}),
	)

	tgt.Composable.Append(
		// oooonly need to inject the proper path interstitial to make this right
		jen.ComposableLatestMajorsOrXJenny(filepath.Join(grafanaVersion, "kinds", "composable"), JsonSchemaJenny{}),
	)

	return tgt
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

	return codejen.NewFile(sfg.Schema.Lineage().Name()+"_types_gen.json", []byte(str), j), nil
}
