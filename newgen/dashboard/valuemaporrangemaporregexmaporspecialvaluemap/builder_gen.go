package valuemaporrangemaporregexmaporspecialvaluemap

import (
	"github.com/grafana/grok/newgen/dashboard/rangemap"
	"github.com/grafana/grok/newgen/dashboard/regexmap"
	"github.com/grafana/grok/newgen/dashboard/specialvaluemap"
	"github.com/grafana/grok/newgen/dashboard/types"
	"github.com/grafana/grok/newgen/dashboard/valuemap"
)

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

func (builder *Builder) Internal() *types.ValueMapOrRangeMapOrRegexMapOrSpecialValueMap {
	return builder.internal
}

func ValValueMap(opts ...valuemap.Option) Option {
	return func(builder *Builder) error {
		resource, err := valuemap.New(opts...)
		if err != nil {
			return err
		}

		builder.internal.ValValueMap = resource.Internal()

		return nil
	}
}

func ValRangeMap(opts ...rangemap.Option) Option {
	return func(builder *Builder) error {
		resource, err := rangemap.New(opts...)
		if err != nil {
			return err
		}

		builder.internal.ValRangeMap = resource.Internal()

		return nil
	}
}

func ValRegexMap(opts ...regexmap.Option) Option {
	return func(builder *Builder) error {
		resource, err := regexmap.New(opts...)
		if err != nil {
			return err
		}

		builder.internal.ValRegexMap = resource.Internal()

		return nil
	}
}

func ValSpecialValueMap(opts ...specialvaluemap.Option) Option {
	return func(builder *Builder) error {
		resource, err := specialvaluemap.New(opts...)
		if err != nil {
			return err
		}

		builder.internal.ValSpecialValueMap = resource.Internal()

		return nil
	}
}

func defaults() []Option {
	return []Option{}
}
