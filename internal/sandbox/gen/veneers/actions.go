package veneers

import (
	"github.com/grafana/grok/internal/sandbox/gen/ast"
)

type OptionRewriteAction func(option ast.Option) []ast.Option

func RenameAction(newName string) OptionRewriteAction {
	return func(option ast.Option) []ast.Option {
		newOption := option
		newOption.Title = newName

		return []ast.Option{newOption}
	}
}

func OmitAction() OptionRewriteAction {
	return func(_ ast.Option) []ast.Option {
		return nil
	}
}

func PromoteToConstructorAction() OptionRewriteAction {
	return func(option ast.Option) []ast.Option {
		newOpt := option
		newOpt.IsConstructorArg = true

		return []ast.Option{newOpt}
	}
}

type BooleanUnfold struct {
	OptionTrue  string
	OptionFalse string
}

func UnfoldBooleanAction(unfoldOpts BooleanUnfold) OptionRewriteAction {
	return func(option ast.Option) []ast.Option {
		return []ast.Option{
			{
				Title:    unfoldOpts.OptionTrue,
				Comments: option.Comments,
				Args:     nil,
				Assignments: []ast.Assignment{
					{
						Path:              option.Assignments[0].Path,
						ValueType:         option.Assignments[0].ValueType,
						ArgumentName:      "",
						Value:             true,
						Constraints:       nil,
						IntoOptionalField: false,
						ValueHasBuilder:   false,
					},
				},
				// TODO: default
			},

			{
				Title:    unfoldOpts.OptionFalse,
				Comments: option.Comments,
				Args:     nil,
				Assignments: []ast.Assignment{
					{
						Path:              option.Assignments[0].Path,
						ValueType:         option.Assignments[0].ValueType,
						ArgumentName:      "",
						Value:             false,
						Constraints:       nil,
						IntoOptionalField: false,
						ValueHasBuilder:   false,
					},
				},
				// TODO: default
			},
		}
	}
}
