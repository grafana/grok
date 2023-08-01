package jsonschema

import (
	"path/filepath"

	"github.com/grafana/grok/internal/jen"
	"github.com/grafana/kindsys/pkg/codegen"
)

func JenniesForJsonSchema(targetGrafanaVersion string) jen.TargetJennies {
	tgt := jen.NewTargetJennies()

	tgt.Core.Append(
		codegen.LatestMajorsOrXJenny(filepath.Join(targetGrafanaVersion, "kinds", "core"), false, codegen.JsonSchemaJenny{}),
	)

	tgt.Composable.Append(
		// oooonly need to inject the proper path interstitial to make this right
		codegen.ComposableLatestMajorsOrXJenny(filepath.Join(targetGrafanaVersion, "kinds", "composable"), codegen.JsonSchemaJenny{}),
	)

	return tgt
}
