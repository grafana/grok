package dashboard

import "github.com/grafana/grok/newgen/dashboard/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.DynamicConfigValue
}

func Id(id string) Option {
	return func(builder *Builder) error {
		
		builder.internal.Id = id

		return nil
	}
}

func Value(value any) Option {
	return func(builder *Builder) error {
		
		builder.internal.Value = &value

		return nil
	}
}

func defaults() []Option {
return []Option{
Id(""),
}
}
