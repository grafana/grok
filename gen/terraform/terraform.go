package terraform

import (
	"fmt"
	"os"
	"path/filepath"

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

func JenniesForTerraform() jen.TargetJennies {
	tgt := jen.NewTargetJennies()

	grafanaVersion := os.Getenv("GRAFANA_VERSION")
	if grafanaVersion == "" {
		fmt.Println("missing GRAFANA_VERSION")
		return tgt
	}

	dataSourcePath := filepath.Join(grafanaVersion, "data_sources")
	tgt.Core.Append(
		jen.LatestJenny(dataSourcePath, TerraformDataSourceJenny{}),
		// codegen.LatestMajorsOrXJenny(filepath.Join("kinds", "core"), TypedSchemaJenny{}),
	)

	// tgt.Composable.Append(
	// 	jen.ComposableLatestMajorsOrXJenny(dataSourcePath, TerraformDataSourceJenny{}),
	// )

	return tgt
}
