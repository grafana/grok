package operatorstate

import "github.com/grafana/grok/newgen/slo/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.OperatorState
}
func New(options ...Option) (Builder, error) {
	resource := &types.OperatorState{}
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

func (builder *Builder) Internal() *types.OperatorState {
	return builder.internal
}
// lastEvaluation is the ResourceVersion last evaluated
func LastEvaluation(lastEvaluation string) Option {
	return func(builder *Builder) error {
		
		builder.internal.LastEvaluation = lastEvaluation

		return nil
	}
}
// state describes the state of the lastEvaluation.
// It is limited to three possible states for machine evaluation.
func State(state types.StateEnum) Option {
	return func(builder *Builder) error {
		
		builder.internal.State = state

		return nil
	}
}
// descriptiveState is an optional more descriptive state field which has no requirements on format
func DescriptiveState(descriptiveState string) Option {
	return func(builder *Builder) error {
		
		builder.internal.DescriptiveState = &descriptiveState

		return nil
	}
}
// details contains any extra information that is operator-specific
func Details(details any) Option {
	return func(builder *Builder) error {
		
		builder.internal.Details = &details

		return nil
	}
}

func defaults() []Option {
return []Option{
}
}
