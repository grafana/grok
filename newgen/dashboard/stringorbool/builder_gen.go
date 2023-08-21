package dashboard

import "github.com/grafana/grok/newgen/dashboard/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.StringOrBool
}

func ValString(ValString string) Option {
	return func(builder *Builder) error {
		
		builder.internal.ValString = &ValString

		return nil
	}
}

func ValBool(ValBool bool) Option {
	return func(builder *Builder) error {
		
		builder.internal.ValBool = &ValBool

		return nil
	}
}

func defaults() []Option {
return []Option{
}
}
