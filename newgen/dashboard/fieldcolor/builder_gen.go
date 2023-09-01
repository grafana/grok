package fieldcolor

import "github.com/grafana/grok/newgen/dashboard/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.FieldColor
}
func New(options ...Option) (Builder, error) {
	resource := &types.FieldColor{}
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

func (builder *Builder) Internal() *types.FieldColor {
	return builder.internal
}
// The main color scheme mode.
func Mode(mode types.FieldColorModeId) Option {
	return func(builder *Builder) error {
		
		builder.internal.Mode = mode

		return nil
	}
}
// The fixed color value for fixed or shades color modes.
func FixedColor(fixedColor string) Option {
	return func(builder *Builder) error {
		
		builder.internal.FixedColor = &fixedColor

		return nil
	}
}
// Some visualizations need to know how to assign a series color from by value color schemes.
func SeriesBy(seriesBy types.FieldColorSeriesByMode) Option {
	return func(builder *Builder) error {
		
		builder.internal.SeriesBy = &seriesBy

		return nil
	}
}

func defaults() []Option {
return []Option{
}
}
