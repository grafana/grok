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

func JenniesForGo() jen.TargetJennies {
	tgt := jen.NewTargetJennies()

	tgt.Core.Append(
		codegen.LatestMajorsOrXJenny(filepath.Join("kinds", "core"), codegen.GoTypesJenny{}),
		codegen.LatestMajorsOrXJenny(filepath.Join("kinds", "core"), TypedSchemaJenny{}),
	)

	// tgt.Composable.Append()

	return tgt
}
