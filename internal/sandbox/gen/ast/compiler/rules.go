package compiler

import (
	"github.com/grafana/grok/internal/sandbox/gen/rewrite"
)

func RewriteEngine() *Rewrite {
	return NewRewrite([]rewrite.FieldRewriteRule{
		// By default, use the field name as Option name
		rewrite.SetDefaultDisplayName(),

		// `Tooltip` looks better than `GraphTooltip`
		rewrite.RenameField(
			rewrite.ExactField("Dashboard", "graphTooltip"),
			"tooltip",
		),

		// We don't want that option at all
		rewrite.OmitField(
			rewrite.ExactField("Dashboard", "schemaVersion"),
		),

		// Editable() + Readonly() instead of Editable(val bool)
		rewrite.UnfoldBoolean(
			rewrite.ExactField("Dashboard", "editable"),
			rewrite.BooleanUnfold{OptionTrue: "readonly", OptionFalse: "editable"},
		),
	})
}
