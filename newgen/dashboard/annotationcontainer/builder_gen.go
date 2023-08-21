package dashboard

import "github.com/grafana/grok/newgen/dashboard/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.AnnotationContainer
}

func List(list []types.AnnotationQuery) Option {
	return func(builder *Builder) error {
		
		builder.internal.List = list

		return nil
	}
}

func defaults() []Option {
return []Option{
}
}
