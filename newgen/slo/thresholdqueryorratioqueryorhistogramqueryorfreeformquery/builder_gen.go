package thresholdqueryorratioqueryorhistogramqueryorfreeformquery

import "github.com/grafana/grok/newgen/slo/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.ThresholdQueryOrRatioQueryOrHistogramQueryOrFreeformQuery
}
func New(options ...Option) (Builder, error) {
	resource := &types.ThresholdQueryOrRatioQueryOrHistogramQueryOrFreeformQuery{}
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

func (builder *Builder) Internal() *types.ThresholdQueryOrRatioQueryOrHistogramQueryOrFreeformQuery {
	return builder.internal
}

func ValThresholdQuery(opts ...thresholdquery.Option) Option {
	return func(builder *Builder) error {
		resource, err := thresholdquery.New(opts...)
		if err != nil {
			return err
		}

		builder.internal.ValThresholdQuery = resource.Internal()

		return nil
	}
}

func ValRatioQuery(opts ...ratioquery.Option) Option {
	return func(builder *Builder) error {
		resource, err := ratioquery.New(opts...)
		if err != nil {
			return err
		}

		builder.internal.ValRatioQuery = resource.Internal()

		return nil
	}
}

func ValHistogramQuery(opts ...histogramquery.Option) Option {
	return func(builder *Builder) error {
		resource, err := histogramquery.New(opts...)
		if err != nil {
			return err
		}

		builder.internal.ValHistogramQuery = resource.Internal()

		return nil
	}
}

func ValFreeformQuery(opts ...freeformquery.Option) Option {
	return func(builder *Builder) error {
		resource, err := freeformquery.New(opts...)
		if err != nil {
			return err
		}

		builder.internal.ValFreeformQuery = resource.Internal()

		return nil
	}
}

func defaults() []Option {
return []Option{
}
}
