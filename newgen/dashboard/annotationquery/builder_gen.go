package annotationquery

import (
	"encoding/json"

	"github.com/grafana/grok/newgen/dashboard/annotationpanelfilter"
	"github.com/grafana/grok/newgen/dashboard/annotationtarget"
	"github.com/grafana/grok/newgen/dashboard/datasourceref"
	"github.com/grafana/grok/newgen/dashboard/types"
)

type Option func(builder *Builder) error

type Builder struct {
	internal *types.AnnotationQuery
}

func New(options ...Option) (Builder, error) {
	resource := &types.AnnotationQuery{}
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

func (builder *Builder) Internal() *types.AnnotationQuery {
	return builder.internal
}

func Name(name string) Option {
	return func(builder *Builder) error {

		builder.internal.Name = name

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

func Enable(enable bool) Option {
	return func(builder *Builder) error {

		builder.internal.Enable = enable

		return nil
	}
}

func Hide(hide bool) Option {
	return func(builder *Builder) error {

		builder.internal.Hide = &hide

		return nil
	}
}

func IconColor(iconColor string) Option {
	return func(builder *Builder) error {

		builder.internal.IconColor = iconColor

		return nil
	}
}

func Filter(opts ...annotationpanelfilter.Option) Option {
	return func(builder *Builder) error {
		resource, err := annotationpanelfilter.New(opts...)
		if err != nil {
			return err
		}

		builder.internal.Filter = resource.Internal()

		return nil
	}
}

func Target(opts ...annotationtarget.Option) Option {
	return func(builder *Builder) error {
		resource, err := annotationtarget.New(opts...)
		if err != nil {
			return err
		}

		builder.internal.Target = resource.Internal()

		return nil
	}
}

func Type(typeArg string) Option {
	return func(builder *Builder) error {

		builder.internal.Type = &typeArg

		return nil
	}
}

func defaults() []Option {
	return []Option{
		Enable(true),
		Hide(false),
	}
}