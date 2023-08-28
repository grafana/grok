package gridpos

import "github.com/grafana/grok/newgen/dashboard/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.GridPos
}
func New(options ...Option) (Builder, error) {
	resource := &types.GridPos{}
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

func (builder *Builder) Internal() *types.GridPos {
	return builder.internal
}
// Panel height. The height is the number of rows from the top edge of the panel.
func H(h uint32) Option {
	return func(builder *Builder) error {
		
		builder.internal.H = h

		return nil
	}
}
// Panel width. The width is the number of columns from the left edge of the panel.
func W(w uint32) Option {
	return func(builder *Builder) error {
		
		builder.internal.W = w

		return nil
	}
}
// Panel x. The x coordinate is the number of columns from the left edge of the grid
func X(x uint32) Option {
	return func(builder *Builder) error {
		
		builder.internal.X = x

		return nil
	}
}
// Panel y. The y coordinate is the number of rows from the top edge of the grid
func Y(y uint32) Option {
	return func(builder *Builder) error {
		
		builder.internal.Y = y

		return nil
	}
}
// Whether the panel is fixed within the grid. If true, the panel will not be affected by other panels' interactions
func Static(static bool) Option {
	return func(builder *Builder) error {
		
		builder.internal.Static = static

		return nil
	}
}

func defaults() []Option {
return []Option{
}
}
