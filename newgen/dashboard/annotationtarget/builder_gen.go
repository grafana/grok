package dashboard

import "github.com/grafana/grok/newgen/dashboard/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.AnnotationTarget
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
