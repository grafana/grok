package timepicker

import (
	"encoding/json"

	"github.com/grafana/grok/newgen/dashboard/types"
)

type Option func(builder *Builder) error

type Builder struct {
	internal *types.TimePicker
}

func New(options ...Option) (Builder, error) {
	resource := &types.TimePicker{}
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

func (builder *Builder) Internal() *types.TimePicker {
	return builder.internal
}

// Whether timepicker is visible or not.
func Hidden(hidden bool) Option {
	return func(builder *Builder) error {

		builder.internal.Hidden = hidden

		return nil
	}
}

// Interval options available in the refresh picker dropdown.
func RefreshIntervals(refreshIntervals []string) Option {
	return func(builder *Builder) error {

		builder.internal.RefreshIntervals = refreshIntervals

		return nil
	}
}

// Whether timepicker is collapsed or not. Has no effect on provisioned dashboard.
func Collapse(collapse bool) Option {
	return func(builder *Builder) error {

		builder.internal.Collapse = collapse

		return nil
	}
}

// Whether timepicker is enabled or not. Has no effect on provisioned dashboard.
func Enable(enable bool) Option {
	return func(builder *Builder) error {

		builder.internal.Enable = enable

		return nil
	}
}

// Selectable options available in the time picker dropdown. Has no effect on provisioned dashboard.
func TimeOptions(timeOptions []string) Option {
	return func(builder *Builder) error {

		builder.internal.TimeOptions = timeOptions

		return nil
	}
}

func defaults() []Option {
	return []Option{
		Hidden(false),
		RefreshIntervals([]string{"5s", "10s", "30s", "1m", "5m", "15m", "30m", "1h", "2h", "1d"}),
		Collapse(false),
		Enable(true),
		TimeOptions([]string{"5m", "15m", "1h", "6h", "12h", "24h", "2d", "7d", "30d"}),
	}
}
