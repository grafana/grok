package thresholdquery

import "github.com/grafana/grok/newgen/slo/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.ThresholdQuery
}
func New(options ...Option) (Builder, error) {
	resource := &types.ThresholdQuery{}
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

func (builder *Builder) Internal() *types.ThresholdQuery {
	return builder.internal
}
func GroupByLabels(groupByLabels []string) Option {
	return func(builder *Builder) error {
		
		builder.internal.GroupByLabels = groupByLabels

		return nil
	}
}

func ThresholdMetric(opts ...metricdef.Option) Option {
	return func(builder *Builder) error {
		resource, err := metricdef.New(opts...)
		if err != nil {
			return err
		}

		builder.internal.ThresholdMetric = resource.Internal()

		return nil
	}
}

func Threshold(opts ...threshold.Option) Option {
	return func(builder *Builder) error {
		resource, err := threshold.New(opts...)
		if err != nil {
			return err
		}

		builder.internal.Threshold = resource.Internal()

		return nil
	}
}

func defaults() []Option {
return []Option{
}
}
