package dashboard

import "github.com/grafana/grok/newgen/dashboard/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.VariableOption
}

func New(options ...Option) (Builder, error) {
	resource := &types.VariableOption{}
	builder := &Builder{internal: resource}

	for _, opt := range append(defaults(), options...) {
		if err := opt(builder); err != nil {
			return *builder, err
		}
	}

	return *builder, nil
}

func Selected(selected bool) Option {
	return func(builder *Builder) error {

		builder.internal.Selected = &selected

		return nil
	}
}

func Text(text types.StringOrArray) Option {
	return func(builder *Builder) error {

		builder.internal.Text = text

		return nil
	}
}

func Value(value types.StringOrArray) Option {
	return func(builder *Builder) error {

		builder.internal.Value = value

		return nil
	}
}

func defaults() []Option {
	return []Option{}
}
