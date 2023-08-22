package dashboard

import "github.com/grafana/grok/newgen/dashboard/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.DataSourceRef
}

func New(options ...Option) (Builder, error) {
	resource := &types.DataSourceRef{}
	builder := &Builder{internal: resource}

	for _, opt := range append(defaults(), options...) {
		if err := opt(builder); err != nil {
			return *builder, err
		}
	}

	return *builder, nil
}

func Type(typeArg string) Option {
	return func(builder *Builder) error {

		builder.internal.Type = &typeArg

		return nil
	}
}

func Uid(uid string) Option {
	return func(builder *Builder) error {

		builder.internal.Uid = &uid

		return nil
	}
}

func defaults() []Option {
	return []Option{}
}
