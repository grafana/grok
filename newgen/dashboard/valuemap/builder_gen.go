package dashboard

import "github.com/grafana/grok/newgen/dashboard/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.ValueMap
}

func Type(type string) Option {
	return func(builder *Builder) error {
		if !(type == value) {
return errors.New("type must be == value")
}

		builder.internal.Type = type

		return nil
	}
}

func Options(options map[string]types.ValueMappingResult) Option {
	return func(builder *Builder) error {
		
		builder.internal.Options = options

		return nil
	}
}

func defaults() []Option {
return []Option{
}
}
