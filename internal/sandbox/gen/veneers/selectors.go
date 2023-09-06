package veneers

import (
	"github.com/grafana/grok/internal/sandbox/gen/ast"
)

type OptionSelector func(builder ast.Builder, option ast.Option) bool

func ExactOption(objectName string, optionName string) OptionSelector {
	return func(builder ast.Builder, option ast.Option) bool {
		return builder.For.Name == objectName && option.Name == optionName
	}
}

func EveryOption() OptionSelector {
	return func(builder ast.Builder, option ast.Option) bool {
		return true
	}
}
