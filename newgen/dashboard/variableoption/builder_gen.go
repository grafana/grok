package variableoption

import (
	"encoding/json"

	"github.com/grafana/grok/newgen/dashboard/types"
)

type Option func(builder *Builder) error

type Builder struct {
	internal *types.VariableOption
}

func New(options ...Option) (Builder, error) {
	resource := &types.VariableOption{}
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

func (builder *Builder) Internal() *types.VariableOption {
	return builder.internal
}

// Whether the option is selected or not
func Selected(selected bool) Option {
	return func(builder *Builder) error {

		builder.internal.Selected = &selected

		return nil
	}
}

func Text(opts ...stringorarray.Option) Option {
	return func(builder *Builder) error {
		resource, err := stringorarray.New(opts...)
		if err != nil {
			return err
		}

		builder.internal.Text = resource.Internal()

		return nil
	}
}

func Value(opts ...stringorarray.Option) Option {
	return func(builder *Builder) error {
		resource, err := stringorarray.New(opts...)
		if err != nil {
			return err
		}

		builder.internal.Value = resource.Internal()

		return nil
	}
}

func defaults() []Option {
	return []Option{}
}
