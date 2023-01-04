package jsonschema

import (
	"encoding/json"
	"path/filepath"

	"cuelang.org/go/cue/cuecontext"
	"github.com/grafana/codejen"
	"github.com/grafana/grafana/pkg/codegen"
	"github.com/grafana/grok/internal/jen"
	"github.com/grafana/thema/encoding/jsonschema"
)

// Key things
// - create a dashboard, raw
// - create a dashboard, with a builder
// - take that object, translate to target grafana version

// go/kinds/core/<machineName>/<v>
// go/kinds/custom/<pluginName>/<machineName>
// go/kinds/composable/<pluginName>/<slot>/<v>
// go/byrelease/{c,c,c}
// go/builder

func JenniesForJsonSchema() jen.TargetJennies {
	tgt := jen.NewTargetJennies()

	tgt.Core.Append(
		codegen.LatestMajorsOrXJenny(filepath.Join("kinds", "core"), JsonSchemaJenny{}),
	)

	tgt.Composable.Append(
		// oooonly need to inject the proper path interstitial to make this right
		jen.ComposableLatestMajorsOrXJenny(filepath.Join("kinds", "composable"), JsonSchemaJenny{}),
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

/*
WITH IMPORT:
-table
stat E
-status-history
-state-timeline
-timeseries
piechart E
bargauge E
gauge E
histogram
barchart

WITHOUT IMPORT:
news
text
-heatmap
dashlist
-canvas
annolist
*/
