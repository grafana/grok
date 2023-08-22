package dashboard

import (
	"errors"

	"github.com/grafana/grok/newgen/dashboard/types"
)

type Option func(builder *Builder) error

type Builder struct {
	internal *types.Panel
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

func Datasource(datasource types.DataSourceRef) Option {
	return func(builder *Builder) error {

		builder.internal.Datasource = &datasource

		return nil
	}
}

func GridPos(gridPos types.GridPos) Option {
	return func(builder *Builder) error {

		builder.internal.GridPos = &gridPos

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

func LibraryPanel(libraryPanel types.LibraryPanelRef) Option {
	return func(builder *Builder) error {

		builder.internal.LibraryPanel = &libraryPanel

		return nil
	}
}

func Options(options any) Option {
	return func(builder *Builder) error {

		builder.internal.Options = options

		return nil
	}
}

func FieldConfig(fieldConfig types.FieldConfigSource) Option {
	return func(builder *Builder) error {

		builder.internal.FieldConfig = fieldConfig

		return nil
	}
}

func defaults() []Option {
	return []Option{
		Transparent(false),
	}
}
