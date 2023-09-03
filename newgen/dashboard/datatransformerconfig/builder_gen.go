package datatransformerconfig

import (
	"encoding/json"

	"github.com/grafana/grok/newgen/dashboard/types"
)

type Option func(builder *Builder) error

type Builder struct {
	internal *types.DataTransformerConfig
}

func New(options ...Option) (Builder, error) {
	resource := &types.DataTransformerConfig{}
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

func (builder *Builder) Internal() *types.DataTransformerConfig {
	return builder.internal
}

// Unique identifier of transformer
func Id(id string) Option {
	return func(builder *Builder) error {

		builder.internal.Id = id

		return nil
	}
}

// Disabled transformations are skipped
func Disabled(disabled bool) Option {
	return func(builder *Builder) error {

		builder.internal.Disabled = &disabled

		return nil
	}
}

func Filter(opts ...matcherconfig.Option) Option {
	return func(builder *Builder) error {
		resource, err := matcherconfig.New(opts...)
		if err != nil {
			return err
		}

		builder.internal.Filter = resource.Internal()

		return nil
	}
}

// Options to be passed to the transformer
// Valid options depend on the transformer id
func Options(options any) Option {
	return func(builder *Builder) error {

		builder.internal.Options = options

		return nil
	}
}

func defaults() []Option {
	return []Option{}
}
