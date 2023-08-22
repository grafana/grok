package dashboard

import "github.com/grafana/grok/newgen/dashboard/types"

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

func Name(name string) Option {
	return func(builder *Builder) error {

		builder.internal.Name = name

		return nil
	}
}

func Datasource(datasource types.DataSourceRef) Option {
	return func(builder *Builder) error {

		builder.internal.Datasource = datasource

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

func Filter(filter types.AnnotationPanelFilter) Option {
	return func(builder *Builder) error {

		builder.internal.Filter = &filter

		return nil
	}
}

func Target(target types.AnnotationTarget) Option {
	return func(builder *Builder) error {

		builder.internal.Target = &target

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
