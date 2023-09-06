package veneers

import (
	"github.com/grafana/grok/internal/sandbox/gen/ast"
	"github.com/grafana/grok/internal/sandbox/gen/veneers/builder"
	"github.com/grafana/grok/internal/sandbox/gen/veneers/option"
)

type Rewriter struct {
	builderRules []builder.RewriteRule
	optionRules  []option.RewriteRule
}

func NewRewrite(builderRules []builder.RewriteRule, optionRules []option.RewriteRule) *Rewriter {
	return &Rewriter{
		builderRules: builderRules,
		optionRules:  optionRules,
	}
}

func (engine *Rewriter) ApplyTo(builders []ast.Builder) []ast.Builder {
	newBuilders := make([]ast.Builder, 0, len(builders))

	for _, b := range builders {
		processed := engine.processBuilder(b)
		// the builder was dismissed
		if len(processed.Options) == 0 {
			continue
		}

		newBuilders = append(newBuilders, processed)
	}

	return newBuilders
}

func (engine *Rewriter) processBuilder(builder ast.Builder) ast.Builder {
	processedBuilder := builder

	for _, rule := range engine.builderRules {
		if rule.Selector(processedBuilder) {
			processedBuilder = rule.Action(processedBuilder)
		}

		// this builder is dismissed, let's return early
		if len(processedBuilder.Options) == 0 {
			return processedBuilder
		}
	}

	processedOptions := make([]ast.Option, 0, len(processedBuilder.Options))
	for _, opt := range builder.Options {
		processedOptions = append(processedOptions, engine.processOption(builder, opt)...)
	}

	processedBuilder.Options = processedOptions

	return processedBuilder
}

func (engine *Rewriter) processOption(parentBuilder ast.Builder, opt ast.Option) []ast.Option {
	toProcess := []ast.Option{opt}
	for _, rule := range engine.optionRules {
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
