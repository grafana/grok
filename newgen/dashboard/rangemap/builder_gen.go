package rangemap

import "github.com/grafana/grok/newgen/dashboard/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.RangeMap
}
func New(options ...Option) (Builder, error) {
	resource := &types.RangeMap{}
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

func (builder *Builder) Internal() *types.RangeMap {
	return builder.internal
}
func Type(typeArg string) Option {
	return func(builder *Builder) error {
		
		builder.internal.Type = typeArg

		return nil
	}
}
// Range to match against and the result to apply when the value is within the range
func Options(options struct {
	// Min value of the range. It can be null which means -Infinity
From disjunction<float64 | null> `json:"from"`
	// Max value of the range. It can be null which means +Infinity
To disjunction<float64 | null> `json:"to"`
	// Config to apply when the value is within the range
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
