package valuemap

import (
	"errors"

	"github.com/grafana/grok/newgen/dashboard/types"
)

type Option func(builder *Builder) error

type Builder struct {
	internal *types.ValueMap
}

func New(options ...Option) (Builder, error) {
	resource := &types.ValueMap{}
	builder := &Builder{internal: resource}

	for _, opt := range append(defaults(), options...) {
		if err := opt(builder); err != nil {
			return *builder, err
		}
	}

	return *builder, nil
}

func (builder *Builder) Internal() *types.ValueMap {
	return builder.internal
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
