// Code generated - EDITING IS FUTILE. DO NOT EDIT.
//
// Generated by pipeline:
//     go
// Using jennies:
//     GoTypesJenny
//     ComposableLatestMajorsOrXJenny
//
// Run 'go generate ./' from repository root to regenerate.

package timeseriespanelcfg

// Defines values for AxisColorMode.
const (
	AxisColorModeSeries AxisColorMode = "series"
	AxisColorModeText   AxisColorMode = "text"
)

// Defines values for AxisPlacement.
const (
	AxisPlacementAuto   AxisPlacement = "auto"
	AxisPlacementBottom AxisPlacement = "bottom"
	AxisPlacementHidden AxisPlacement = "hidden"
	AxisPlacementLeft   AxisPlacement = "left"
	AxisPlacementRight  AxisPlacement = "right"
	AxisPlacementTop    AxisPlacement = "top"
)

// Defines values for BarAlignment.
const (
	BarAlignmentMinus1 BarAlignment = -1
	BarAlignmentN0     BarAlignment = 0
	BarAlignmentN1     BarAlignment = 1
)

// Defines values for GraphDrawStyle.
const (
	GraphDrawStyleBars   GraphDrawStyle = "bars"
	GraphDrawStyleLine   GraphDrawStyle = "line"
	GraphDrawStylePoints GraphDrawStyle = "points"
)

// Defines values for GraphGradientMode.
const (
	GraphGradientModeHue     GraphGradientMode = "hue"
	GraphGradientModeNone    GraphGradientMode = "none"
	GraphGradientModeOpacity GraphGradientMode = "opacity"
	GraphGradientModeScheme  GraphGradientMode = "scheme"
)

// Defines values for GraphTransform.
const (
	GraphTransformConstant  GraphTransform = "constant"
	GraphTransformNegativeY GraphTransform = "negative-Y"
)

// Defines values for GraphTresholdsStyleMode.
const (
	GraphTresholdsStyleModeArea       GraphTresholdsStyleMode = "area"
	GraphTresholdsStyleModeDashed     GraphTresholdsStyleMode = "dashed"
	GraphTresholdsStyleModeDashedArea GraphTresholdsStyleMode = "dashed+area"
	GraphTresholdsStyleModeLine       GraphTresholdsStyleMode = "line"
	GraphTresholdsStyleModeLineArea   GraphTresholdsStyleMode = "line+area"
	GraphTresholdsStyleModeOff        GraphTresholdsStyleMode = "off"
	GraphTresholdsStyleModeSeries     GraphTresholdsStyleMode = "series"
)

// Defines values for LegendDisplayMode.
const (
	LegendDisplayModeHidden LegendDisplayMode = "hidden"
	LegendDisplayModeList   LegendDisplayMode = "list"
	LegendDisplayModeTable  LegendDisplayMode = "table"
)

// Defines values for LegendPlacement.
const (
	LegendPlacementBottom LegendPlacement = "bottom"
	LegendPlacementRight  LegendPlacement = "right"
)

// Defines values for LineInterpolation.
const (
	LineInterpolationLinear     LineInterpolation = "linear"
	LineInterpolationSmooth     LineInterpolation = "smooth"
	LineInterpolationStepAfter  LineInterpolation = "stepAfter"
	LineInterpolationStepBefore LineInterpolation = "stepBefore"
)

// Defines values for LineStyleFill.
const (
	LineStyleFillDash   LineStyleFill = "dash"
	LineStyleFillDot    LineStyleFill = "dot"
	LineStyleFillSolid  LineStyleFill = "solid"
	LineStyleFillSquare LineStyleFill = "square"
)

// Defines values for ScaleDistribution.
const (
	ScaleDistributionLinear  ScaleDistribution = "linear"
	ScaleDistributionLog     ScaleDistribution = "log"
	ScaleDistributionOrdinal ScaleDistribution = "ordinal"
	ScaleDistributionSymlog  ScaleDistribution = "symlog"
)

// Defines values for SortOrder.
const (
	SortOrderAsc  SortOrder = "asc"
	SortOrderDesc SortOrder = "desc"
	SortOrderNone SortOrder = "none"
)

// Defines values for StackingMode.
const (
	StackingModeNone    StackingMode = "none"
	StackingModeNormal  StackingMode = "normal"
	StackingModePercent StackingMode = "percent"
)

// Defines values for TooltipDisplayMode.
const (
	TooltipDisplayModeMulti  TooltipDisplayMode = "multi"
	TooltipDisplayModeNone   TooltipDisplayMode = "none"
	TooltipDisplayModeSingle TooltipDisplayMode = "single"
)

