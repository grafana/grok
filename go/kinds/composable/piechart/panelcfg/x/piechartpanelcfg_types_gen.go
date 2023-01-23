package piechartpanelcfg

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

// Defines values for PieChartLabels.
const (
	PieChartLabelsName PieChartLabels = "name"

	PieChartLabelsPercent PieChartLabels = "percent"

	PieChartLabelsValue PieChartLabels = "value"
)

// Defines values for PieChartLegendValues.
const (
	PieChartLegendValuesPercent PieChartLegendValues = "percent"

	PieChartLegendValuesValue PieChartLegendValues = "value"
)

// Defines values for PieChartType.
const (
	PieChartTypeDonut PieChartType = "donut"

	PieChartTypePie PieChartType = "pie"
)

// Defines values for SortOrder.
const (
	SortOrderAsc SortOrder = "asc"

	SortOrderDesc SortOrder = "desc"

	SortOrderNone SortOrder = "none"
)

// Defines values for TooltipDisplayMode.
const (
	TooltipDisplayModeMulti TooltipDisplayMode = "multi"

	TooltipDisplayModeNone TooltipDisplayMode = "none"

	TooltipDisplayModeSingle TooltipDisplayMode = "single"
)

// Defines values for VizOrientation.
const (
	VizOrientationAuto VizOrientation = "auto"

	VizOrientationHorizontal VizOrientation = "horizontal"

	VizOrientationVertical VizOrientation = "vertical"
)

// TODO docs
type HideSeriesConfig struct {
	Legend  bool `json:"legend"`
	Tooltip bool `json:"tooltip"`
	Viz     bool `json:"viz"`
}

// TODO docs
// Note: "hidden" needs to remain as an option for plugins compatibility
type LegendDisplayMode string

// TODO docs
type LegendPlacement string

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
	// TODO docs
	HideFrom *HideSeriesConfig `json:"hideFrom,omitempty"`
}

// PanelOptions defines model for PanelOptions.
type PanelOptions struct {
	// Embedded struct due to allOf(#/components/schemas/OptionsWithTooltip)
	OptionsWithTooltip `yaml:",inline"`
	// Embedded struct due to allOf(#/components/schemas/SingleStatBaseOptions)
	SingleStatBaseOptions `yaml:",inline"`
	// Embedded fields due to inline allOf schema
}

// Select labels to display on the pie chart.
//   - Name - The series or field name.
//   - Percent - The percentage of the whole.
//   - Value - The raw numerical value.
type PieChartLabels string

// PieChartLegendOptions defines model for PieChartLegendOptions.
type PieChartLegendOptions struct {
	// Embedded struct due to allOf(#/components/schemas/VizLegendOptions)
	VizLegendOptions `yaml:",inline"`
	// Embedded fields due to inline allOf schema
}

// Select values to display in the legend.
//   - Percent: The percentage of the whole.
//   - Value: The raw numerical value.
type PieChartLegendValues string

// Select the pie chart display style.
type PieChartType string

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
type SortOrder string

// TODO docs
type TooltipDisplayMode string

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
type VizOrientation string

// TODO docs
type VizTextDisplayOptions struct {
	// Explicit title text size
	TitleSize *float32 `json:"titleSize,omitempty"`

	// Explicit value text size
	ValueSize *float32 `json:"valueSize,omitempty"`
}

// TODO docs
type VizTooltipOptions struct {
	// TODO docs
	Mode TooltipDisplayMode `json:"mode"`

	// TODO docs
	Sort SortOrder `json:"sort"`
}
