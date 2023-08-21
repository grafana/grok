package dashboard

import "github.com/grafana/grok/newgen/dashboard/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.AnnotationPanelFilter
}

func Exclude(exclude bool) Option {
	return func(builder *Builder) error {
		
		builder.internal.Exclude = &exclude

		return nil
	}
}

func Ids(ids []uint8) Option {
	return func(builder *Builder) error {
		
		builder.internal.Ids = ids

		return nil
	}
}

func defaults() []Option {
return []Option{
Exclude(false),
}
}
