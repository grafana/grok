package barchart

// Defines values for AxisColorMode.
const (
	AxisColorModeSeries AxisColorMode = "series"

	AxisColorModeText AxisColorMode = "text"
)

// Defines values for AxisPlacement.
const (
	AxisPlacementAuto AxisPlacement = "auto"

	AxisPlacementBottom AxisPlacement = "bottom"

	AxisPlacementHidden AxisPlacement = "hidden"

	AxisPlacementLeft AxisPlacement = "left"

	AxisPlacementRight AxisPlacement = "right"

	AxisPlacementTop AxisPlacement = "top"
)

// Defines values for GraphGradientMode.
const (
	GraphGradientModeHue GraphGradientMode = "hue"

	GraphGradientModeNone GraphGradientMode = "none"

	GraphGradientModeOpacity GraphGradientMode = "opacity"

	GraphGradientModeScheme GraphGradientMode = "scheme"
)

// Defines values for GraphTresholdsStyleMode.
const (
	GraphTresholdsStyleModeArea GraphTresholdsStyleMode = "area"

	GraphTresholdsStyleModeDashed GraphTresholdsStyleMode = "dashed"

	GraphTresholdsStyleModeDashedArea GraphTresholdsStyleMode = "dashed+area"

	GraphTresholdsStyleModeLine GraphTresholdsStyleMode = "line"

	GraphTresholdsStyleModeLineArea GraphTresholdsStyleMode = "line+area"

	GraphTresholdsStyleModeOff GraphTresholdsStyleMode = "off"

	GraphTresholdsStyleModeSeries GraphTresholdsStyleMode = "series"
)

// Defines values for LegendDisplayMode.
const (
	LegendDisplayModeHidden LegendDisplayMode = "hidden"

	LegendDisplayModeList LegendDisplayMode = "list"

	LegendDisplayModeTable LegendDisplayMode = "table"
)

// Defines values for LegendPlacement.
const (
	LegendPlacementBottom LegendPlacement = "bottom"

	LegendPlacementRight LegendPlacement = "right"
)

// Defines values for ScaleDistribution.
const (
	ScaleDistributionLinear ScaleDistribution = "linear"

	ScaleDistributionLog ScaleDistribution = "log"

	ScaleDistributionOrdinal ScaleDistribution = "ordinal"

	ScaleDistributionSymlog ScaleDistribution = "symlog"
)

// Defines values for SortOrder.
const (
	SortOrderAsc SortOrder = "asc"

	SortOrderDesc SortOrder = "desc"

	SortOrderNone SortOrder = "none"
)

// Defines values for StackingMode.
const (
	StackingModeNone StackingMode = "none"

	StackingModeNormal StackingMode = "normal"

	StackingModePercent StackingMode = "percent"
)

// Defines values for TooltipDisplayMode.
const (
	TooltipDisplayModeMulti TooltipDisplayMode = "multi"

	TooltipDisplayModeNone TooltipDisplayMode = "none"

	TooltipDisplayModeSingle TooltipDisplayMode = "single"
)

// Defines values for VisibilityMode.
const (
	VisibilityModeAlways VisibilityMode = "always"

	VisibilityModeAuto VisibilityMode = "auto"

	VisibilityModeNever VisibilityMode = "never"
)

// Defines values for VizOrientation.
const (
	VizOrientationAuto VizOrientation = "auto"

	VizOrientationHorizontal VizOrientation = "horizontal"

	VizOrientationVertical VizOrientation = "vertical"
)

// TODO docs
type AxisColorMode string

