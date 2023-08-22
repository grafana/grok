package dashboard

import "github.com/grafana/grok/newgen/dashboard/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.Threshold
}

func New(options ...Option) (Builder, error) {
	resource := &types.Threshold{}
	builder := &Builder{internal: resource}

	for _, opt := range append(defaults(), options...) {
		if err := opt(builder); err != nil {
			return *builder, err
		}
	}

	return *builder, nil
}

func Value(value float64) Option {
	return func(builder *Builder) error {

		builder.internal.Value = &value

		return nil
	}
}

func Color(color string) Option {
	return func(builder *Builder) error {

		builder.internal.Color = color

		return nil
	}
}

func defaults() []Option {
	return []Option{}
}
