package dashboard

import "github.com/grafana/grok/newgen/dashboard/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.MatcherConfig
}

func Id(id string) Option {
	return func(builder *Builder) error {
		
		builder.internal.Id = id

		return nil
	}
}

func Options(options any) Option {
	return func(builder *Builder) error {
		
		builder.internal.Options = &options

		return nil
	}
}

func defaults() []Option {
return []Option{
Id(""),
}
}
