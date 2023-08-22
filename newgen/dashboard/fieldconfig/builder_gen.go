package fieldconfig

import (
	"github.com/grafana/grok/newgen/dashboard/fieldcolor"
	"github.com/grafana/grok/newgen/dashboard/thresholdsconfig"
	"github.com/grafana/grok/newgen/dashboard/types"
)

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

func (builder *Builder) Internal() *types.FieldConfig {
	return builder.internal
}

func DisplayName(displayName string) Option {
	return func(builder *Builder) error {

		builder.internal.DisplayName = &displayName

		return nil
	}
}

func DisplayNameFromDS(displayNameFromDS string) Option {
	return func(builder *Builder) error {

		builder.internal.DisplayNameFromDS = &displayNameFromDS

		return nil
	}
}

func Description(description string) Option {
	return func(builder *Builder) error {

		builder.internal.Description = &description

		return nil
	}
}

func Path(path string) Option {
	return func(builder *Builder) error {

		builder.internal.Path = &path

		return nil
	}
}

func Writeable(writeable bool) Option {
	return func(builder *Builder) error {

		builder.internal.Writeable = &writeable

		return nil
	}
}

func Filterable(filterable bool) Option {
	return func(builder *Builder) error {

		builder.internal.Filterable = &filterable

		return nil
	}
}

func Unit(unit string) Option {
	return func(builder *Builder) error {

		builder.internal.Unit = &unit

		return nil
	}
}

func Decimals(decimals float64) Option {
	return func(builder *Builder) error {

		builder.internal.Decimals = &decimals

		return nil
	}
}

func Min(min float64) Option {
	return func(builder *Builder) error {

		builder.internal.Min = &min

		return nil
	}
}

func Max(max float64) Option {
	return func(builder *Builder) error {

		builder.internal.Max = &max

		return nil
	}
}

func Mappings(mappings []types.ValueMapOrRangeMapOrRegexMapOrSpecialValueMap) Option {
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

func Links(links []any) Option {
	return func(builder *Builder) error {

		builder.internal.Links = links

		return nil
	}
}

func NoValue(noValue string) Option {
	return func(builder *Builder) error {

		builder.internal.NoValue = &noValue

		return nil
	}
}

func Custom(custom any) Option {
	return func(builder *Builder) error {

		builder.internal.Custom = &custom

		return nil
	}
}

func defaults() []Option {
	return []Option{}
}
