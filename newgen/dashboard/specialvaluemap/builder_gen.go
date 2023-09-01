package specialvaluemap

import "github.com/grafana/grok/newgen/dashboard/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.SpecialValueMap
}
func New(options ...Option) (Builder, error) {
	resource := &types.SpecialValueMap{}
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

func (builder *Builder) Internal() *types.SpecialValueMap {
	return builder.internal
}
func Type(typeArg string) Option {
	return func(builder *Builder) error {
		
		builder.internal.Type = typeArg

		return nil
	}
}
func Options(options struct {
	// Special value to match against
Match types.SpecialValueMatch `json:"match"`
	// Config to apply when the value matches the special value
Result types.ValueMappingResult `json:"result"`
}) Option {
	return func(builder *Builder) error {
		
		builder.internal.Options = options

		return nil
	}
}

func defaults() []Option {
return []Option{
}
}
