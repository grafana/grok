package veneers

func Engine() *Rewriter {
	return NewRewrite([]OptionRewriteRule{
		// `Tooltip` looks better than `GraphTooltip`
		RenameOption(
			ExactOption("Dashboard", "graphTooltip"),
			"tooltip",
		),

		// We don't want that option at all
		OmitOption(
			ExactOption("Dashboard", "schemaVersion"),
		),

		// Editable() + Readonly() instead of Editable(val bool)
		UnfoldBooleanOption(
			ExactOption("Dashboard", "editable"),
			BooleanUnfold{OptionTrue: "readonly", OptionFalse: "editable"},
		),
	})
}
