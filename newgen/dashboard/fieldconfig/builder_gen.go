package fieldconfig

import "github.com/grafana/grok/newgen/dashboard/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.FieldConfig
}
func New(options ...Option) (Builder, error) {
	resource := &types.FieldConfig{}
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

func (builder *Builder) Internal() *types.FieldConfig {
	return builder.internal
}
// The display value for this field.  This supports template variables blank is auto
func DisplayName(displayName string) Option {
	return func(builder *Builder) error {
		
		builder.internal.DisplayName = displayName

		return nil
	}
}
// This can be used by data sources that return and explicit naming structure for values and labels
// When this property is configured, this value is used rather than the default naming strategy.
func DisplayNameFromDS(displayNameFromDS string) Option {
	return func(builder *Builder) error {
		
		builder.internal.DisplayNameFromDS = displayNameFromDS

		return nil
	}
}
// Human readable field metadata
func Description(description string) Option {
	return func(builder *Builder) error {
		
		builder.internal.Description = description

		return nil
	}
}
// An explicit path to the field in the datasource.  When the frame meta includes a path,
// This will default to `${frame.meta.path}/${field.name}
// 
// When defined, this value can be used as an identifier within the datasource scope, and
// may be used to update the results
func Path(path string) Option {
	return func(builder *Builder) error {
		
		builder.internal.Path = path

		return nil
	}
}
// True if data source can write a value to the path. Auth/authz are supported separately
func Writeable(writeable bool) Option {
	return func(builder *Builder) error {
		
		builder.internal.Writeable = writeable

		return nil
	}
}
// True if data source field supports ad-hoc filters
func Filterable(filterable bool) Option {
	return func(builder *Builder) error {
		
		builder.internal.Filterable = filterable

		return nil
	}
}
// Unit a field should use. The unit you select is applied to all fields except time.
// You can use the units ID availables in Grafana or a custom unit.
// Available units in Grafana: https://github.com/grafana/grafana/blob/main/packages/grafana-data/src/valueFormats/categories.ts
// As custom unit, you can use the following formats:
// `suffix:<suffix>` for custom unit that should go after value.
// `prefix:<prefix>` for custom unit that should go before value.
// `time:<format>` For custom date time formats type for example `time:YYYY-MM-DD`.
// `si:<base scale><unit characters>` for custom SI units. For example: `si: mF`. This one is a bit more advanced as you can specify both a unit and the source data scale. So if your source data is represented as milli (thousands of) something prefix the unit with that SI scale character.
// `count:<unit>` for a custom count unit.
// `currency:<unit>` for custom a currency unit.
func Unit(unit string) Option {
	return func(builder *Builder) error {
		
		builder.internal.Unit = unit

		return nil
	}
}
// Specify the number of decimals Grafana includes in the rendered value.
// If you leave this field blank, Grafana automatically truncates the number of decimals based on the value.
// For example 1.1234 will display as 1.12 and 100.456 will display as 100.
// To display all decimals, set the unit to `String`.
func Decimals(decimals float64) Option {
	return func(builder *Builder) error {
		
		builder.internal.Decimals = decimals

		return nil
	}
}
// The minimum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.
func Min(min float64) Option {
	return func(builder *Builder) error {
		
		builder.internal.Min = min

		return nil
	}
}
// The maximum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.
func Max(max float64) Option {
	return func(builder *Builder) error {
		
		builder.internal.Max = max

		return nil
	}
}
// Convert input values into a display string
func Mappings(mappings []disjunction<types.ValueMap | types.RangeMap | types.RegexMap | types.SpecialValueMap>) Option {
	return func(builder *Builder) error {
		
		builder.internal.Mappings = mappings

		return nil
	}
}

func Thresholds(opts ...thresholdsconfig.Option) Option {
	return func(builder *Builder) error {
		resource, err := thresholdsconfig.New(opts...)
		if err != nil {
			return err
		}

		builder.internal.Thresholds = resource.Internal()

		return nil
	}
}

func Color(opts ...fieldcolor.Option) Option {
	return func(builder *Builder) error {
		resource, err := fieldcolor.New(opts...)
		if err != nil {
			return err
		}

		builder.internal.Color = resource.Internal()

		return nil
	}
}
// The behavior when clicking on a result
func Links(links []any) Option {
	return func(builder *Builder) error {
		
		builder.internal.Links = links

		return nil
	}
}
// Alternative to empty string
func NoValue(noValue string) Option {
	return func(builder *Builder) error {
		
		builder.internal.NoValue = noValue

		return nil
	}
}
// custom is specified by the FieldConfig field
// in panel plugin schemas.
func Custom(custom any) Option {
	return func(builder *Builder) error {
		
		builder.internal.Custom = custom

		return nil
	}
}

func defaults() []Option {
return []Option{
}
}
