package stringorarray

import "github.com/grafana/grok/newgen/dashboard/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.StringOrArray
}

func New(options ...Option) (Builder, error) {
	resource := &types.StringOrArray{}
	builder := &Builder{internal: resource}

	for _, opt := range append(defaults(), options...) {
		if err := opt(builder); err != nil {
			return *builder, err
		}
	}

	return *builder, nil
}

func (builder *Builder) Internal() *types.StringOrArray {
	return builder.internal
}

func ValString(ValString string) Option {
	return func(builder *Builder) error {

		builder.internal.ValString = &ValString

		return nil
	}
}

func ValArray(ValArray []string) Option {
	return func(builder *Builder) error {

		builder.internal.ValArray = &ValArray

		return nil
	}
}

func defaults() []Option {
	return []Option{}
}
