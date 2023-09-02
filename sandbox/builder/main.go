package main

import (
	"fmt"

	"github.com/grafana/grok/newgen/dashboard/dashboard"
	"github.com/grafana/grok/newgen/dashboard/timepicker"
	"github.com/grafana/grok/newgen/dashboard/types"
)

func main() {
	builder, err := dashboard.New(
		"Some title",
		dashboard.Description("Some description"),
		dashboard.Timepicker(
			timepicker.RefreshIntervals([]string{"30s", "1m", "5m"}),
		),
		dashboard.Style(types.StyleDark),
		dashboard.Timezone("utc"),
		dashboard.Tooltip(types.Crosshair),
		dashboard.Tags([]string{"generated", "from", "cue"}),
		dashboard.Links([]types.DashboardLink{
			{
				Title:       "Some link",
				Url:         "http://google.com",
				AsDropdown:  false,
				TargetBlank: true,
			},
		}),
	)
	if err != nil {
		panic(err)
	}

	dashboardJson, err := builder.MarshalIndentJSON()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dashboardJson))
}