// Defines values for VisibilityMode.
const (
	VisibilityModeAlways VisibilityMode = "always"
	VisibilityModeAuto   VisibilityMode = "auto"
	VisibilityModeNever  VisibilityMode = "never"
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
type BarAlignment int

// TODO docs
type BarConfig struct {
	// TODO docs
	BarAlignment   *BarAlignment `json:"barAlignment,omitempty"`
	BarMaxWidth    *float32      `json:"barMaxWidth,omitempty"`
	BarWidthFactor *float32      `json:"barWidthFactor,omitempty"`
}

// TODO docs
type FillConfig struct {
	FillBelowTo *string  `json:"fillBelowTo,omitempty"`
	FillColor   *string  `json:"fillColor,omitempty"`
	FillOpacity *float32 `json:"fillOpacity,omitempty"`
}

// TODO docs
type GraphDrawStyle string

// GraphFieldConfig defines model for GraphFieldConfig.
type GraphFieldConfig struct {
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
	BarAlignment   *BarAlignment `json:"barAlignment,omitempty"`
	BarMaxWidth    *float32      `json:"barMaxWidth,omitempty"`
	BarWidthFactor *float32      `json:"barWidthFactor,omitempty"`
	FillBelowTo    *string       `json:"fillBelowTo,omitempty"`
	FillColor      *string       `json:"fillColor,omitempty"`
	FillOpacity    *float32      `json:"fillOpacity,omitempty"`

	// TODO docs
	HideFrom  *HideSeriesConfig `json:"hideFrom,omitempty"`
	LineColor *string           `json:"lineColor,omitempty"`

	// TODO docs
	LineInterpolation *LineInterpolation `json:"lineInterpolation,omitempty"`

	// TODO docs
	LineStyle   *LineStyle `json:"lineStyle,omitempty"`
	LineWidth   *float32   `json:"lineWidth,omitempty"`
	PointColor  *string    `json:"pointColor,omitempty"`
	PointSize   *float32   `json:"pointSize,omitempty"`
	PointSymbol *string    `json:"pointSymbol,omitempty"`

	// TODO docs
	ScaleDistribution *ScaleDistributionConfig `json:"scaleDistribution,omitempty"`

	// TODO docs
	ShowPoints *VisibilityMode `json:"showPoints,omitempty"`

	// Indicate if null values should be treated as gaps or connected.
	// When the value is a number, it represents the maximum delta in the
	// X axis that should be considered connected.  For timeseries, this is milliseconds
	SpanNulls *interface{} `json:"spanNulls,omitempty"`

	// TODO docs
	Stacking *StackingConfig `json:"stacking,omitempty"`
}

// TODO docs
type GraphGradientMode string

// TODO docs
type GraphThresholdsStyleConfig struct {
	// TODO docs
	Mode GraphTresholdsStyleMode `json:"mode"`
}

// TODO docs
type GraphTransform string

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
type LineConfig struct {
	LineColor *string `json:"lineColor,omitempty"`

	// TODO docs
	LineInterpolation *LineInterpolation `json:"lineInterpolation,omitempty"`

	// TODO docs
	LineStyle *LineStyle `json:"lineStyle,omitempty"`
	LineWidth *float32   `json:"lineWidth,omitempty"`

	// Indicate if null values should be treated as gaps or connected.
	// When the value is a number, it represents the maximum delta in the
	// X axis that should be considered connected.  For timeseries, this is milliseconds
	SpanNulls *interface{} `json:"spanNulls,omitempty"`
}

// TODO docs
type LineInterpolation string

// TODO docs
type LineStyle struct {
	Dash []float32      `json:"dash,omitempty"`
	Fill *LineStyleFill `json:"fill,omitempty"`
}

// LineStyleFill defines model for LineStyle.Fill.
type LineStyleFill string

// TODO docs
type PanelFieldConfig = GraphFieldConfig

// PanelOptions defines model for PanelOptions.
type PanelOptions struct {
	// TODO docs
	Legend VizLegendOptions `json:"legend"`

	// TODO docs
	Tooltip VizTooltipOptions `json:"tooltip"`
}

// TODO docs
type PointsConfig struct {
	PointColor  *string  `json:"pointColor,omitempty"`
	PointSize   *float32 `json:"pointSize,omitempty"`
	PointSymbol *string  `json:"pointSymbol,omitempty"`

	// TODO docs
	ShowPoints *VisibilityMode `json:"showPoints,omitempty"`
}

// TODO docs
type ScaleDistribution string

// TODO docs
type ScaleDistributionConfig struct {
	LinearThreshold *float32 `json:"linearThreshold,omitempty"`
	Log             *float32 `json:"log,omitempty"`

	// TODO docs
	Type ScaleDistribution `json:"type"`
}

// TODO docs
type SortOrder string

// TODO docs
type StackableFieldConfig struct {
	// TODO docs
	Stacking *StackingConfig `json:"stacking,omitempty"`
}

// TODO docs
type StackingConfig struct {
	Group *string `json:"group,omitempty"`

	// TODO docs
	Mode *StackingMode `json:"mode,omitempty"`
}

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
	Width      *float32        `json:"width,omitempty"`
}

// TODO docs
type VizTooltipOptions struct {
	// TODO docs
	Mode TooltipDisplayMode `json:"mode"`

	// TODO docs
	Sort SortOrder `json:"sort"`
}
