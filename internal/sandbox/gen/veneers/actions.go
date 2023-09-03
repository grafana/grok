package veneers

import (
	"github.com/grafana/grok/internal/sandbox/gen/ast"
)

type OptionRewriteAction func(option ast.Option) []ast.Option

func Rename(newName string) OptionRewriteAction {
	return func(option ast.Option) []ast.Option {
		newOption := option
		newOption.Title = newName

		return []ast.Option{newOption}
	}
}

func Omit() OptionRewriteAction {
	return func(_ ast.Option) []ast.Option {
		return nil
	}
}

type BooleanUnfold struct {
	OptionTrue  string
	OptionFalse string
}

func UnfoldBoolean(unfoldOpts BooleanUnfold) OptionRewriteAction {
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
