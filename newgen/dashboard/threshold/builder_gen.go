package threshold

import "github.com/grafana/grok/newgen/dashboard/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.Threshold
}
func New(options ...Option) (Builder, error) {
	resource := &types.Threshold{}
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

func (builder *Builder) Internal() *types.Threshold {
	return builder.internal
}
// Value represents a specified metric for the threshold, which triggers a visual change in the dashboard when this value is met or exceeded.
// Nulls currently appear here when serializing -Infinity to JSON.
func Value(value float64) Option {
	return func(builder *Builder) error {
		
		builder.internal.Value = value

		return nil
	}
}
// Color represents the color of the visual change that will occur in the dashboard when the threshold value is met or exceeded.
func Color(color string) Option {
	return func(builder *Builder) error {
		
		builder.internal.Color = color

		return nil
	}
}

func defaults() []Option {
return []Option{
}
}
