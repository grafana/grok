package rowpanel

import (
	"encoding/json"

	"github.com/grafana/grok/newgen/dashboard/datasourceref"
	"github.com/grafana/grok/newgen/dashboard/gridpos"
	"github.com/grafana/grok/newgen/dashboard/types"
)

type Option func(builder *Builder) error

type Builder struct {
	internal *types.RowPanel
}

func New(options ...Option) (Builder, error) {
	resource := &types.RowPanel{}
	builder := &Builder{internal: resource}

	for _, opt := range append(defaults(), options...) {
		if err := opt(builder); err != nil {
			return *builder, err
		}
	}

	return *builder, nil
}

// MarshalJSON implements the encoding/json.Marshaler interface.
//
// This method can be used to render the resource as JSON
// which your configuration management tool of choice can then feed into
// Grafana.
func (builder *Builder) MarshalJSON() ([]byte, error) {
	return json.Marshal(builder.internal)
}

// MarshalIndentJSON renders the resource as indented JSON
// which your configuration management tool of choice can then feed into
// Grafana.
func (builder *Builder) MarshalIndentJSON() ([]byte, error) {
	return json.MarshalIndent(builder.internal, "", "  ")
}

func (builder *Builder) Internal() *types.RowPanel {
	return builder.internal
}

func Type(typeArg string) Option {
	return func(builder *Builder) error {

		builder.internal.Type = typeArg

		return nil
	}
}

func Collapsed(collapsed bool) Option {
	return func(builder *Builder) error {

		builder.internal.Collapsed = collapsed

		return nil
	}
}

func Title(title string) Option {
	return func(builder *Builder) error {

		builder.internal.Title = &title

		return nil
	}
}

func Datasource(opts ...datasourceref.Option) Option {
	return func(builder *Builder) error {
		resource, err := datasourceref.New(opts...)
		if err != nil {
			return err
		}

		builder.internal.Datasource = resource.Internal()

		return nil
	}
}

func GridPos(opts ...gridpos.Option) Option {
	return func(builder *Builder) error {
		resource, err := gridpos.New(opts...)
		if err != nil {
			return err
		}

		builder.internal.GridPos = resource.Internal()

		return nil
	}
}

func Id(id uint32) Option {
	return func(builder *Builder) error {

		builder.internal.Id = id

		return nil
	}
}

func Panels(panels []types.Panel) Option {
	return func(builder *Builder) error {

		builder.internal.Panels = panels

		return nil
	}
}

func Repeat(repeat string) Option {
	return func(builder *Builder) error {

		builder.internal.Repeat = &repeat

		return nil
	}
}

func defaults() []Option {
	return []Option{
		Collapsed(false),
	}
}
