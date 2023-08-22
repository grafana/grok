package dashboard

import "github.com/grafana/grok/newgen/dashboard/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.ValueMappingResult
}

func Text(text string) Option {
	return func(builder *Builder) error {

		builder.internal.Text = &text

		return nil
	}
}

func Color(color string) Option {
	return func(builder *Builder) error {

		builder.internal.Color = &color

		return nil
	}
}

func Icon(icon string) Option {
	return func(builder *Builder) error {

		builder.internal.Icon = &icon

		return nil
	}
}

func Index(index int32) Option {
	return func(builder *Builder) error {

		builder.internal.Index = &index

		return nil
	}
}

func defaults() []Option {
	return []Option{}
}
