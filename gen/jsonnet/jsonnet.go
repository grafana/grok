package jsonnet

import (
	"path/filepath"
	"strings"

	"github.com/grafana/grafana/pkg/codegen"
	"github.com/grafana/grok/internal/jen"
)

func JenniesForJsonnet(targetGrafanaVersion string) jen.TargetJennies {
	tgt := jen.NewTargetJennies()

	tgt.Core.Append(
		&JsonnetCoreImportsJenny{},
		codegen.LatestMajorsOrXJenny(filepath.Join(targetGrafanaVersion, "kinds", "core"), false, JsonnetSchemaJenny{}),
	)
	tgt.Composable.Append(
		// oooonly need to inject the proper path interstitial to make this right
		&JsonnetComposableImportsJenny{},
		jen.ComposableLatestMajorsOrXJenny(filepath.Join(targetGrafanaVersion, "kinds", "composable"), JsonnetSchemaJenny{}),
	)

	return tgt
}

func fixKindName(name string) string {
	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, "panelcfg", "")
	name = strings.ReplaceAll(name, "dataquery", "")
	return name
}
