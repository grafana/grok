package dashboard

import "github.com/grafana/grok/newgen/dashboard/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.LibraryPanelRef
}

func Name(name string) Option {
	return func(builder *Builder) error {
		
		builder.internal.Name = name

		return nil
	}
}

func Uid(uid string) Option {
	return func(builder *Builder) error {
		
		builder.internal.Uid = uid

		return nil
	}
}

func defaults() []Option {
return []Option{
}
}
