package dashboard

import (
	"encoding/json"
	"fmt"
)

// Defines values for CursorSync.
const (
	CursorSyncN0 CursorSync = 0

	CursorSyncN1 CursorSync = 1

	CursorSyncN2 CursorSync = 2
)

// Defines values for LinkType.
const (
	LinkTypeDashboards LinkType = "dashboards"

	LinkTypeLink LinkType = "link"
)

// Defines values for FieldColorModeId.
const (
	FieldColorModeIdContinuousGrYlRd FieldColorModeId = "continuous-GrYlRd"

	FieldColorModeIdFixed FieldColorModeId = "fixed"

	FieldColorModeIdPaletteClassic FieldColorModeId = "palette-classic"

	FieldColorModeIdPaletteSaturated FieldColorModeId = "palette-saturated"

	FieldColorModeIdThresholds FieldColorModeId = "thresholds"
)

// Defines values for FieldColorSeriesByMode.
const (
	FieldColorSeriesByModeLast FieldColorSeriesByMode = "last"

	FieldColorSeriesByModeMax FieldColorSeriesByMode = "max"

	FieldColorSeriesByModeMin FieldColorSeriesByMode = "min"
)

// Defines values for GraphPanelType.
const (
	GraphPanelTypeGraph GraphPanelType = "graph"
)

// Defines values for HeatmapPanelType.
const (
	HeatmapPanelTypeHeatmap HeatmapPanelType = "heatmap"
)

// Defines values for MappingType.
const (
	MappingTypeRange MappingType = "range"

	MappingTypeRegex MappingType = "regex"

	MappingTypeSpecial MappingType = "special"

	MappingTypeValue MappingType = "value"
)

// Defines values for PanelRepeatDirection.
const (
	PanelRepeatDirectionH PanelRepeatDirection = "h"

	PanelRepeatDirectionV PanelRepeatDirection = "v"
)

// Defines values for RowPanelType.
const (
	RowPanelTypeRow RowPanelType = "row"
)

// Defines values for SpecialValueMapOptionsMatch.
const (
	SpecialValueMapOptionsMatchFalse SpecialValueMapOptionsMatch = "false"

	SpecialValueMapOptionsMatchTrue SpecialValueMapOptionsMatch = "true"
)

// Defines values for SpecialValueMatch.
const (
	SpecialValueMatchEmpty SpecialValueMatch = "empty"

	SpecialValueMatchFalse SpecialValueMatch = "false"

	SpecialValueMatchNan SpecialValueMatch = "nan"

	SpecialValueMatchNull SpecialValueMatch = "null"

	SpecialValueMatchNullNan SpecialValueMatch = "null+nan"

	SpecialValueMatchTrue SpecialValueMatch = "true"
)

// Defines values for ThresholdsMode.
const (
	ThresholdsModeAbsolute ThresholdsMode = "absolute"

	ThresholdsModePercentage ThresholdsMode = "percentage"
)

// Defines values for VariableType.
const (
	VariableTypeAdhoc VariableType = "adhoc"

	VariableTypeConstant VariableType = "constant"

	VariableTypeCustom VariableType = "custom"

	VariableTypeDatasource VariableType = "datasource"

	VariableTypeInterval VariableType = "interval"

	VariableTypeQuery VariableType = "query"

	VariableTypeSystem VariableType = "system"

	VariableTypeTextbox VariableType = "textbox"
)

// Defines values for Style.
const (
	StyleDark Style = "dark"

	StyleLight Style = "light"
)

// Defines values for Timezone.
const (
	TimezoneBrowser Timezone = "browser"

	TimezoneEmpty Timezone = ""

	TimezoneUtc Timezone = "utc"
)

// TODO docs
// FROM: AnnotationQuery in grafana-data/src/types/annotations.ts
type AnnotationQuery struct {
	BuiltIn int `json:"builtIn"`

	// Datasource to use for annotation.
	Datasource struct {
		Type *string `json:"type,omitempty"`
		Uid  *string `json:"uid,omitempty"`
	} `json:"datasource"`

	// Whether annotation is enabled.
	Enable bool `json:"enable"`

	// Whether to hide annotation.
	Hide *bool `json:"hide,omitempty"`

	// Annotation icon color.
	IconColor *string `json:"iconColor,omitempty"`

	// Name of annotation.
	Name *string `json:"name,omitempty"`

	// Query for annotation data.
	RawQuery *string `json:"rawQuery,omitempty"`
	ShowIn   int     `json:"showIn"`

	// TODO docs
	Target *AnnotationTarget `json:"target,omitempty"`
	Type   string            `json:"type"`
}

