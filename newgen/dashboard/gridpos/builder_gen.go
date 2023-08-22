package dashboard

import "github.com/grafana/grok/newgen/dashboard/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.GridPos
}

func H(h uint32) Option {
	return func(builder *Builder) error {

		builder.internal.H = h

		return nil
	}
}

func W(w uint32) Option {
	return func(builder *Builder) error {

		builder.internal.W = w

		return nil
	}
}

func X(x uint32) Option {
	return func(builder *Builder) error {

		builder.internal.X = x

		return nil
	}
}

func Y(y uint32) Option {
	return func(builder *Builder) error {

		builder.internal.Y = y

		return nil
	}
}

func Static(static bool) Option {
	return func(builder *Builder) error {

		builder.internal.Static = &static

		return nil
	}
}

func defaults() []Option {
	return []Option{
		H(9),
		W(12),
		X(0),
		Y(0),
	}
}
