package jsonnet

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/grafana/grafana/pkg/codegen"
	"github.com/grafana/grok/internal/jen"
)

func JenniesForJsonnet() jen.TargetJennies {
	tgt := jen.NewTargetJennies()

	grafanaVersion := os.Getenv("GRAFANA_VERSION")
	if grafanaVersion == "" {
		return tgt
	}
	tgt.Core.Append(
		&JsonnetCoreImportsJenny{},
		codegen.LatestMajorsOrXJenny(filepath.Join(grafanaVersion, "kinds", "core"), JsonnetSchemaJenny{}),
	)
	tgt.Composable.Append(
		// oooonly need to inject the proper path interstitial to make this right
		&JsonnetComposableImportsJenny{},
		jen.ComposableLatestMajorsOrXJenny(filepath.Join(grafanaVersion, "kinds", "composable"), JsonnetSchemaJenny{}),
	)

	return tgt
}

func fixKindName(name string) string {
	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, "panelcfg", "")
	name = strings.ReplaceAll(name, "dataquery", "")
	return name
}
