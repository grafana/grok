package timeinterval

import (
	"encoding/json"

	"github.com/grafana/grok/newgen/dashboard/types"
)

type Option func(builder *Builder) error

type Builder struct {
	internal *types.TimeInterval
}

func New(options ...Option) (Builder, error) {
	resource := &types.TimeInterval{}
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

func (builder *Builder) Internal() *types.TimeInterval {
	return builder.internal
}

func From(from string) Option {
	return func(builder *Builder) error {

		builder.internal.From = from

		return nil
	}
}

func To(to string) Option {
	return func(builder *Builder) error {

		builder.internal.To = to

		return nil
	}
}

func defaults() []Option {
	return []Option{
		From("now-6h"),
		To("now"),
	}
}
