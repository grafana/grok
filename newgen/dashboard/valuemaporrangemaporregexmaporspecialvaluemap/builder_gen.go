package dashboard

import "github.com/grafana/grok/newgen/dashboard/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.ValueMapOrRangeMapOrRegexMapOrSpecialValueMap
}

func New(options ...Option) (Builder, error) {
	resource := &types.ValueMapOrRangeMapOrRegexMapOrSpecialValueMap{}
	builder := &Builder{internal: resource}

	for _, opt := range append(defaults(), options...) {
		if err := opt(builder); err != nil {
			return *builder, err
		}
	}

	return *builder, nil
}

func ValValueMap(ValValueMap types.ValueMap) Option {
	return func(builder *Builder) error {

		builder.internal.ValValueMap = &ValValueMap

		return nil
	}
}

func ValRangeMap(ValRangeMap types.RangeMap) Option {
	return func(builder *Builder) error {

		builder.internal.ValRangeMap = &ValRangeMap

		return nil
	}
}

func ValRegexMap(ValRegexMap types.RegexMap) Option {
	return func(builder *Builder) error {

		builder.internal.ValRegexMap = &ValRegexMap

		return nil
	}
}

func ValSpecialValueMap(ValSpecialValueMap types.SpecialValueMap) Option {
	return func(builder *Builder) error {

		builder.internal.ValSpecialValueMap = &ValSpecialValueMap

		return nil
	}
}

func defaults() []Option {
	return []Option{}
}
