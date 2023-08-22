package panel

import (
	"encoding/json"
	"errors"

	"github.com/grafana/grok/newgen/dashboard/datasourceref"
	"github.com/grafana/grok/newgen/dashboard/fieldconfigsource"
	"github.com/grafana/grok/newgen/dashboard/gridpos"
	"github.com/grafana/grok/newgen/dashboard/librarypanelref"
	"github.com/grafana/grok/newgen/dashboard/types"
)

type Option func(builder *Builder) error

type Builder struct {
	internal *types.Panel
}

func New(options ...Option) (Builder, error) {
	resource := &types.Panel{}
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

func (builder *Builder) Internal() *types.Panel {
	return builder.internal
}

func Type(typeArg string) Option {
	return func(builder *Builder) error {
		if !(len([]rune(typeArg)) >= 1) {
			return errors.New("typeArg must be minLength 1")
		}

		builder.internal.Type = typeArg

		return nil
	}
}

func Id(id uint32) Option {
	return func(builder *Builder) error {

		builder.internal.Id = &id

		return nil
	}
}

func PluginVersion(pluginVersion string) Option {
	return func(builder *Builder) error {

		builder.internal.PluginVersion = &pluginVersion

		return nil
	}
}

func Tags(tags []string) Option {
	return func(builder *Builder) error {

		builder.internal.Tags = tags

		return nil
	}
}

func Targets(targets []types.Target) Option {
	return func(builder *Builder) error {

		builder.internal.Targets = targets

		return nil
	}
}

func Title(title string) Option {
	return func(builder *Builder) error {

		builder.internal.Title = &title

		return nil
	}
}

func Description(description string) Option {
	return func(builder *Builder) error {

		builder.internal.Description = &description

		return nil
	}
}

func Transparent(transparent bool) Option {
	return func(builder *Builder) error {

		builder.internal.Transparent = transparent

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

func GridPos(opts ...gridpos.Option) Option {
	return func(builder *Builder) error {
		resource, err := gridpos.New(opts...)
		if err != nil {
			return err
		}

		builder.internal.GridPos = resource.Internal()

		return nil
	}
}

func Links(links []types.DashboardLink) Option {
	return func(builder *Builder) error {

		builder.internal.Links = links

		return nil
	}
}

func Repeat(repeat string) Option {
	return func(builder *Builder) error {

		builder.internal.Repeat = &repeat

		return nil
	}
}

func RepeatDirection(repeatDirection types.PanelRepeatDirection) Option {
	return func(builder *Builder) error {

		builder.internal.RepeatDirection = &repeatDirection

		return nil
	}
}

func RepeatPanelId(repeatPanelId int64) Option {
	return func(builder *Builder) error {

		builder.internal.RepeatPanelId = &repeatPanelId

		return nil
	}
}

func MaxDataPoints(maxDataPoints float64) Option {
	return func(builder *Builder) error {

		builder.internal.MaxDataPoints = &maxDataPoints

		return nil
	}
}

func Transformations(transformations []types.DataTransformerConfig) Option {
	return func(builder *Builder) error {

		builder.internal.Transformations = transformations

		return nil
	}
}

func Interval(interval string) Option {
	return func(builder *Builder) error {

		builder.internal.Interval = &interval

		return nil
	}
}

func TimeFrom(timeFrom string) Option {
	return func(builder *Builder) error {

		builder.internal.TimeFrom = &timeFrom

		return nil
	}
}

func TimeShift(timeShift string) Option {
	return func(builder *Builder) error {

		builder.internal.TimeShift = &timeShift

		return nil
	}
}

func LibraryPanel(opts ...librarypanelref.Option) Option {
	return func(builder *Builder) error {
		resource, err := librarypanelref.New(opts...)
		if err != nil {
			return err
		}

		builder.internal.LibraryPanel = resource.Internal()

		return nil
	}
}

func Options(options any) Option {
	return func(builder *Builder) error {

		builder.internal.Options = options

		return nil
	}
}

func FieldConfig(opts ...fieldconfigsource.Option) Option {
	return func(builder *Builder) error {
		resource, err := fieldconfigsource.New(opts...)
		if err != nil {
			return err
		}

		builder.internal.FieldConfig = resource.Internal()

		return nil
	}
}

func defaults() []Option {
	return []Option{
		Transparent(false),
	}
}
