package dashboard

import "github.com/grafana/grok/newgen/dashboard/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.DashboardTemplating
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
