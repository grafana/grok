package jsonschema

import (
	"path/filepath"

	"github.com/grafana/grafana/pkg/codegen"
	"github.com/grafana/grok/internal/jen"
)

func JenniesForJsonSchema() jen.TargetJennies {
	tgt := jen.NewTargetJennies()

	tgt.Core.Append(
		codegen.LatestMajorsOrXJenny(filepath.Join("kinds", "core"), JsonSchemaJenny{}),
		&JsonSchemaCoreIndexJenny{
			parentdir: filepath.Join("kinds", "core"),
		},
	)

	tgt.Composable.Append(
		// oooonly need to inject the proper path interstitial to make this right
		jen.ComposableLatestMajorsOrXJenny(filepath.Join("kinds", "composable"), JsonSchemaJenny{}),
		&JsonSchemaComposableIndexJenny{
			parentdir: filepath.Join("kinds", "composable"),
		},
	)

	return tgt
}
