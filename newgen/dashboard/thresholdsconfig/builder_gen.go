package dashboard

import "github.com/grafana/grok/newgen/dashboard/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.ThresholdsConfig
}

func New(options ...Option) (Builder, error) {
	resource := &types.ThresholdsConfig{}
	builder := &Builder{internal: resource}

	for _, opt := range append(defaults(), options...) {
		if err := opt(builder); err != nil {
			return *builder, err
		}
	}

	return *builder, nil
}

func Mode(mode types.ThresholdsMode) Option {
	return func(builder *Builder) error {

		builder.internal.Mode = mode

		return nil
	}
}

func Steps(steps []types.Threshold) Option {
	return func(builder *Builder) error {

		builder.internal.Steps = steps

		return nil
	}
}

func defaults() []Option {
	return []Option{}
}
