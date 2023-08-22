package dashboard

import (
	"errors"

	"github.com/grafana/grok/newgen/dashboard/types"
)

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

func Type(typeArg string) Option {
	return func(builder *Builder) error {
		if !(typeArg == "special") {
			return errors.New("typeArg must be == special")
		}

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
	return []Option{}
}