// TODO docs
type AnnotationTarget struct {
	Limit    int64    `json:"limit"`
	MatchAny bool     `json:"matchAny"`
	Tags     []string `json:"tags"`
	Type     string   `json:"type"`
}

// 0 for no shared crosshair or tooltip (default).
// 1 for shared crosshair.
// 2 for shared crosshair AND shared tooltip.
type CursorSync int

// FROM public/app/features/dashboard/state/Models.ts - ish
// TODO docs
type Link struct {
	AsDropdown  bool     `json:"asDropdown"`
	Icon        *string  `json:"icon,omitempty"`
	IncludeVars bool     `json:"includeVars"`
	KeepTime    bool     `json:"keepTime"`
	Tags        []string `json:"tags"`
	TargetBlank bool     `json:"targetBlank"`
	Title       string   `json:"title"`
	Tooltip     *string  `json:"tooltip,omitempty"`

	// TODO docs
	Type LinkType `json:"type"`
	Url  *string  `json:"url,omitempty"`
}

// TODO docs
type LinkType string

// DynamicConfigValue defines model for DynamicConfigValue.
type DynamicConfigValue struct {
	Id    string       `json:"id"`
	Value *interface{} `json:"value,omitempty"`
}

// TODO docs
type FieldColor struct {
	// Stores the fixed color value if mode is fixed
	FixedColor *string `json:"fixedColor,omitempty"`

	// The main color scheme mode
	Mode interface{} `json:"mode"`

	// TODO docs
	SeriesBy *FieldColorSeriesByMode `json:"seriesBy,omitempty"`
}

// TODO docs
type FieldColorModeId string

// TODO docs
type FieldColorSeriesByMode string

// FieldConfig defines model for FieldConfig.
type FieldConfig struct {
	// TODO docs
	Color *FieldColor `json:"color,omitempty"`

	// custom is specified by the PanelFieldConfig field
	// in panel plugin schemas.
	Custom *map[string]interface{} `json:"custom,omitempty"`

	// Significant digits (for display)
	Decimals *float32 `json:"decimals,omitempty"`

	// Human readable field metadata
	Description *string `json:"description,omitempty"`

	// The display value for this field.  This supports template variables blank is auto
	DisplayName *string `json:"displayName,omitempty"`

	// This can be used by data sources that return and explicit naming structure for values and labels
	// When this property is configured, this value is used rather than the default naming strategy.
	DisplayNameFromDS *string `json:"displayNameFromDS,omitempty"`

	// True if data source field supports ad-hoc filters
	Filterable *bool `json:"filterable,omitempty"`

	// The behavior when clicking on a result
	Links *[]interface{} `json:"links,omitempty"`

	// Convert input values into a display string
	Mappings *[]ValueMapping `json:"mappings,omitempty"`
	Max      *float32        `json:"max,omitempty"`
	Min      *float32        `json:"min,omitempty"`

	// Alternative to empty string
	NoValue *string `json:"noValue,omitempty"`

	// An explicit path to the field in the datasource.  When the frame meta includes a path,
	// This will default to `${frame.meta.path}/${field.name}
	//
	// When defined, this value can be used as an identifier within the datasource scope, and
	// may be used to update the results
	Path       *string           `json:"path,omitempty"`
	Thresholds *ThresholdsConfig `json:"thresholds,omitempty"`

	// Numeric Options
	Unit *string `json:"unit,omitempty"`

	// True if data source can write a value to the path.  Auth/authz are supported separately
	Writeable *bool `json:"writeable,omitempty"`
}

