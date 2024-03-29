package _go

import (
	"path/filepath"

	"github.com/grafana/grafana/pkg/codegen"
	"github.com/grafana/grok/internal/jen"
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

func JenniesForGo(targetGrafanaVersion string) jen.TargetJennies {
	tgt := jen.NewTargetJennies()

	tgt.Core.Append(
		codegen.LatestMajorsOrXJenny(filepath.Join(targetGrafanaVersion, "kinds", "core"), false, codegen.GoTypesJenny{}),
		//codegen.LatestMajorsOrXJenny(filepath.Join(targetGrafanaVersion, "kinds", "core"), false, TypedSchemaJenny{}), // not good enough: references stuff from grafana/grafana
	)

	tgt.Composable.Append(
		// oooonly need to inject the proper path interstitial to make this right
		jen.ComposableLatestMajorsOrXJenny(filepath.Join(targetGrafanaVersion, "kinds", "composable"), GoTypesJenny{}),
	)

	return tgt
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
