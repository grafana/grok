package dashboard

import "github.com/grafana/grok/newgen/dashboard/types"

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

func Defaults(defaults types.FieldConfig) Option {
	return func(builder *Builder) error {

		builder.internal.Defaults = defaults

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
