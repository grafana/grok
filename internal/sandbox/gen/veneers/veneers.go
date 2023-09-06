package veneers

import (
	"github.com/grafana/grok/internal/sandbox/gen/veneers/builder"
	"github.com/grafana/grok/internal/sandbox/gen/veneers/option"
)

func Engine() *Rewriter {
	return NewRewrite(
		[]builder.RewriteRule{
			// We don't want these builders at all
			builder.Omit(builder.ByName("GridPos")),
			builder.Omit(builder.ByName("DataSourceRef")),
			builder.Omit(builder.ByName("LibraryPanelRef")),

			// rearrange things a bit
			builder.MergeInto(
				builder.ByName("Panel"),
				"FieldConfig",
				"fieldConfig.defaults",
				// don't copy these over as they clash with a similarly named option from Panel
				[]string{"description", "links"},
			),

			// remove builders that were previously merged into something else
			builder.Omit(builder.ByName("FieldConfig")),
			builder.Omit(builder.ByName("FieldConfigSource")),
		},

		[]option.RewriteRule{
			/********************************************
			 * Dashboards
			 ********************************************/

			// Let's make the dashboard constructor more friendly
			option.PromoteToConstructor(
				option.ByName("Dashboard", "title"),
			),

			// `Tooltip` looks better than `GraphTooltip`
			option.Rename(
				option.ByName("Dashboard", "graphTooltip"),
				"tooltip",
			),

			// Editable() + Readonly() instead of Editable(val bool)
			option.UnfoldBoolean(
				option.ByName("Dashboard", "editable"),
				option.BooleanUnfold{OptionTrue: "editable", OptionFalse: "readonly"},
			),

			// Time(from, to) instead of time(struct {From string `json:"from"`, To   string `json:"to"`}{From: "lala", To: "lala})
			option.StructFieldsAsArguments(
				option.ByName("Dashboard", "time"),
			),

			// We don't want these options at all
			option.Omit(option.ByName("Dashboard", "schemaVersion")),

			/********************************************
			 * Panels
			 ********************************************/

			option.Omit(option.ByName("Panel", "fieldConfig")),
			option.Omit(option.ByName("Panel", "options")), // comes from a panel plugin
			option.Omit(option.ByName("Panel", "custom")),  // comes from a panel plugin

			/********************************************
			 * Rows
			 ********************************************/

			// Let's make the row constructor more friendly
			option.PromoteToConstructor(
				option.ByName("RowPanel", "title"),
			),
		},
	)
}
