package variablemodel

import "github.com/grafana/grok/newgen/dashboard/types"

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

func (builder *Builder) Internal() *types.VariableModel {
	return builder.internal
}
// Unique numeric identifier for the variable.
func Id(id string) Option {
	return func(builder *Builder) error {
		
		builder.internal.Id = id

		return nil
	}
}
// Type of variable
func Type(typeArg types.VariableType) Option {
	return func(builder *Builder) error {
		
		builder.internal.Type = typeArg

		return nil
	}
}
// Name of variable
func Name(name string) Option {
	return func(builder *Builder) error {
		
		builder.internal.Name = name

		return nil
	}
}
// Optional display name
func Label(label string) Option {
	return func(builder *Builder) error {
		
		builder.internal.Label = &label

		return nil
	}
}
// Visibility configuration for the variable
func Hide(hide types.VariableHide) Option {
	return func(builder *Builder) error {
		
		builder.internal.Hide = hide

		return nil
	}
}
// Whether the variable value should be managed by URL query params or not
func SkipUrlSync(skipUrlSync bool) Option {
	return func(builder *Builder) error {
		
		builder.internal.SkipUrlSync = skipUrlSync

		return nil
	}
}
// Description of variable. It can be defined but `null`.
func Description(description string) Option {
	return func(builder *Builder) error {
		
		builder.internal.Description = &description

		return nil
	}
}
// Query used to fetch values for a variable
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
// Format to use while fetching all values from data source, eg: wildcard, glob, regex, pipe, etc.
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
// Whether multiple values can be selected or not from variable value list
func Multi(multi bool) Option {
	return func(builder *Builder) error {
		
		builder.internal.Multi = &multi

		return nil
	}
}
// Options that can be selected for a variable.
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
}
}
