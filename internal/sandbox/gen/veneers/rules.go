package veneers

type OptionRewriteRule struct {
	Selector OptionSelector
	Action   OptionRewriteAction
}

func RenameOption(selector OptionSelector, newName string) OptionRewriteRule {
	return OptionRewriteRule{
		Selector: selector,
		Action:   Rename(newName),
	}
}

func OmitOption(selector OptionSelector) OptionRewriteRule {
	return OptionRewriteRule{
		Selector: selector,
		Action:   Omit(),
	}
}

func UnfoldBooleanOption(selector OptionSelector, unfoldOpts BooleanUnfold) OptionRewriteRule {
	return OptionRewriteRule{
		Selector: selector,
		Action:   UnfoldBoolean(unfoldOpts),
	}
}
