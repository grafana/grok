package status

import "github.com/grafana/grok/newgen/slo/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.Status
}
func New(options ...Option) (Builder, error) {
	resource := &types.Status{}
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

func (builder *Builder) Internal() *types.Status {
	return builder.internal
}
func DrillDownDashboard(drillDownDashboard struct {
	Uid string `json:"uid"`
	// The generation of the SLO when this dashboard was last updated.
ReconciledForGeneration string `json:"reconciledForGeneration"`
	LastError string `json:"lastError"`
}) Option {
	return func(builder *Builder) error {
		
		builder.internal.DrillDownDashboard = drillDownDashboard

		return nil
	}
}
func OperatorState(operatorState struct {
	// lastEvaluation is the ResourceVersion last evaluated
LastEvaluation string `json:"lastEvaluation"`
	// state describes the state of the lastEvaluation.
// It is limited to three possible states for machine evaluation.
State types.StateEnum `json:"state"`
	// descriptiveState is an optional more descriptive state field which has no requirements on format
DescriptiveState *string `json:"descriptiveState,omitempty"`
	// details contains any extra information that is operator-specific
Details any `json:"details,omitempty"`
}) Option {
	return func(builder *Builder) error {
		
		builder.internal.OperatorState = operatorState

		return nil
	}
}
// operatorStates is a map of operator ID to operator state evaluations.
// Any operator which consumes this kind SHOULD add its state evaluation information to this field.
func OperatorStates(operatorStates any) Option {
	return func(builder *Builder) error {
		
		builder.internal.OperatorStates = &operatorStates

		return nil
	}
}
func PrometheusRules(prometheusRules struct {
	// The generation of the SLO when these rules were last updated.
ReconciledForGeneration string `json:"reconciledForGeneration"`
	LastError string `json:"lastError"`
}) Option {
	return func(builder *Builder) error {
		
		builder.internal.PrometheusRules = prometheusRules

		return nil
	}
}
// additionalFields is reserved for future use
func AdditionalFields(additionalFields any) Option {
	return func(builder *Builder) error {
		
		builder.internal.AdditionalFields = &additionalFields

		return nil
	}
}

func defaults() []Option {
return []Option{
}
}
