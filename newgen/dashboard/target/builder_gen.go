package target

import "github.com/grafana/grok/newgen/dashboard/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.Target
}

func New(options ...Option) (Builder, error) {
	resource := &types.Target{}
	builder := &Builder{internal: resource}

	for _, opt := range append(defaults(), options...) {
		if err := opt(builder); err != nil {
			return *builder, err
		}
	}

	return *builder, nil
}

func (builder *Builder) Internal() *types.Target {
	return builder.internal
}

func defaults() []Option {
	return []Option{}
}
