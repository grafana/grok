package objective

import "github.com/grafana/grok/newgen/slo/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.Objective
}
func New(options ...Option) (Builder, error) {
	resource := &types.Objective{}
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

func (builder *Builder) Internal() *types.Objective {
	return builder.internal
}
// is a value between 0 and 1 if the value of the query's output
// is above the objective, the SLO is met.
func Value(value float64) Option {
	return func(builder *Builder) error {
		
		builder.internal.Value = value

		return nil
	}
}
// is a Prometheus-parsable time duration string like 24h, 60m. This is the time
// window the objective is measured over.
func Window(window string) Option {
	return func(builder *Builder) error {
		
		builder.internal.Window = window

		return nil
	}
}

func defaults() []Option {
return []Option{
}
}
