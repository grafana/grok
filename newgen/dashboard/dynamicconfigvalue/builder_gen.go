package dashboard

import "github.com/grafana/grok/newgen/dashboard/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.DynamicConfigValue
}

func New(options ...Option) (Builder, error) {
	resource := &types.DynamicConfigValue{}
	builder := &Builder{internal: resource}

	for _, opt := range append(defaults(), options...) {
		if err := opt(builder); err != nil {
			return *builder, err
		}
	}

	return *builder, nil
}

func Id(id string) Option {
	return func(builder *Builder) error {

		builder.internal.Id = id

		return nil
	}
}

func Value(value any) Option {
	return func(builder *Builder) error {

		builder.internal.Value = &value

		return nil
	}
}

func defaults() []Option {
	return []Option{
		Id(""),
	}
}
