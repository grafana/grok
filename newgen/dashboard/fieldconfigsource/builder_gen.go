package fieldconfigsource

import (
	"github.com/grafana/grok/newgen/dashboard/fieldconfig"
	"github.com/grafana/grok/newgen/dashboard/types"
)

type Option func(builder *Builder) error

type Builder struct {
	internal *types.FieldConfigSource
}

func New(options ...Option) (Builder, error) {
	resource := &types.FieldConfigSource{}
	builder := &Builder{internal: resource}

	for _, opt := range append(defaults(), options...) {
		if err := opt(builder); err != nil {
			return *builder, err
		}
	}

	return *builder, nil
}

func (builder *Builder) Internal() *types.FieldConfigSource {
	return builder.internal
}

func Defaults(opts ...fieldconfig.Option) Option {
	return func(builder *Builder) error {
		resource, err := fieldconfig.New(opts...)
		if err != nil {
			return err
		}

		builder.internal.Defaults = resource.Internal()

		return nil
	}
}

func Overrides(overrides []types.FieldConfigSourceOverride) Option {
	return func(builder *Builder) error {

		builder.internal.Overrides = overrides

		return nil
	}
}

func defaults() []Option {
	return []Option{}
}
