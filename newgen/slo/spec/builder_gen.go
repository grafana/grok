package spec

import (
	"encoding/json"

	"github.com/grafana/grok/newgen/slo/alerting"
	"github.com/grafana/grok/newgen/slo/grafanametadata"
	"github.com/grafana/grok/newgen/slo/types"
)

type Option func(builder *Builder) error

type Builder struct {
	internal *types.Spec
}

func New(options ...Option) (Builder, error) {
	resource := &types.Spec{}
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

func (builder *Builder) Internal() *types.Spec {
	return builder.internal
}

func GrafanaMetadata(opts ...grafanametadata.Option) Option {
	return func(builder *Builder) error {
		resource, err := grafanametadata.New(opts...)
		if err != nil {
			return err
		}

		builder.internal.GrafanaMetadata = resource.Internal()

		return nil
	}
}

// A unique, random identifier. This value will also be the name of the
// resource stored in the API server. Must be set for a PUT.
func Uuid(uuid string) Option {
	return func(builder *Builder) error {

		builder.internal.Uuid = uuid

		return nil
	}
}

// should be a short description of your indicator. Consider names like
// "API Availability"
func Name(name string) Option {
	return func(builder *Builder) error {

		builder.internal.Name = name

		return nil
	}
}

// is a free-text field that can provide more context to an
// SLO. It is shown on SLO drill-down dashboards and in hover text on
// the SLO summary dashboard.
func Description(description string) Option {
	return func(builder *Builder) error {

		builder.internal.Description = description

		return nil
	}
}

// describes the indicator that will be measured against the
// objective. Four query types are supported:
//  1. Ratio Queries provide a successMetric and totalMetric whose ratio is the SLI.
//  2. Threshold Queries provide a thresholdMetric and a threshold. The
//     SLI is the boolean result of evaluating the threshould.
//  3. Histogram Queries are similar to threshold queries, but the use a
//     Prometheus histogram metric, percentile value, and a threshold to
//     generate the boolean output.
//  4. Freeform Queries supply a single freeFormQuery string that is
//     evaluated to produce the SLI output. The value should range beween 0
//     and 1.0. Freeform queries should include a time variable named
//     either `$__rate_interval`,`$__interval` or `$__range`. This will be used by the
//     tool to evaluate the burn rate of an SLO over various time
//     windows. Queries that don't include this interval will have
//     sensitive and imprecise alerting.
//
// Additionally, "groupByLabels" are used in the first three query types
// to define how to group series for evaluation. They are discarded for
// freeform queries.
func Query(query types.Query) Option {
	return func(builder *Builder) error {

		builder.internal.Query = query

		return nil
	}
}

// You can have multiple time windows and objectives associated with an
// SLO. Over each rolling time window, the remaining error budget will
// be calculated, and separate alerts can be generated for each time
// window based on the SLO burn rate or remaining error budget.
func Objectives(objectives []types.Objective) Option {
	return func(builder *Builder) error {

		builder.internal.Objectives = objectives

		return nil
	}
}

// Any additional labels that will be attached to all metrics generated
// from the query. These labels are useful for grouping SLOs in
// dashboard views that you create by hand.
// The key must match the prometheus label requirements regex:
// "^[a-zA-Z_][a-zA-Z0-9_]*$"
func Labels(labels []types.Label) Option {
	return func(builder *Builder) error {

		builder.internal.Labels = labels

		return nil
	}
}

func Alerting(opts ...alerting.Option) Option {
	return func(builder *Builder) error {
		resource, err := alerting.New(opts...)
		if err != nil {
			return err
		}

		builder.internal.Alerting = resource.Internal()

		return nil
	}
}

func defaults() []Option {
	return []Option{}
}