// FieldConfigSource defines model for FieldConfigSource.
type FieldConfigSource struct {
	Defaults  FieldConfig `json:"defaults"`
	Overrides []struct {
		Matcher    MatcherConfig        `json:"matcher"`
		Properties []DynamicConfigValue `json:"properties"`
	} `json:"overrides"`
}

// Support for legacy graph and heatmap panels.
type GraphPanel struct {
	Type GraphPanelType `json:"type"`
}

// GraphPanelType defines model for GraphPanel.Type.
type GraphPanelType string

// GridPos defines model for GridPos.
type GridPos struct {
	// Panel
	H int `json:"h"`

	// true if fixed
	Static *bool `json:"static,omitempty"`

	// Panel
	W int `json:"w"`

	// Panel x
	X int `json:"x"`

	// Panel y
	Y int `json:"y"`
}

// HeatmapPanel defines model for HeatmapPanel.
type HeatmapPanel struct {
	Type HeatmapPanelType `json:"type"`
}

// HeatmapPanelType defines model for HeatmapPanel.Type.
type HeatmapPanelType string

// TODO docs
type MappingType string

// MatcherConfig defines model for MatcherConfig.
type MatcherConfig struct {
	Id      string       `json:"id"`
	Options *interface{} `json:"options,omitempty"`
}

// Dashboard panels. Panels are canonically defined inline
// because they share a version timeline with the dashboard
// schema; they do not evolve independently.
type Panel struct {
	// The datasource used in all targets.
	Datasource *struct {
		Type *string `json:"type,omitempty"`
		Uid  *string `json:"uid,omitempty"`
	} `json:"datasource,omitempty"`

	// Description.
	Description *string           `json:"description,omitempty"`
	FieldConfig FieldConfigSource `json:"fieldConfig"`
	GridPos     *GridPos          `json:"gridPos,omitempty"`

	// TODO docs
	Id *int `json:"id,omitempty"`

	// TODO docs
	// TODO tighter constraint
	Interval *string `json:"interval,omitempty"`

	// Panel links.
	// TODO fill this out - seems there are a couple variants?
	Links *[]Link `json:"links,omitempty"`

	// TODO docs
	MaxDataPoints *float32 `json:"maxDataPoints,omitempty"`

	// options is specified by the PanelOptions field in panel
	// plugin schemas.
	Options map[string]interface{} `json:"options"`

	// FIXME this almost certainly has to be changed in favor of scuemata versions
	PluginVersion *string `json:"pluginVersion,omitempty"`

	// Name of template variable to repeat for.
	Repeat *string `json:"repeat,omitempty"`

	// Direction to repeat in if 'repeat' is set.
	// "h" for horizontal, "v" for vertical.
	RepeatDirection PanelRepeatDirection `json:"repeatDirection"`

	// TODO docs
	Tags *[]string `json:"tags,omitempty"`

	// TODO docs
	Targets *[]Target `json:"targets,omitempty"`

	// TODO docs - seems to be an old field from old dashboard alerts?
	Thresholds *[]interface{} `json:"thresholds,omitempty"`

	// TODO docs
	// TODO tighter constraint
	TimeFrom *string `json:"timeFrom,omitempty"`

	// TODO docs
	TimeRegions *[]interface{} `json:"timeRegions,omitempty"`

	// TODO docs
	// TODO tighter constraint
	TimeShift *string `json:"timeShift,omitempty"`

	// Panel title.
	Title           *string          `json:"title,omitempty"`
	Transformations []Transformation `json:"transformations"`

	// Whether to display the panel without a background.
	Transparent bool `json:"transparent"`

	// The panel plugin type id. May not be empty.
	Type string `json:"type"`
}

// Direction to repeat in if 'repeat' is set.
// "h" for horizontal, "v" for vertical.
type PanelRepeatDirection string

// TODO docs
type RangeMap struct {
	Options struct {
		// to and from are `number | null` in current ts, really not sure what to do
		From float64 `json:"from"`

		// TODO docs
		Result ValueMappingResult `json:"result"`
		To     float64            `json:"to"`
	} `json:"options"`
	Type struct {
		// Embedded struct due to allOf(#/components/schemas/MappingType)
		MappingType `yaml:",inline"`
		// Embedded fields due to inline allOf schema
	} `json:"type"`
}

