package dashboard

import (
	"errors"

	"github.com/grafana/grok/newgen/dashboard/annotationcontainer"
	"github.com/grafana/grok/newgen/dashboard/dashboardtemplating"
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

func (builder *Builder) Internal() *types.Dashboard {
	return builder.internal
}

func Id(id int64) Option {
	return func(builder *Builder) error {

		builder.internal.Id = &id

		return nil
	}
}

func Uid(uid string) Option {
	return func(builder *Builder) error {

		builder.internal.Uid = &uid

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

func Revision(revision int64) Option {
	return func(builder *Builder) error {

		builder.internal.Revision = &revision

		return nil
	}
}

func GnetId(gnetId string) Option {
	return func(builder *Builder) error {

		builder.internal.GnetId = &gnetId

		return nil
	}
}

func Tags(tags []string) Option {
	return func(builder *Builder) error {

		builder.internal.Tags = tags

		return nil
	}
}

func Style(style types.DashboardStyle) Option {
	return func(builder *Builder) error {

		builder.internal.Style = style

		return nil
	}
}

func Timezone(timezone string) Option {
	return func(builder *Builder) error {

		builder.internal.Timezone = &timezone

		return nil
	}
}

func Editable(editable bool) Option {
	return func(builder *Builder) error {

		builder.internal.Editable = editable

		return nil
	}
}

func GraphTooltip(graphTooltip types.DashboardCursorSync) Option {
	return func(builder *Builder) error {

		builder.internal.GraphTooltip = graphTooltip

		return nil
	}
}

func Time(time struct {
	// Default: "now-6h"
	From string `json:"from"`
	// Default: "now"
	To string `json:"to"`
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

func FiscalYearStartMonth(fiscalYearStartMonth uint8) Option {
	return func(builder *Builder) error {
		if !(fiscalYearStartMonth < 12) {
			return errors.New("fiscalYearStartMonth must be < 12")
		}

		builder.internal.FiscalYearStartMonth = &fiscalYearStartMonth

		return nil
	}
}

func LiveNow(liveNow bool) Option {
	return func(builder *Builder) error {

		builder.internal.LiveNow = &liveNow

		return nil
	}
}

func WeekStart(weekStart string) Option {
	return func(builder *Builder) error {

		builder.internal.WeekStart = &weekStart

		return nil
	}
}

func Refresh(refresh types.StringOrBool) Option {
	return func(builder *Builder) error {

		builder.internal.Refresh = &refresh

		return nil
	}
}

func SchemaVersion(schemaVersion uint16) Option {
	return func(builder *Builder) error {

		builder.internal.SchemaVersion = schemaVersion

		return nil
	}
}

func Version(version uint32) Option {
	return func(builder *Builder) error {

		builder.internal.Version = &version

		return nil
	}
}

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
		Editable(true),
		GraphTooltip(0),
		FiscalYearStartMonth(0),
		SchemaVersion(36),
	}
}
