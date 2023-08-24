package playlist

import (
	"encoding/json"

	"github.com/grafana/grok/newgen/playlist/types"
)

type Option func(builder *Builder) error

type Builder struct {
	internal *types.Playlist
}

func New(options ...Option) (Builder, error) {
	resource := &types.Playlist{}
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

func (builder *Builder) Internal() *types.Playlist {
	return builder.internal
}

func Interval(interval string) Option {
	return func(builder *Builder) error {

		builder.internal.Interval = interval

		return nil
	}
}

func Items(items []types.PlaylistItem) Option {
	return func(builder *Builder) error {

		builder.internal.Items = items

		return nil
	}
}

func Name(name string) Option {
	return func(builder *Builder) error {

		builder.internal.Name = name

		return nil
	}
}

func Xxx(xxx string) Option {
	return func(builder *Builder) error {

		builder.internal.Xxx = xxx

		return nil
	}
}

func defaults() []Option {
	return []Option{}
}