// TODO docs
type RegexMap struct {
	Options struct {
		Pattern string `json:"pattern"`

		// TODO docs
		Result ValueMappingResult `json:"result"`
	} `json:"options"`
	Type struct {
		// Embedded struct due to allOf(#/components/schemas/MappingType)
		MappingType `yaml:",inline"`
		// Embedded fields due to inline allOf schema
	} `json:"type"`
}

// Row panel
type RowPanel struct {
	Collapsed bool `json:"collapsed"`

	// Name of default datasource.
	Datasource *struct {
		Type *string `json:"type,omitempty"`
		Uid  *string `json:"uid,omitempty"`
	} `json:"datasource,omitempty"`
	GridPos *GridPos      `json:"gridPos,omitempty"`
	Id      int           `json:"id"`
	Panels  []interface{} `json:"panels"`

	// Name of template variable to repeat for.
	Repeat *string      `json:"repeat,omitempty"`
	Title  *string      `json:"title,omitempty"`
	Type   RowPanelType `json:"type"`
}

// RowPanelType defines model for RowPanel.Type.
type RowPanelType string

// TODO docs
type SpecialValueMap struct {
	Options struct {
		Match   SpecialValueMapOptionsMatch `json:"match"`
		Pattern string                      `json:"pattern"`

		// TODO docs
		Result ValueMappingResult `json:"result"`
	} `json:"options"`
	Type struct {
		// Embedded struct due to allOf(#/components/schemas/MappingType)
		MappingType `yaml:",inline"`
		// Embedded fields due to inline allOf schema
	} `json:"type"`
}

// SpecialValueMapOptionsMatch defines model for SpecialValueMap.Options.Match.
type SpecialValueMapOptionsMatch string

// TODO docs
type SpecialValueMatch string

// Schema for panel targets is specified by datasource
// plugins. We use a placeholder definition, which the Go
// schema loader either left open/as-is with the Base
// variant of the Dashboard and Panel families, or filled
// with types derived from plugins in the Instance variant.
// When working directly from CUE, importers can extend this
// type directly to achieve the same effect.
type Target map[string]interface{}

// TODO docs
type Threshold struct {
	// TODO docs
	Color string `json:"color"`

	// TODO docs
	// TODO are the values here enumerable into a disjunction?
	// Some seem to be listed in typescript comment
	State *string `json:"state,omitempty"`

	// TODO docs
	// FIXME the corresponding typescript field is required/non-optional, but nulls currently appear here when serializing -Infinity to JSON
	Value *float32 `json:"value,omitempty"`
}

// ThresholdsConfig defines model for ThresholdsConfig.
type ThresholdsConfig struct {
	Mode ThresholdsMode `json:"mode"`

	// Must be sorted by 'value', first value is always -Infinity
	Steps []Threshold `json:"steps"`
}

// ThresholdsMode defines model for ThresholdsMode.
type ThresholdsMode string

// TODO docs
// FIXME this is extremely underspecfied; wasn't obvious which typescript types corresponded to it
type Transformation struct {
	Id      string                 `json:"id"`
	Options map[string]interface{} `json:"options"`
}

// TODO docs
type ValueMap struct {
	Options ValueMap_Options `json:"options"`
	Type    struct {
		// Embedded struct due to allOf(#/components/schemas/MappingType)
		MappingType `yaml:",inline"`
		// Embedded fields due to inline allOf schema
	} `json:"type"`
}

// ValueMap_Options defines model for ValueMap.Options.
type ValueMap_Options struct {
	AdditionalProperties map[string]ValueMappingResult `json:"-"`
}

// TODO docs
type ValueMapping interface{}

// TODO docs
type ValueMappingResult struct {
	Color *string `json:"color,omitempty"`
	Icon  *string `json:"icon,omitempty"`
	Index *int32  `json:"index,omitempty"`
	Text  *string `json:"text,omitempty"`
}

// FROM: packages/grafana-data/src/types/templateVars.ts
// TODO docs
// TODO what about what's in public/app/features/types.ts?
// TODO there appear to be a lot of different kinds of [template] vars here? if so need a disjunction
type VariableModel struct {
	Label *string `json:"label,omitempty"`
	Name  string  `json:"name"`

	// FROM: packages/grafana-data/src/types/templateVars.ts
	// TODO docs
	// TODO this implies some wider pattern/discriminated union, probably?
	Type VariableType `json:"type"`
}

