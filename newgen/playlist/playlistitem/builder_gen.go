package playlistitem

import "github.com/grafana/grok/newgen/playlist/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.PlaylistItem
}
func New(options ...Option) (Builder, error) {
	resource := &types.PlaylistItem{}
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

func (builder *Builder) Internal() *types.PlaylistItem {
	return builder.internal
}
// Type of the item.
func Type(typeArg types.TypeEnum) Option {
	return func(builder *Builder) error {
		
		builder.internal.Type = typeArg

		return nil
	}
}
// Value depends on type and describes the playlist item.
// 
//  - dashboard_by_id: The value is an internal numerical identifier set by Grafana. This
//  is not portable as the numerical identifier is non-deterministic between different instances.
//  Will be replaced by dashboard_by_uid in the future. (deprecated)
//  - dashboard_by_tag: The value is a tag which is set on any number of dashboards. All
//  dashboards behind the tag will be added to the playlist.
//  - dashboard_by_uid: The value is the dashboard UID
func Value(value string) Option {
	return func(builder *Builder) error {
		
		builder.internal.Value = value

		return nil
	}
}
// Title is an unused property -- it will be removed in the future
func Title(title string) Option {
	return func(builder *Builder) error {
		
		builder.internal.Title = title

		return nil
	}
}

func defaults() []Option {
return []Option{
}
}
