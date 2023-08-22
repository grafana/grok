package dashboard

import "github.com/grafana/grok/newgen/dashboard/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.TimeInterval
}

func New(options ...Option) (Builder, error) {
	resource := &types.TimeInterval{}
	builder := &Builder{internal: resource}

	for _, opt := range append(defaults(), options...) {
		if err := opt(builder); err != nil {
			return *builder, err
		}
	}

	return *builder, nil
}

func From(from string) Option {
	return func(builder *Builder) error {

		builder.internal.From = from

		return nil
	}
}

func To(to string) Option {
	return func(builder *Builder) error {

		builder.internal.To = to

		return nil
	}
}

func defaults() []Option {
	return []Option{
		From("now-6h"),
		To("now"),
	}
}
