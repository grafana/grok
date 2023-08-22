package annotationpanelfilter

import "github.com/grafana/grok/newgen/dashboard/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.AnnotationPanelFilter
}

func New(options ...Option) (Builder, error) {
	resource := &types.AnnotationPanelFilter{}
	builder := &Builder{internal: resource}

	for _, opt := range append(defaults(), options...) {
		if err := opt(builder); err != nil {
			return *builder, err
		}
	}

	return *builder, nil
}

func (builder *Builder) Internal() *types.AnnotationPanelFilter {
	return builder.internal
}

func Exclude(exclude bool) Option {
	return func(builder *Builder) error {

		builder.internal.Exclude = &exclude

		return nil
	}
}

func Ids(ids []uint8) Option {
	return func(builder *Builder) error {

		builder.internal.Ids = ids

		return nil
	}
}

func defaults() []Option {
	return []Option{
		Exclude(false),
	}
}
