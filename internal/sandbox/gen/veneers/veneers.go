package veneers

func Engine() *Rewriter {
	return NewRewrite([]OptionRewriteRule{
		/********************************************
		 * Dashboards
		 ********************************************/

		// Let's make the dashboard constructor more friendly
		PromoteToConstructor(
			ExactOption("Dashboard", "title"),
		),

		// `Tooltip` looks better than `GraphTooltip`
		Rename(
			ExactOption("Dashboard", "graphTooltip"),
			"tooltip",
		),

		// We don't want that option at all
		Omit(
			ExactOption("Dashboard", "schemaVersion"),
		),

		// Editable() + Readonly() instead of Editable(val bool)
		UnfoldBoolean(
			ExactOption("Dashboard", "editable"),
			BooleanUnfold{OptionTrue: "readonly", OptionFalse: "editable"},
		),

		/********************************************
		 * Rows
		 ********************************************/

		// Let's make the row constructor more friendly
		PromoteToConstructor(
			ExactOption("RowPanel", "title"),
		),
	})
}
