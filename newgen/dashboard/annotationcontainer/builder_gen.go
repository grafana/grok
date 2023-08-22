package annotationcontainer

import "github.com/grafana/grok/newgen/dashboard/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.AnnotationContainer
}

func New(options ...Option) (Builder, error) {
	resource := &types.AnnotationContainer{}
	builder := &Builder{internal: resource}

	for _, opt := range append(defaults(), options...) {
		if err := opt(builder); err != nil {
			return *builder, err
		}
	}

	return *builder, nil
}

func (builder *Builder) Internal() *types.AnnotationContainer {
	return builder.internal
}

func List(list []types.AnnotationQuery) Option {
	return func(builder *Builder) error {

		builder.internal.List = list

		return nil
	}
}

func defaults() []Option {
	return []Option{}
}