// TODO docs
type AxisConfig struct {
	AxisCenteredZero *bool `json:"axisCenteredZero,omitempty"`

	// TODO docs
	AxisColorMode *AxisColorMode `json:"axisColorMode,omitempty"`
	AxisGridShow  *bool          `json:"axisGridShow,omitempty"`
	AxisLabel     *string        `json:"axisLabel,omitempty"`

	// TODO docs
	AxisPlacement *AxisPlacement `json:"axisPlacement,omitempty"`
	AxisSoftMax   *float32       `json:"axisSoftMax,omitempty"`
	AxisSoftMin   *float32       `json:"axisSoftMin,omitempty"`
	AxisWidth     *float32       `json:"axisWidth,omitempty"`

	// TODO docs
	ScaleDistribution *ScaleDistributionConfig `json:"scaleDistribution,omitempty"`
}

// TODO docs
type AxisPlacement string

// TODO docs
type GraphGradientMode string

// TODO docs
type GraphThresholdsStyleConfig struct {
	// TODO docs
	Mode GraphTresholdsStyleMode `json:"mode"`
}

// TODO docs
type GraphTresholdsStyleMode string

// TODO docs
type HideSeriesConfig struct {
	Legend  bool `json:"legend"`
	Tooltip bool `json:"tooltip"`
	Viz     bool `json:"viz"`
}

// TODO docs
type HideableFieldConfig struct {
	// TODO docs
	HideFrom *HideSeriesConfig `json:"hideFrom,omitempty"`
}

// TODO docs
// Note: "hidden" needs to remain as an option for plugins compatibility
type LegendDisplayMode string

// TODO docs
type LegendPlacement string

// TODO docs
type OptionsWithLegend struct {
	// TODO docs
	Legend VizLegendOptions `json:"legend"`
}

// TODO docs
type OptionsWithTextFormatting struct {
	// TODO docs
	Text *VizTextDisplayOptions `json:"text,omitempty"`
}

// TODO docs
type OptionsWithTooltip struct {
	// TODO docs
	Tooltip VizTooltipOptions `json:"tooltip"`
}

// PanelFieldConfig defines model for PanelFieldConfig.
type PanelFieldConfig struct {
	// Embedded struct due to allOf(#/components/schemas/AxisConfig)
	AxisConfig `yaml:",inline"`
	// Embedded struct due to allOf(#/components/schemas/HideableFieldConfig)
	HideableFieldConfig `yaml:",inline"`
}

// PanelOptions defines model for PanelOptions.
type PanelOptions struct {
	// Embedded struct due to allOf(#/components/schemas/OptionsWithLegend)
	OptionsWithLegend `yaml:",inline"`
	// Embedded struct due to allOf(#/components/schemas/OptionsWithTooltip)
	OptionsWithTooltip `yaml:",inline"`
	// Embedded struct due to allOf(#/components/schemas/OptionsWithTextFormatting)
	OptionsWithTextFormatting `yaml:",inline"`
	// Embedded fields due to inline allOf schema
}

// TODO docs
type ScaleDistribution string

// TODO docs
type ScaleDistributionConfig struct {
	LinearThreshold *int `json:"linearThreshold,omitempty"`
	Log             *int `json:"log,omitempty"`

	// TODO docs
	Type ScaleDistribution `json:"type"`
}

// TODO docs
type SortOrder string

// TODO docs
type StackingMode string

// TODO docs
type TooltipDisplayMode string

// TODO docs
type VisibilityMode string

// TODO docs
type VizLegendOptions struct {
	AsTable *bool    `json:"asTable,omitempty"`
	Calcs   []string `json:"calcs"`

	// TODO docs
	// Note: "hidden" needs to remain as an option for plugins compatibility
	DisplayMode LegendDisplayMode `json:"displayMode"`
	IsVisible   *bool             `json:"isVisible,omitempty"`

	// TODO docs
	Placement  LegendPlacement `json:"placement"`
	ShowLegend bool            `json:"showLegend"`
	SortBy     *string         `json:"sortBy,omitempty"`
	SortDesc   *bool           `json:"sortDesc,omitempty"`
	Width      *int32          `json:"width,omitempty"`
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

// TODO docs
type VizTooltipOptions struct {
	// TODO docs
	Mode TooltipDisplayMode `json:"mode"`

	// TODO docs
	Sort SortOrder `json:"sort"`
}