// FROM: packages/grafana-data/src/types/templateVars.ts
// TODO docs
// TODO this implies some wider pattern/discriminated union, probably?
type VariableType string

// Dashboard defines model for dashboard.
type Dashboard struct {
	// TODO docs
	Annotations *struct {
		List []AnnotationQuery `json:"list"`
	} `json:"annotations,omitempty"`

	// Description of dashboard.
	Description *string `json:"description,omitempty"`

	// Whether a dashboard is editable or not.
	Editable bool `json:"editable"`

	// TODO docs
	FiscalYearStartMonth *int    `json:"fiscalYearStartMonth,omitempty"`
	GnetId               *string `json:"gnetId,omitempty"`

	// 0 for no shared crosshair or tooltip (default).
	// 1 for shared crosshair.
	// 2 for shared crosshair AND shared tooltip.
	GraphTooltip CursorSync `json:"graphTooltip"`

	// Unique numeric identifier for the dashboard.
	// TODO must isolate or remove identifiers local to a Grafana instance...?
	Id *int64 `json:"id,omitempty"`

	// TODO docs
	Links *[]Link `json:"links,omitempty"`

	// TODO docs
	LiveNow *bool          `json:"liveNow,omitempty"`
	Panels  *[]interface{} `json:"panels,omitempty"`

	// TODO docs
	Refresh *interface{} `json:"refresh,omitempty"`

	// Version of the JSON schema, incremented each time a Grafana update brings
	// changes to said schema.
	// TODO this is the existing schema numbering system. It will be replaced by Thema's themaVersion
	SchemaVersion int `json:"schemaVersion"`

	// Theme of dashboard.
	Style Style `json:"style"`

	// Tags associated with dashboard.
	Tags *[]string `json:"tags,omitempty"`

	// TODO docs
	Templating *struct {
		List []VariableModel `json:"list"`
	} `json:"templating,omitempty"`

	// Time range for dashboard, e.g. last 6 hours, last 7 days, etc
	Time *struct {
		From string `json:"from"`
		To   string `json:"to"`
	} `json:"time,omitempty"`

	// TODO docs
	// TODO this appears to be spread all over in the frontend. Concepts will likely need tidying in tandem with schema changes
	Timepicker *struct {
		// Whether timepicker is collapsed or not.
		Collapse bool `json:"collapse"`

		// Whether timepicker is enabled or not.
		Enable bool `json:"enable"`

		// Whether timepicker is visible or not.
		Hidden bool `json:"hidden"`

		// Selectable intervals for auto-refresh.
		RefreshIntervals []string `json:"refresh_intervals"`

		// TODO docs
		TimeOptions []string `json:"time_options"`
	} `json:"timepicker,omitempty"`

	// Timezone of dashboard,
	Timezone *Timezone `json:"timezone,omitempty"`

	// Title of dashboard.
	Title *string `json:"title,omitempty"`

	// Unique dashboard identifier that can be generated by anyone. string (8-40)
	Uid *string `json:"uid,omitempty"`

	// Version of the dashboard, incremented each time the dashboard is updated.
	Version *int `json:"version,omitempty"`

	// TODO docs
	WeekStart *string `json:"weekStart,omitempty"`
}

// Theme of dashboard.
type Style string

// Timezone of dashboard,
type Timezone string

// Getter for additional properties for ValueMap_Options. Returns the specified
// element and whether it was found
func (a ValueMap_Options) Get(fieldName string) (value ValueMappingResult, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for ValueMap_Options
func (a *ValueMap_Options) Set(fieldName string, value ValueMappingResult) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]ValueMappingResult)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for ValueMap_Options to handle AdditionalProperties
func (a *ValueMap_Options) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]ValueMappingResult)
		for fieldName, fieldBuf := range object {
			var fieldVal ValueMappingResult
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for ValueMap_Options to handle AdditionalProperties
func (a ValueMap_Options) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}
