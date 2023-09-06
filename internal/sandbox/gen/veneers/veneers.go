package veneers

import (
	"github.com/grafana/grok/internal/sandbox/gen/veneers/builder"
	"github.com/grafana/grok/internal/sandbox/gen/veneers/option"
)

func Engine() *Rewriter {
	return NewRewrite(
		[]builder.RewriteRule{
			// We don't want that builder at all
			builder.Omit(
				builder.ExactBuilder("GridPos"),
			),
		},

		[]option.RewriteRule{
			/********************************************
			 * Dashboards
			 ********************************************/

			// Let's make the dashboard constructor more friendly
			option.PromoteToConstructor(
				option.ExactOption("Dashboard", "title"),
			),

			// `Tooltip` looks better than `GraphTooltip`
			option.Rename(
				option.ExactOption("Dashboard", "graphTooltip"),
				"tooltip",
			),

			// We don't want that option at all
			option.Omit(
				option.ExactOption("Dashboard", "schemaVersion"),
			),

			// Editable() + Readonly() instead of Editable(val bool)
			option.UnfoldBoolean(
				option.ExactOption("Dashboard", "editable"),
				option.BooleanUnfold{OptionTrue: "editable", OptionFalse: "readonly"},
			),

			// Time(from, to) instead of time(struct {From string `json:"from"`, To   string `json:"to"`}{From: "lala", To: "lala})
			option.StructFieldsAsArguments(
				option.ExactOption("Dashboard", "time"),
			),

			/********************************************
			 * Rows
			 ********************************************/

			// Let's make the row constructor more friendly
			option.PromoteToConstructor(
				option.ExactOption("RowPanel", "title"),
			),
		},
	)
}
