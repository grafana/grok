package option

type RewriteRule struct {
	Selector Selector
	Action   RewriteAction
}

func Rename(selector Selector, newName string) RewriteRule {
	return RewriteRule{
		Selector: selector,
		Action:   RenameAction(newName),
	}
}

func Omit(selector Selector) RewriteRule {
	return RewriteRule{
		Selector: selector,
		Action:   OmitAction(),
	}
}

func UnfoldBoolean(selector Selector, unfoldOpts BooleanUnfold) RewriteRule {
	return RewriteRule{
		Selector: selector,
		Action:   UnfoldBooleanAction(unfoldOpts),
	}
}

func PromoteToConstructor(selector Selector) RewriteRule {
	return RewriteRule{
		Selector: selector,
		Action:   PromoteToConstructorAction(),
	}
}

func StructFieldsAsArguments(selector Selector, explicitFields ...string) RewriteRule {
	return RewriteRule{
		Selector: selector,
		Action:   StructFieldsAsArgumentsAction(explicitFields...),
	}
}
