package playlist

import "github.com/grafana/grok/newgen/playlist/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.Playlist
}

func Name(name string) Option {
	return func(builder *Builder) error {

		builder.internal.Name = name

		return nil
	}
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

func Xxx(xxx string) Option {
	return func(builder *Builder) error {

		builder.internal.Xxx = xxx

		return nil
	}
}
