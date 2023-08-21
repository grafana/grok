package main

import (
	"fmt"

	"github.com/grafana/grok/newgen/dashboard"
	"github.com/grafana/grok/newgen/dashboard/types"
)

func main() {
	builder, err := dashboard.New(
		"Some title",
		dashboard.Description("Some description"),
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
