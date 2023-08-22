package dashboard

import "github.com/grafana/grok/newgen/dashboard/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.DashboardTemplating
}

func New(options ...Option) (Builder, error) {
	resource := &types.DashboardTemplating{}
	builder := &Builder{internal: resource}

	for _, opt := range append(defaults(), options...) {
		if err := opt(builder); err != nil {
			return *builder, err
		}
	}

	return *builder, nil
}

func List(list []types.VariableModel) Option {
	return func(builder *Builder) error {

		builder.internal.List = list

		return nil
	}
}

func defaults() []Option {
	return []Option{}
}
