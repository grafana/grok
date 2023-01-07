package bargauge

// Defines values for BarGaugeDisplayMode.
const (
	BarGaugeDisplayModeBasic BarGaugeDisplayMode = "basic"

	BarGaugeDisplayModeGradient BarGaugeDisplayMode = "gradient"

	BarGaugeDisplayModeLcd BarGaugeDisplayMode = "lcd"
)

// Defines values for VizOrientation.
const (
	VizOrientationAuto VizOrientation = "auto"

	VizOrientationHorizontal VizOrientation = "horizontal"

	VizOrientationVertical VizOrientation = "vertical"
)

// TODO docs
type BarGaugeDisplayMode string

// TODO docs
type OptionsWithTextFormatting struct {
	// TODO docs
	Text *VizTextDisplayOptions `json:"text,omitempty"`
}

// PanelOptions defines model for PanelOptions.
type PanelOptions struct {
	// Embedded struct due to allOf(#/components/schemas/SingleStatBaseOptions)
	SingleStatBaseOptions `yaml:",inline"`
	// Embedded fields due to inline allOf schema
}

// TODO docs
type ReduceDataOptions struct {
	// When !values, pick one value for the whole field
	Calcs []string `json:"calcs"`

	// Which fields to show.  By default this is only numeric fields
	Fields *string `json:"fields,omitempty"`

	// if showing all values limit
	Limit *float32 `json:"limit,omitempty"`

	// If true show each row value
	Values *bool `json:"values,omitempty"`
}

// SingleStatBaseOptions defines model for SingleStatBaseOptions.
type SingleStatBaseOptions struct {
	// Embedded struct due to allOf(#/components/schemas/OptionsWithTextFormatting)
	OptionsWithTextFormatting `yaml:",inline"`
	// Embedded fields due to inline allOf schema
}

// TODO docs
type VizOrientation string

// TODO docs
type VizTextDisplayOptions struct {
	// Explicit title text size
	TitleSize *float32 `json:"titleSize,omitempty"`

	// Explicit value text size
	ValueSize *float32 `json:"valueSize,omitempty"`
}
