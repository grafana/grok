package dashboard

import "github.com/grafana/grok/newgen/dashboard/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.FieldConfigSourceOverride
}

func Matcher(matcher types.MatcherConfig) Option {
	return func(builder *Builder) error {

		builder.internal.Matcher = matcher

		return nil
	}
}

func Properties(properties []types.DynamicConfigValue) Option {
	return func(builder *Builder) error {

		builder.internal.Properties = properties

		return nil
	}
}

func defaults() []Option {
	return []Option{}
}
