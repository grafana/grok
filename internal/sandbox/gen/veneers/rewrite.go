package veneers

import (
	"github.com/grafana/grok/internal/sandbox/gen/ast"
)

type Rewriter struct {
	optionRules []OptionRewriteRule
}

func NewRewrite(fieldRules []OptionRewriteRule) *Rewriter {
	return &Rewriter{
		optionRules: fieldRules,
	}
}

func (pass *Rewriter) ApplyTo(builders []ast.Builder) []ast.Builder {
	newBuilders := make([]ast.Builder, 0, len(builders))

	for _, builder := range builders {
		newBuilders = append(newBuilders, pass.processBuilder(builder))
	}

	return newBuilders
}

func (pass *Rewriter) processBuilder(builder ast.Builder) ast.Builder {
	processedOptions := make([]ast.Option, 0, len(builder.Options))
	for _, opt := range builder.Options {
		processedOptions = append(processedOptions, pass.processOption(builder, opt)...)
	}

	return ast.Builder{
		Package: builder.Package,
		For:     builder.For,
		Options: processedOptions,
	}
}

func (pass *Rewriter) processOption(parentBuilder ast.Builder, opt ast.Option) []ast.Option {
	toProcess := []ast.Option{opt}
	for _, rule := range pass.optionRules {
		if !rule.Selector(parentBuilder, opt) {
			continue
		}

		var wip []ast.Option
		for _, modifiedField := range toProcess {
			wip = append(wip, rule.Action(modifiedField)...)
		}
		toProcess = wip
	}

	return toProcess
}
