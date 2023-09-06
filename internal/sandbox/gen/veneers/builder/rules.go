package builder

type RewriteRule struct {
	Selector Selector
	Action   RewriteAction
}

func Omit(selector Selector) RewriteRule {
	return RewriteRule{
		Selector: selector,
		Action:   OmitAction(),
	}
}
