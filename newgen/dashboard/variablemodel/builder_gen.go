package variablemodel

import (
	"github.com/grafana/grok/newgen/dashboard/datasourceref"
	"github.com/grafana/grok/newgen/dashboard/types"
	"github.com/grafana/grok/newgen/dashboard/variableoption"
)

type Option func(builder *Builder) error

type Builder struct {
	internal *types.VariableModel
}

func New(options ...Option) (Builder, error) {
	resource := &types.VariableModel{}
	builder := &Builder{internal: resource}

	for _, opt := range append(defaults(), options...) {
		if err := opt(builder); err != nil {
			return *builder, err
		}
	}

	return *builder, nil
}

func (builder *Builder) Internal() *types.VariableModel {
	return builder.internal
}

func Id(id string) Option {
	return func(builder *Builder) error {

		builder.internal.Id = id

		return nil
	}
}

func Type(typeArg types.VariableType) Option {
	return func(builder *Builder) error {

		builder.internal.Type = typeArg

		return nil
	}
}

func Name(name string) Option {
	return func(builder *Builder) error {

		builder.internal.Name = name

		return nil
	}
}

func Label(label string) Option {
	return func(builder *Builder) error {

		builder.internal.Label = &label

		return nil
	}
}

func Hide(hide types.VariableHide) Option {
	return func(builder *Builder) error {

		builder.internal.Hide = hide

		return nil
	}
}

func SkipUrlSync(skipUrlSync bool) Option {
	return func(builder *Builder) error {

		builder.internal.SkipUrlSync = skipUrlSync

		return nil
	}
}

func Description(description string) Option {
	return func(builder *Builder) error {

		builder.internal.Description = &description

		return nil
	}
}

func Query(query any) Option {
	return func(builder *Builder) error {

		builder.internal.Query = &query

		return nil
	}
}

func Datasource(opts ...datasourceref.Option) Option {
	return func(builder *Builder) error {
		resource, err := datasourceref.New(opts...)
		if err != nil {
			return err
		}

		builder.internal.Datasource = resource.Internal()

		return nil
	}
}

func AllFormat(allFormat string) Option {
	return func(builder *Builder) error {

		builder.internal.AllFormat = &allFormat

		return nil
	}
}

func Current(opts ...variableoption.Option) Option {
	return func(builder *Builder) error {
		resource, err := variableoption.New(opts...)
		if err != nil {
			return err
		}

		builder.internal.Current = resource.Internal()

		return nil
	}
}

func Multi(multi bool) Option {
	return func(builder *Builder) error {

		builder.internal.Multi = &multi

		return nil
	}
}

func Options(options []types.VariableOption) Option {
	return func(builder *Builder) error {

		builder.internal.Options = options

		return nil
	}
}

func Refresh(refresh types.VariableRefresh) Option {
	return func(builder *Builder) error {

		builder.internal.Refresh = &refresh

		return nil
	}
}

func defaults() []Option {
	return []Option{
		Id("00000000-0000-0000-0000-000000000000"),
		SkipUrlSync(false),
		Multi(false),
	}
}
