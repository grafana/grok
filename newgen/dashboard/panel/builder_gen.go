package panel

import (
	"encoding/json"
	"errors"

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

// The panel plugin type id. This is used to find the plugin to display the panel.
func Type(typeArg string) Option {
	return func(builder *Builder) error {
		if !(len([]rune(typeArg)) >= 1) {
			return errors.New("typeArg must be minLength 1")
		}

		builder.internal.Type = typeArg

		return nil
	}
}

// Unique identifier of the panel. Generated by Grafana when creating a new panel. It must be unique within a dashboard, but not globally.
func Id(id uint32) Option {
	return func(builder *Builder) error {

		builder.internal.Id = &id

		return nil
	}
}

// The version of the plugin that is used for this panel. This is used to find the plugin to display the panel and to migrate old panel configs.
func PluginVersion(pluginVersion string) Option {
	return func(builder *Builder) error {

		builder.internal.PluginVersion = &pluginVersion

		return nil
	}
}

// Tags for the panel.
func Tags(tags []string) Option {
	return func(builder *Builder) error {

		builder.internal.Tags = tags

		return nil
	}
}

// Depends on the panel plugin. See the plugin documentation for details.
func Targets(targets []types.Target) Option {
	return func(builder *Builder) error {

		builder.internal.Targets = targets

		return nil
	}
}

// Panel title.
func Title(title string) Option {
	return func(builder *Builder) error {

		builder.internal.Title = &title

		return nil
	}
}

// Panel description.
func Description(description string) Option {
	return func(builder *Builder) error {

		builder.internal.Description = &description

		return nil
	}
}

// Whether to display the panel without a background.
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

// Panel links.
func Links(links []types.DashboardLink) Option {
	return func(builder *Builder) error {

		builder.internal.Links = links

		return nil
	}
}

// Name of template variable to repeat for.
func Repeat(repeat string) Option {
	return func(builder *Builder) error {

		builder.internal.Repeat = &repeat

		return nil
	}
}

// Direction to repeat in if 'repeat' is set.
// `h` for horizontal, `v` for vertical.
func RepeatDirection(repeatDirection types.PanelRepeatDirection) Option {
	return func(builder *Builder) error {

		builder.internal.RepeatDirection = &repeatDirection

		return nil
	}
}

// Id of the repeating panel.
func RepeatPanelId(repeatPanelId int64) Option {
	return func(builder *Builder) error {

		builder.internal.RepeatPanelId = &repeatPanelId

		return nil
	}
}

// The maximum number of data points that the panel queries are retrieving.
func MaxDataPoints(maxDataPoints float64) Option {
	return func(builder *Builder) error {

		builder.internal.MaxDataPoints = &maxDataPoints

		return nil
	}
}

// List of transformations that are applied to the panel data before rendering.
// When there are multiple transformations, Grafana applies them in the order they are listed.
// Each transformation creates a result set that then passes on to the next transformation in the processing pipeline.
func Transformations(transformations []types.DataTransformerConfig) Option {
	return func(builder *Builder) error {

		builder.internal.Transformations = transformations

		return nil
	}
}

// The min time interval setting defines a lower limit for the $__interval and $__interval_ms variables.
// This value must be formatted as a number followed by a valid time
// identifier like: "40s", "3d", etc.
// See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
func Interval(interval string) Option {
	return func(builder *Builder) error {

		builder.internal.Interval = &interval

		return nil
	}
}

// Overrides the relative time range for individual panels,
// which causes them to be different than what is selected in
// the dashboard time picker in the top-right corner of the dashboard. You can use this to show metrics from different
// time periods or days on the same dashboard.
// The value is formatted as time operation like: `now-5m` (Last 5 minutes), `now/d` (the day so far),
// `now-5d/d`(Last 5 days), `now/w` (This week so far), `now-2y/y` (Last 2 years).
// Note: Panel time overrides have no effect when the dashboard’s time range is absolute.
// See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
func TimeFrom(timeFrom string) Option {
	return func(builder *Builder) error {

		builder.internal.TimeFrom = &timeFrom

		return nil
	}
}

// Overrides the time range for individual panels by shifting its start and end relative to the time picker.
// For example, you can shift the time range for the panel to be two hours earlier than the dashboard time picker setting `2h`.
// Note: Panel time overrides have no effect when the dashboard’s time range is absolute.
// See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
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

// It depends on the panel plugin. They are specified by the Options field in panel plugin schemas.
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
