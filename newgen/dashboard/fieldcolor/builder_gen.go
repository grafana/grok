package dashboard

import "github.com/grafana/grok/newgen/dashboard/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.FieldColor
}

func Mode(mode types.FieldColorModeId) Option {
	return func(builder *Builder) error {

		builder.internal.Mode = mode

		return nil
	}
}

func FixedColor(fixedColor string) Option {
	return func(builder *Builder) error {

		builder.internal.FixedColor = &fixedColor

		return nil
	}
}

func SeriesBy(seriesBy types.FieldColorSeriesByMode) Option {
	return func(builder *Builder) error {

		builder.internal.SeriesBy = &seriesBy

		return nil
	}
}

func defaults() []Option {
	return []Option{}
}
