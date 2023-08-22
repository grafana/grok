package gridpos

import (
	"errors"

	"github.com/grafana/grok/newgen/dashboard/types"
)

type Option func(builder *Builder) error

type Builder struct {
	internal *types.GridPos
}

func New(options ...Option) (Builder, error) {
	resource := &types.GridPos{}
	builder := &Builder{internal: resource}

	for _, opt := range append(defaults(), options...) {
		if err := opt(builder); err != nil {
			return *builder, err
		}
	}

	return *builder, nil
}

func (builder *Builder) Internal() *types.GridPos {
	return builder.internal
}

func H(h uint32) Option {
	return func(builder *Builder) error {
		if !(h > 0) {
			return errors.New("h must be > 0")
		}

		builder.internal.H = h

		return nil
	}
}

func W(w uint32) Option {
	return func(builder *Builder) error {
		if !(w > 0) {
			return errors.New("w must be > 0")
		}

		if !(w <= 24) {
			return errors.New("w must be <= 24")
		}

		builder.internal.W = w

		return nil
	}
}

func X(x uint32) Option {
	return func(builder *Builder) error {
		if !(x >= 0) {
			return errors.New("x must be >= 0")
		}

		if !(x < 24) {
			return errors.New("x must be < 24")
		}

		builder.internal.X = x

		return nil
	}
}

func Y(y uint32) Option {
	return func(builder *Builder) error {
		if !(y >= 0) {
			return errors.New("y must be >= 0")
		}

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
