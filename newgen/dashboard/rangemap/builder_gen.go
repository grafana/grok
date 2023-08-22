package dashboard

import (
	"errors"

	"github.com/grafana/grok/newgen/dashboard/types"
)

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

func Type(typeArg string) Option {
	return func(builder *Builder) error {
		if !(typeArg == "range") {
			return errors.New("typeArg must be == range")
		}

		builder.internal.Type = typeArg

		return nil
	}
}

func Options(options struct {
	// Min value of the range. It can be null which means -Infinity
	From *float64 `json:"from"`
	// Max value of the range. It can be null which means +Infinity
	To *float64 `json:"to"`
	// Config to apply when the value is within the range
	Result types.ValueMappingResult `json:"result"`
}) Option {
	return func(builder *Builder) error {

		builder.internal.Options = options

		return nil
	}
}

func defaults() []Option {
	return []Option{}
}
