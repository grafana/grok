package stat

// Defines values for BigValueColorMode.
const (
	BigValueColorModeBackground BigValueColorMode = "background"

	BigValueColorModeNone BigValueColorMode = "none"

	BigValueColorModeValue BigValueColorMode = "value"
)

// Defines values for BigValueGraphMode.
const (
	BigValueGraphModeArea BigValueGraphMode = "area"

	BigValueGraphModeLine BigValueGraphMode = "line"

	BigValueGraphModeNone BigValueGraphMode = "none"
)

// Defines values for BigValueJustifyMode.
const (
	BigValueJustifyModeAuto BigValueJustifyMode = "auto"

	BigValueJustifyModeCenter BigValueJustifyMode = "center"
)

// Defines values for BigValueTextMode.
const (
	BigValueTextModeAuto BigValueTextMode = "auto"

	BigValueTextModeName BigValueTextMode = "name"

	BigValueTextModeNone BigValueTextMode = "none"

	BigValueTextModeValue BigValueTextMode = "value"

	BigValueTextModeValueAndName BigValueTextMode = "value_and_name"
)

// Defines values for VizOrientation.
const (
	VizOrientationAuto VizOrientation = "auto"

	VizOrientationHorizontal VizOrientation = "horizontal"

	VizOrientationVertical VizOrientation = "vertical"
)

// TODO docs
type BigValueColorMode string

// TODO docs
type BigValueGraphMode string

// TODO docs
type BigValueJustifyMode string

// TODO docs
type BigValueTextMode string

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
	Limit *int32 `json:"limit,omitempty"`

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
	TitleSize *int32 `json:"titleSize,omitempty"`

	// Explicit value text size
	ValueSize *int32 `json:"valueSize,omitempty"`
}
