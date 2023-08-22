package librarypanelref

import (
	"encoding/json"

	"github.com/grafana/grok/newgen/dashboard/types"
)

type Option func(builder *Builder) error

type Builder struct {
	internal *types.LibraryPanelRef
}

func New(options ...Option) (Builder, error) {
	resource := &types.LibraryPanelRef{}
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

func (builder *Builder) Internal() *types.LibraryPanelRef {
	return builder.internal
}

func Name(name string) Option {
	return func(builder *Builder) error {

		builder.internal.Name = name

		return nil
	}
}

func Uid(uid string) Option {
	return func(builder *Builder) error {

		builder.internal.Uid = uid

		return nil
	}
}

func defaults() []Option {
	return []Option{}
}
