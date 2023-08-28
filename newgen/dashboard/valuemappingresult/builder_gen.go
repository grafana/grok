package valuemappingresult

import "github.com/grafana/grok/newgen/dashboard/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.ValueMappingResult
}
func New(options ...Option) (Builder, error) {
	resource := &types.ValueMappingResult{}
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

func (builder *Builder) Internal() *types.ValueMappingResult {
	return builder.internal
}
// Text to display when the value matches
func Text(text string) Option {
	return func(builder *Builder) error {
		
		builder.internal.Text = text

		return nil
	}
}
// Text to use when the value matches
func Color(color string) Option {
	return func(builder *Builder) error {
		
		builder.internal.Color = color

		return nil
	}
}
// Icon to display when the value matches. Only specific visualizations.
func Icon(icon string) Option {
	return func(builder *Builder) error {
		
		builder.internal.Icon = icon

		return nil
	}
}
// Position in the mapping array. Only used internally.
func Index(index int32) Option {
	return func(builder *Builder) error {
		
		builder.internal.Index = index

		return nil
	}
}

func defaults() []Option {
return []Option{
}
}
