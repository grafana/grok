package dashboard

import (
	"errors"

	"github.com/grafana/grok/newgen/dashboard/types"
)

type Option func(builder *Builder) error

type Builder struct {
	internal *types.ValueMap
}

func Type(typeArg string) Option {
	return func(builder *Builder) error {
		if !(typeArg == "value") {
			return errors.New("typeArg must be == value")
		}

		builder.internal.Type = typeArg

		return nil
	}
}

func Options(options map[string]types.ValueMappingResult) Option {
	return func(builder *Builder) error {

		builder.internal.Options = options

		return nil
	}
}

func defaults() []Option {
	return []Option{}
}
