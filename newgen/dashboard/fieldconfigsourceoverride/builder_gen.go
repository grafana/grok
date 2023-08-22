package fieldconfigsourceoverride

import (
	"github.com/grafana/grok/newgen/dashboard/matcherconfig"
	"github.com/grafana/grok/newgen/dashboard/types"
)

type Option func(builder *Builder) error

type Builder struct {
	internal *types.FieldConfigSourceOverride
}

func New(options ...Option) (Builder, error) {
	resource := &types.FieldConfigSourceOverride{}
	builder := &Builder{internal: resource}

	for _, opt := range append(defaults(), options...) {
		if err := opt(builder); err != nil {
			return *builder, err
		}
	}

	return *builder, nil
}

func (builder *Builder) Internal() *types.FieldConfigSourceOverride {
	return builder.internal
}

func Matcher(opts ...matcherconfig.Option) Option {
	return func(builder *Builder) error {
		resource, err := matcherconfig.New(opts...)
		if err != nil {
			return err
		}

		builder.internal.Matcher = resource.Internal()

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
