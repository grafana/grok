package dashboard

import "github.com/grafana/grok/newgen/dashboard/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.Threshold
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
