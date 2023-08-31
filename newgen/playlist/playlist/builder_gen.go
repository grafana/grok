package playlist

import "github.com/grafana/grok/newgen/playlist/types"

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
// Name of the playlist.
func Name(name string) Option {
	return func(builder *Builder) error {
		
		builder.internal.Name = name

		return nil
	}
}
// Interval sets the time between switching views in a playlist.
// FIXME: Is this based on a standardized format or what options are available? Can datemath be used?
func Interval(interval string) Option {
	return func(builder *Builder) error {
		
		builder.internal.Interval = interval

		return nil
	}
}
// The ordered list of items that the playlist will iterate over.
// FIXME! This should not be optional, but changing it makes the godegen awkward
func Items(items []types.PlaylistItem) Option {
	return func(builder *Builder) error {
		
		builder.internal.Items = items

		return nil
	}
}
// Adding a required new field...
// This is only hear so that thema breaking change detection allows
// defining this as a new major version
func Xxx(xxx string) Option {
	return func(builder *Builder) error {
		
		builder.internal.Xxx = xxx

		return nil
	}
}

func defaults() []Option {
return []Option{
}
}
