package annotationquery

import (
	"encoding/json"

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

// Name of annotation.
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

// When enabled the annotation query is issued with every dashboard refresh
func Enable(enable bool) Option {
	return func(builder *Builder) error {

		builder.internal.Enable = enable

		return nil
	}
}

// Annotation queries can be toggled on or off at the top of the dashboard.
// When hide is true, the toggle is not shown in the dashboard.
func Hide(hide bool) Option {
	return func(builder *Builder) error {

		builder.internal.Hide = &hide

		return nil
	}
}

// Color to use for the annotation event markers
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

// TODO -- this should not exist here, it is based on the --grafana-- datasource
func Type(typeArg string) Option {
	return func(builder *Builder) error {

		builder.internal.Type = &typeArg

		return nil
	}
}

func defaults() []Option {
	return []Option{}
}
