package dashboard

import (
	"encoding/json"

	"github.com/grafana/grok/newgen/dashboard/annotationcontainer"
	"github.com/grafana/grok/newgen/dashboard/dashboardtemplating"
	"github.com/grafana/grok/newgen/dashboard/stringorbool"
	"github.com/grafana/grok/newgen/dashboard/timepicker"
	"github.com/grafana/grok/newgen/dashboard/types"
)

type Option func(builder *Builder) error

type Builder struct {
	internal *types.Dashboard
}

func New(title string, options ...Option) (Builder, error) {
	dashboard := &types.Dashboard{
		Title: &title,
	}

	builder := &Builder{internal: dashboard}

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

func (builder *Builder) Internal() *types.Dashboard {
	return builder.internal
}

// Unique numeric identifier for the dashboard.
// `id` is internal to a specific Grafana instance. `uid` should be used to identify a dashboard across Grafana instances.
func Id(id int64) Option {
	return func(builder *Builder) error {

		builder.internal.Id = &id

		return nil
	}
}

// Unique dashboard identifier that can be generated by anyone. string (8-40)
func Uid(uid string) Option {
	return func(builder *Builder) error {

		builder.internal.Uid = &uid

		return nil
	}
}

// Title of dashboard.
func Title(title string) Option {
	return func(builder *Builder) error {

		builder.internal.Title = &title

		return nil
	}
}

// Description of dashboard.
func Description(description string) Option {
	return func(builder *Builder) error {

		builder.internal.Description = &description

		return nil
	}
}

// This property should only be used in dashboards defined by plugins.  It is a quick check
// to see if the version has changed since the last time.
func Revision(revision int64) Option {
	return func(builder *Builder) error {

		builder.internal.Revision = &revision

		return nil
	}
}

// ID of a dashboard imported from the https://grafana.com/grafana/dashboards/ portal
func GnetId(gnetId string) Option {
	return func(builder *Builder) error {

		builder.internal.GnetId = &gnetId

		return nil
	}
}

// Tags associated with dashboard.
func Tags(tags []string) Option {
	return func(builder *Builder) error {

		builder.internal.Tags = tags

		return nil
	}
}

// Theme of dashboard.
func Style(style types.StyleEnum) Option {
	return func(builder *Builder) error {

		builder.internal.Style = style

		return nil
	}
}

// Timezone of dashboard. Accepted values are IANA TZDB zone ID or "browser" or "utc".
func Timezone(timezone string) Option {
	return func(builder *Builder) error {

		builder.internal.Timezone = &timezone

		return nil
	}
}
func Readonly() Option {
	return func(builder *Builder) error {
		builder.internal.Editable = true

		return nil
	}
}
func Editable() Option {
	return func(builder *Builder) error {
		builder.internal.Editable = false

		return nil
	}
}

// Configuration of dashboard cursor sync behavior.
// Accepted values are 0 (sync turned off), 1 (shared crosshair), 2 (shared crosshair and tooltip).
func Tooltip(tooltip types.DashboardCursorSync) Option {
	return func(builder *Builder) error {

		builder.internal.GraphTooltip = tooltip

		return nil
	}
}

// Time range for dashboard.
// Accepted values are relative time strings like {from: 'now-6h', to: 'now'} or absolute time strings like {from: '2020-07-10T08:00:00.000Z', to: '2020-07-10T14:00:00.000Z'}.
func Time(time struct {
	From string `json:"from"`
	To   string `json:"to"`
}) Option {
	return func(builder *Builder) error {

		builder.internal.Time = time

		return nil
	}
}

func Timepicker(opts ...timepicker.Option) Option {
	return func(builder *Builder) error {
		resource, err := timepicker.New(opts...)
		if err != nil {
			return err
		}

		builder.internal.Timepicker = resource.Internal()

		return nil
	}
}

// The month that the fiscal year starts on.  0 = January, 11 = December
func FiscalYearStartMonth(fiscalYearStartMonth uint8) Option {
	return func(builder *Builder) error {

		builder.internal.FiscalYearStartMonth = &fiscalYearStartMonth

		return nil
	}
}

// When set to true, the dashboard will redraw panels at an interval matching the pixel width.
// This will keep data "moving left" regardless of the query refresh rate. This setting helps
// avoid dashboards presenting stale live data
func LiveNow(liveNow bool) Option {
	return func(builder *Builder) error {

		builder.internal.LiveNow = &liveNow

		return nil
	}
}

// Day when the week starts. Expressed by the name of the day in lowercase, e.g. "monday".
func WeekStart(weekStart string) Option {
	return func(builder *Builder) error {

		builder.internal.WeekStart = &weekStart

		return nil
	}
}

func Refresh(opts ...stringorbool.Option) Option {
	return func(builder *Builder) error {
		resource, err := stringorbool.New(opts...)
		if err != nil {
			return err
		}

		builder.internal.Refresh = resource.Internal()

		return nil
	}
}

// Version of the dashboard, incremented each time the dashboard is updated.
func Version(version uint32) Option {
	return func(builder *Builder) error {

		builder.internal.Version = &version

		return nil
	}
}

// List of dashboard panels
func Panels(panels []types.RowPanel) Option {
	return func(builder *Builder) error {

		builder.internal.Panels = panels

		return nil
	}
}

func Templating(opts ...dashboardtemplating.Option) Option {
	return func(builder *Builder) error {
		resource, err := dashboardtemplating.New(opts...)
		if err != nil {
			return err
		}

		builder.internal.Templating = resource.Internal()

		return nil
	}
}

func Annotations(opts ...annotationcontainer.Option) Option {
	return func(builder *Builder) error {
		resource, err := annotationcontainer.New(opts...)
		if err != nil {
			return err
		}

		builder.internal.Annotations = resource.Internal()

		return nil
	}
}

// Links with references to other dashboards or external websites.
func Links(links []types.DashboardLink) Option {
	return func(builder *Builder) error {

		builder.internal.Links = links

		return nil
	}
}

func defaults() []Option {
	return []Option{
		Style("dark"),
		Timezone("browser"),
		Tooltip(0),
		FiscalYearStartMonth(0),
	}
}
