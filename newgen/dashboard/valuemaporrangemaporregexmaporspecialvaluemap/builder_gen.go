package valuemaporrangemaporregexmaporspecialvaluemap

import (
	"encoding/json"

	"github.com/grafana/grok/newgen/dashboard/types"
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

// MarshalJSON implements the encoding/json.Marshaler interface.
//
// This method can be used to render the resource as JSON
// which your configuration management tool of choice can then feed into
// Grafana.
func (builder *Builder) MarshalJSON() ([]byte, error) {
	return json.Marshal(builder.internal)
}

// MarshalIndentJSON renders the resource as indented JSON
// which your configuration management tool of choice can then feed into
// Grafana.
func (builder *Builder) MarshalIndentJSON() ([]byte, error) {
	return json.MarshalIndent(builder.internal, "", "  ")
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
