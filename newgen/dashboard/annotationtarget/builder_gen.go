package annotationtarget

import "github.com/grafana/grok/newgen/dashboard/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.AnnotationTarget
}

func New(options ...Option) (Builder, error) {
	resource := &types.AnnotationTarget{}
	builder := &Builder{internal: resource}

	for _, opt := range append(defaults(), options...) {
		if err := opt(builder); err != nil {
			return *builder, err
		}
	}

	return *builder, nil
}

func (builder *Builder) Internal() *types.AnnotationTarget {
	return builder.internal
}

func Limit(limit int64) Option {
	return func(builder *Builder) error {

		builder.internal.Limit = limit

		return nil
	}
}

func MatchAny(matchAny bool) Option {
	return func(builder *Builder) error {

		builder.internal.MatchAny = matchAny

		return nil
	}
}

func Tags(tags []string) Option {
	return func(builder *Builder) error {

		builder.internal.Tags = tags

		return nil
	}
}

func Type(typeArg string) Option {
	return func(builder *Builder) error {

		builder.internal.Type = typeArg

		return nil
	}
}

func defaults() []Option {
	return []Option{}
}
