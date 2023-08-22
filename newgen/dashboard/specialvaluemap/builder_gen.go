package dashboard

import (
	"errors"

	"github.com/grafana/grok/newgen/dashboard/types"
)

type Option func(builder *Builder) error

type Builder struct {
	internal *types.SpecialValueMap
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
