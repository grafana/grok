package annotationtarget

import (
	"encoding/json"

	"github.com/grafana/grok/newgen/dashboard/types"
)

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

func (builder *Builder) Internal() *types.AnnotationTarget {
	return builder.internal
}

// Only required/valid for the grafana datasource...
// but code+tests is already depending on it so hard to change
func Limit(limit int64) Option {
	return func(builder *Builder) error {

		builder.internal.Limit = limit

		return nil
	}
}

// Only required/valid for the grafana datasource...
// but code+tests is already depending on it so hard to change
func MatchAny(matchAny bool) Option {
	return func(builder *Builder) error {

		builder.internal.MatchAny = matchAny

		return nil
	}
}

// Only required/valid for the grafana datasource...
// but code+tests is already depending on it so hard to change
func Tags(tags []string) Option {
	return func(builder *Builder) error {

		builder.internal.Tags = tags

		return nil
	}
}

// Only required/valid for the grafana datasource...
// but code+tests is already depending on it so hard to change
func Type(typeArg string) Option {
	return func(builder *Builder) error {

		builder.internal.Type = typeArg

		return nil
	}
}

func defaults() []Option {
	return []Option{}
}
