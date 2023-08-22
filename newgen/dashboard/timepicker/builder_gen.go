package dashboard

import "github.com/grafana/grok/newgen/dashboard/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.TimePicker
}

func Hidden(hidden bool) Option {
	return func(builder *Builder) error {

		builder.internal.Hidden = hidden

		return nil
	}
}

func Refresh_intervals(refresh_intervals []string) Option {
	return func(builder *Builder) error {

		builder.internal.Refresh_intervals = refresh_intervals

		return nil
	}
}

func Collapse(collapse bool) Option {
	return func(builder *Builder) error {

		builder.internal.Collapse = collapse

		return nil
	}
}

func Enable(enable bool) Option {
	return func(builder *Builder) error {

		builder.internal.Enable = enable

		return nil
	}
}

func Time_options(time_options []string) Option {
	return func(builder *Builder) error {

		builder.internal.Time_options = time_options

		return nil
	}
}

func defaults() []Option {
	return []Option{
		Hidden(false),
		Refresh_intervals([]interface{}{"5s", "10s", "30s", "1m", "5m", "15m", "30m", "1h", "2h", "1d"}),
		Collapse(false),
		Enable(true),
		Time_options([]interface{}{"5m", "15m", "1h", "6h", "12h", "24h", "2d", "7d", "30d"}),
	}
}
