package dashboard

import "github.com/grafana/grok/newgen/dashboard/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.StringOrArray
}

func ValString(ValString string) Option {
	return func(builder *Builder) error {

		builder.internal.ValString = &ValString

		return nil
	}
}

func ValArray(ValArray []string) Option {
	return func(builder *Builder) error {

		builder.internal.ValArray = &ValArray

		return nil
	}
}

func defaults() []Option {
	return []Option{}
}
