package stringorbool

import "github.com/grafana/grok/newgen/dashboard/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.StringOrBool
}

func New(options ...Option) (Builder, error) {
	resource := &types.StringOrBool{}
	builder := &Builder{internal: resource}

	for _, opt := range append(defaults(), options...) {
		if err := opt(builder); err != nil {
			return *builder, err
		}
	}

	return *builder, nil
}

func (builder *Builder) Internal() *types.StringOrBool {
	return builder.internal
}

func ValString(ValString string) Option {
	return func(builder *Builder) error {

		builder.internal.ValString = &ValString

		return nil
	}
}

func ValBool(ValBool bool) Option {
	return func(builder *Builder) error {

		builder.internal.ValBool = &ValBool

		return nil
	}
}

func defaults() []Option {
	return []Option{}
}
