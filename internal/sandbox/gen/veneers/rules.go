package veneers

type OptionRewriteRule struct {
	Selector OptionSelector
	Action   OptionRewriteAction
}

func Rename(selector OptionSelector, newName string) OptionRewriteRule {
	return OptionRewriteRule{
		Selector: selector,
		Action:   RenameAction(newName),
	}
}

func Omit(selector OptionSelector) OptionRewriteRule {
	return OptionRewriteRule{
		Selector: selector,
		Action:   OmitAction(),
	}
}

func UnfoldBoolean(selector OptionSelector, unfoldOpts BooleanUnfold) OptionRewriteRule {
	return OptionRewriteRule{
		Selector: selector,
		Action:   UnfoldBooleanAction(unfoldOpts),
	}
}

func PromoteToConstructor(selector OptionSelector) OptionRewriteRule {
	return OptionRewriteRule{
		Selector: selector,
		Action:   PromoteToConstructorAction(),
	}
}
