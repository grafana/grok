
export enum DashboardStyle {
	Light = "light",
	Dark = "dark",
}

export interface DashboardTemplating {
	// List of configured template variables with their saved values along with some other metadata
	list?: VariableModel[];
}

export interface TimePicker {
	// Whether timepicker is visible or not.
	hidden: boolean;
	// Interval options available in the refresh picker dropdown.
	refreshIntervals: string[];
	// Whether timepicker is collapsed or not. Has no effect on provisioned dashboard.
	collapse: boolean;
	// Whether timepicker is enabled or not. Has no effect on provisioned dashboard.
	enable: boolean;
	// Selectable options available in the time picker dropdown. Has no effect on provisioned dashboard.
	timeOptions: string[];
}

export interface TimeInterval {
	from: string;
	to: string;
}

// TODO: this should be a regular DataQuery that depends on the selected dashboard
// these match the properties of the "grafana" datasource that is default in most dashboards
export interface AnnotationTarget {
	// Only required/valid for the grafana datasource...
	// but code+tests is already depending on it so hard to change
	limit: number;
	// Only required/valid for the grafana datasource...
	// but code+tests is already depending on it so hard to change
	matchAny: boolean;
	// Only required/valid for the grafana datasource...
	// but code+tests is already depending on it so hard to change
	tags: string[];
	// Only required/valid for the grafana datasource...
	// but code+tests is already depending on it so hard to change
	type: string;
}

export interface AnnotationPanelFilter {
	// Should the specified panels be included or excluded
	exclude?: boolean;
	// Panel IDs that should be included or excluded
	ids: number[];
}

// Contains the list of annotations that are associated with the dashboard.
// Annotations are used to overlay event markers and overlay event tags on graphs.
// Grafana comes with a native annotation store and the ability to add annotation events directly from the graph panel or via the HTTP API.
// See https://grafana.com/docs/grafana/latest/dashboards/build-dashboards/annotate-visualizations/
export interface AnnotationContainer {
	// List of annotations
	list?: AnnotationQuery[];
}

// TODO docs
// FROM: AnnotationQuery in grafana-data/src/types/annotations.ts
export interface AnnotationQuery {
	// Name of annotation.
	name: string;
	// Datasource where the annotations data is
	datasource: DataSourceRef;
	// When enabled the annotation query is issued with every dashboard refresh
	enable: boolean;
	// Annotation queries can be toggled on or off at the top of the dashboard.
	// When hide is true, the toggle is not shown in the dashboard.
	hide?: boolean;
	// Color to use for the annotation event markers
	iconColor: string;
	// Filters to apply when fetching annotations
	filter?: AnnotationPanelFilter;
	// TODO.. this should just be a normal query target
	target?: AnnotationTarget;
	// TODO -- this should not exist here, it is based on the --grafana-- datasource
	type?: string;
}

// A variable is a placeholder for a value. You can use variables in metric queries and in panel titles.
export interface VariableModel {
	// Unique numeric identifier for the variable.
	id: string;
	// Type of variable
	type: VariableType;
	// Name of variable
	name: string;
	// Optional display name
	label?: string;
	// Visibility configuration for the variable
	hide: VariableHide;
	// Whether the variable value should be managed by URL query params or not
	skipUrlSync: boolean;
	// Description of variable. It can be defined but `null`.
	description?: string;
	// Query used to fetch values for a variable
	query?: any;
	// Data source used to fetch values for a variable. It can be defined but `null`.
	datasource?: DataSourceRef;
	// Format to use while fetching all values from data source, eg: wildcard, glob, regex, pipe, etc.
	allFormat?: string;
	// Shows current selected variable text/value on the dashboard
	current?: VariableOption;
	// Whether multiple values can be selected or not from variable value list
	multi?: boolean;
	// Options that can be selected for a variable.
	options?: VariableOption[];
	refresh?: VariableRefresh;
}

// Option to be selected in a variable.
export interface VariableOption {
	// Whether the option is selected or not
	selected?: boolean;
	// Text to be displayed for the option
	text: StringOrArray;
	// Value of the option
	value: StringOrArray;
}

// Options to config when to refresh a variable
// `0`: Never refresh the variable
// `1`: Queries the data source every time the dashboard loads.
// `2`: Queries the data source when the dashboard time range changes.
export enum VariableRefresh {
	Never = 0,
	OnDashboardLoad = 1,
	OnTimeRangeChanged = 2,
}

// Determine if the variable shows on dashboard
// Accepted values are 0 (show label and value), 1 (show value only), 2 (show nothing).
export enum VariableHide {
	DontHide = 0,
	HideLabel = 1,
	HideVariable = 2,
}

// Sort variable options
// Accepted values are:
// `0`: No sorting
// `1`: Alphabetical ASC
// `2`: Alphabetical DESC
// `3`: Numerical ASC
// `4`: Numerical DESC
// `5`: Alphabetical Case Insensitive ASC
// `6`: Alphabetical Case Insensitive DESC
export enum VariableSort {
	Disabled = 0,
	AlphabeticalAsc = 1,
	AlphabeticalDesc = 2,
	NumericalAsc = 3,
	NumericalDesc = 4,
	AlphabeticalCaseInsensitiveAsc = 5,
	AlphabeticalCaseInsensitiveDesc = 6,
}

// Loading status
// Accepted values are `NotStarted` (the request is not started), `Loading` (waiting for response), `Streaming` (pulling continuous data), `Done` (response received successfully) or `Error` (failed request).
export enum LoadingState {
	NotStarted = "NotStarted",
	Loading = "Loading",
	Streaming = "Streaming",
	Done = "Done",
	Error = "Error",
}

// Ref to a DataSource instance
export interface DataSourceRef {
	// The plugin type-id
	type?: string;
	// Specific datasource instance
	uid?: string;
}

// Links with references to other dashboards or external resources
export interface DashboardLink {
	// Title to display with the link
	title: string;
	// Link type. Accepted values are dashboards (to refer to another dashboard) and link (to refer to an external resource)
	type: DashboardLinkType;
	// Icon name to be displayed with the link
	icon: string;
	// Tooltip to display when the user hovers their mouse over it
	tooltip: string;
	// Link URL. Only required/valid if the type is link
	url: string;
	// List of tags to limit the linked dashboards. If empty, all dashboards will be displayed. Only valid if the type is dashboards
	tags: string[];
	// If true, all dashboards links will be displayed in a dropdown. If false, all dashboards links will be displayed side by side. Only valid if the type is dashboards
	asDropdown: boolean;
	// If true, the link will be opened in a new tab
	targetBlank: boolean;
	// If true, includes current template variables values in the link as query params
	includeVars: boolean;
	// If true, includes current time range in the link as query params
	keepTime: boolean;
}

// Dashboard Link type. Accepted values are dashboards (to refer to another dashboard) and link (to refer to an external resource)
export enum DashboardLinkType {
	Link = "link",
	Dashboards = "dashboards",
}

// Dashboard variable type
// `query`: Query-generated list of values such as metric names, server names, sensor IDs, data centers, and so on.
// `adhoc`: Key/value filters that are automatically added to all metric queries for a data source (Prometheus, Loki, InfluxDB, and Elasticsearch only).
// `constant`: 	Define a hidden constant.
// `datasource`: Quickly change the data source for an entire dashboard.
// `interval`: Interval variables represent time spans.
// `textbox`: Display a free text input field with an optional default value.
// `custom`: Define the variable options manually using a comma-separated list.
// `system`: Variables defined by Grafana. See: https://grafana.com/docs/grafana/latest/dashboards/variables/add-template-variables/#global-variables
export enum VariableType {
	Query = "query",
	Adhoc = "adhoc",
	Constant = "constant",
	Datasource = "datasource",
	Interval = "interval",
	Textbox = "textbox",
	Custom = "custom",
	System = "system",
}

// Color mode for a field. You can specify a single color, or select a continuous (gradient) color schemes, based on a value.
// Continuous color interpolates a color using the percentage of a value relative to min and max.
// Accepted values are:
// `thresholds`: From thresholds. Informs Grafana to take the color from the matching threshold
// `palette-classic`: Classic palette. Grafana will assign color by looking up a color in a palette by series index. Useful for Graphs and pie charts and other categorical data visualizations
// `palette-classic-by-name`: Classic palette (by name). Grafana will assign color by looking up a color in a palette by series name. Useful for Graphs and pie charts and other categorical data visualizations
// `continuous-GrYlRd`: ontinuous Green-Yellow-Red palette mode
// `continuous-RdYlGr`: Continuous Red-Yellow-Green palette mode
// `continuous-BlYlRd`: Continuous Blue-Yellow-Red palette mode
// `continuous-YlRd`: Continuous Yellow-Red palette mode
// `continuous-BlPu`: Continuous Blue-Purple palette mode
// `continuous-YlBl`: Continuous Yellow-Blue palette mode
// `continuous-blues`: Continuous Blue palette mode
// `continuous-reds`: Continuous Red palette mode
// `continuous-greens`: Continuous Green palette mode
// `continuous-purples`: Continuous Purple palette mode
// `shades`: Shades of a single color. Specify a single color, useful in an override rule.
// `fixed`: Fixed color mode. Specify a single color, useful in an override rule.
export enum FieldColorModeId {
	Thresholds = "thresholds",
	PaletteClassic = "palette-classic",
	PaletteClassicByName = "palette-classic-by-name",
	ContinuousGrYlRd = "continuous-GrYlRd",
	ContinuousRdYlGr = "continuous-RdYlGr",
	ContinuousBlYlRd = "continuous-BlYlRd",
	ContinuousYlRd = "continuous-YlRd",
	ContinuousBlPu = "continuous-BlPu",
	ContinuousYlBl = "continuous-YlBl",
	ContinuousBlues = "continuous-blues",
	ContinuousReds = "continuous-reds",
	ContinuousGreens = "continuous-greens",
	ContinuousPurples = "continuous-purples",
	Fixed = "fixed",
	Shades = "shades",
}

// Defines how to assign a series color from "by value" color schemes. For example for an aggregated data points like a timeseries, the color can be assigned by the min, max or last value.
export enum FieldColorSeriesByMode {
	Min = "min",
	Max = "max",
	Last = "last",
}

// Map a field to a color.
export interface FieldColor {
	// The main color scheme mode.
	mode: FieldColorModeId;
	// The fixed color value for fixed or shades color modes.
	fixedColor?: string;
	// Some visualizations need to know how to assign a series color from by value color schemes.
	seriesBy?: FieldColorSeriesByMode;
}

// Position and dimensions of a panel in the grid
export interface GridPos {
	// Panel height. The height is the number of rows from the top edge of the panel.
	h: number;
	// Panel width. The width is the number of columns from the left edge of the panel.
	w: number;
	// Panel x. The x coordinate is the number of columns from the left edge of the grid
	x: number;
	// Panel y. The y coordinate is the number of rows from the top edge of the grid
	y: number;
	// Whether the panel is fixed within the grid. If true, the panel will not be affected by other panels' interactions
	static?: boolean;
}

// User-defined value for a metric that triggers visual changes in a panel when this value is met or exceeded
// They are used to conditionally style and color visualizations based on query results , and can be applied to most visualizations.
export interface Threshold {
	// Value represents a specified metric for the threshold, which triggers a visual change in the dashboard when this value is met or exceeded.
	// Nulls currently appear here when serializing -Infinity to JSON.
	value: number;
	// Color represents the color of the visual change that will occur in the dashboard when the threshold value is met or exceeded.
	color: string;
}

// Thresholds can either be `absolute` (specific number) or `percentage` (relative to min or max, it will be values between 0 and 1).
export enum ThresholdsMode {
	Absolute = "absolute",
	Percentage = "percentage",
}

// Thresholds configuration for the panel
export interface ThresholdsConfig {
	// Thresholds mode.
	mode: ThresholdsMode;
	// Must be sorted by 'value', first value is always -Infinity
	steps: Threshold[];
}

// Supported value mapping types
// `value`: Maps text values to a color or different display text and color. For example, you can configure a value mapping so that all instances of the value 10 appear as Perfection! rather than the number.
// `range`: Maps numerical ranges to a display text and color. For example, if a value is within a certain range, you can configure a range value mapping to display Low or High rather than the number.
// `regex`: Maps regular expressions to replacement text and a color. For example, if a value is www.example.com, you can configure a regex value mapping so that Grafana displays www and truncates the domain.
// `special`: Maps special values like Null, NaN (not a number), and boolean values like true and false to a display text and color. See SpecialValueMatch to see the list of special values. For example, you can configure a special value mapping so that null values appear as N/A.
export enum MappingType {
	ValueToText = "value",
	RangeToText = "range",
	RegexToText = "regex",
	SpecialValue = "special",
}

// Maps text values to a color or different display text and color.
// For example, you can configure a value mapping so that all instances of the value 10 appear as Perfection! rather than the number.
export interface ValueMap {
	type: string;
	// Map with <value_to_match>: ValueMappingResult. For example: { "10": { text: "Perfection!", color: "green" } }
	options: Record<&{string}, ValueMappingResult>;
}

// Maps numerical ranges to a display text and color.
// For example, if a value is within a certain range, you can configure a range value mapping to display Low or High rather than the number.
export interface RangeMap {
	type: string;
	// Range to match against and the result to apply when the value is within the range
	options: {
		// Min value of the range. It can be null which means -Infinity
		from: number;
		// Max value of the range. It can be null which means +Infinity
		to: number;
		// Config to apply when the value is within the range
		result: ValueMappingResult;
	};
}

// Maps regular expressions to replacement text and a color.
// For example, if a value is www.example.com, you can configure a regex value mapping so that Grafana displays www and truncates the domain.
export interface RegexMap {
	type: string;
	// Regular expression to match against and the result to apply when the value matches the regex
	options: {
		// Regular expression to match against
		pattern: string;
		// Config to apply when the value matches the regex
		result: ValueMappingResult;
	};
}

// Maps special values like Null, NaN (not a number), and boolean values like true and false to a display text and color.
// See SpecialValueMatch to see the list of special values.
// For example, you can configure a special value mapping so that null values appear as N/A.
export interface SpecialValueMap {
	type: string;
	options: {
		// Special value to match against
		match: SpecialValueMatch;
		// Config to apply when the value matches the special value
		result: ValueMappingResult;
	};
}

// Special value types supported by the `SpecialValueMap`
export enum SpecialValueMatch {
	True = "true",
	False = "false",
	Null = "null",
	NaN = "nan",
	NullAndNan = "null+nan",
	Empty = "empty",
}

// Result used as replacement with text and color when the value matches
export interface ValueMappingResult {
	// Text to display when the value matches
	text?: string;
	// Text to use when the value matches
	color?: string;
	// Icon to display when the value matches. Only specific visualizations.
	icon?: string;
	// Position in the mapping array. Only used internally.
	index?: number;
}

// Transformations allow to manipulate data returned by a query before the system applies a visualization.
// Using transformations you can: rename fields, join time series data, perform mathematical operations across queries,
// use the output of one transformation as the input to another transformation, etc.
export interface DataTransformerConfig {
	// Unique identifier of transformer
	id: string;
	// Disabled transformations are skipped
	disabled?: boolean;
	// Optional frame matcher. When missing it will be applied to all results
	filter?: MatcherConfig;
	// Options to be passed to the transformer
	// Valid options depend on the transformer id
	options: any;
}

// 0 for no shared crosshair or tooltip (default).
// 1 for shared crosshair.
// 2 for shared crosshair AND shared tooltip.
export enum DashboardCursorSync {
	Off = 0,
	Crosshair = 1,
	Tooltip = 2,
}

// Schema for panel targets is specified by datasource
// plugins. We use a placeholder definition, which the Go
// schema loader either left open/as-is with the Base
// variant of the Dashboard and Panel families, or filled
// with types derived from plugins in the Instance variant.
// When working directly from CUE, importers can extend this
// type directly to achieve the same effect.
export interface Target {

}

// Dashboard panels are the basic visualization building blocks.
export interface Panel {
	// The panel plugin type id. This is used to find the plugin to display the panel.
	type: string;
	// Unique identifier of the panel. Generated by Grafana when creating a new panel. It must be unique within a dashboard, but not globally.
	id?: number;
	// The version of the plugin that is used for this panel. This is used to find the plugin to display the panel and to migrate old panel configs.
	pluginVersion?: string;
	// Tags for the panel.
	tags?: string[];
	// Depends on the panel plugin. See the plugin documentation for details.
	targets?: Target[];
	// Panel title.
	title?: string;
	// Panel description.
	description?: string;
	// Whether to display the panel without a background.
	transparent: boolean;
	// The datasource used in all targets.
	datasource?: DataSourceRef;
	// Grid position.
	gridPos?: GridPos;
	// Panel links.
	links?: DashboardLink[];
	// Name of template variable to repeat for.
	repeat?: string;
	// Direction to repeat in if 'repeat' is set.
	// `h` for horizontal, `v` for vertical.
	repeatDirection?: PanelRepeatDirection;
	// Id of the repeating panel.
	repeatPanelId?: number;
	// The maximum number of data points that the panel queries are retrieving.
	maxDataPoints?: number;
	// List of transformations that are applied to the panel data before rendering.
	// When there are multiple transformations, Grafana applies them in the order they are listed.
	// Each transformation creates a result set that then passes on to the next transformation in the processing pipeline.
	transformations: DataTransformerConfig[];
	// The min time interval setting defines a lower limit for the $__interval and $__interval_ms variables.
	// This value must be formatted as a number followed by a valid time
	// identifier like: "40s", "3d", etc.
	// See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
	interval?: string;
	// Overrides the relative time range for individual panels,
	// which causes them to be different than what is selected in
	// the dashboard time picker in the top-right corner of the dashboard. You can use this to show metrics from different
	// time periods or days on the same dashboard.
	// The value is formatted as time operation like: `now-5m` (Last 5 minutes), `now/d` (the day so far),
	// `now-5d/d`(Last 5 days), `now/w` (This week so far), `now-2y/y` (Last 2 years).
	// Note: Panel time overrides have no effect when the dashboard’s time range is absolute.
	// See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
	timeFrom?: string;
	// Overrides the time range for individual panels by shifting its start and end relative to the time picker.
	// For example, you can shift the time range for the panel to be two hours earlier than the dashboard time picker setting `2h`.
	// Note: Panel time overrides have no effect when the dashboard’s time range is absolute.
	// See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
	timeShift?: string;
	// Dynamically load the panel
	libraryPanel?: LibraryPanelRef;
	// It depends on the panel plugin. They are specified by the Options field in panel plugin schemas.
	options: any;
	// Field options allow you to change how the data is displayed in your visualizations.
	fieldConfig: FieldConfigSource;
}

export enum PanelRepeatDirection {
	Horizontal = "h",
	Vertical = "v",
}

// The data model used in Grafana, namely the data frame, is a columnar-oriented table structure that unifies both time series and table query results.
// Each column within this structure is called a field. A field can represent a single time series or table column.
// Field options allow you to change how the data is displayed in your visualizations.
export interface FieldConfigSource {
	// Defaults are the options applied to all fields.
	defaults: FieldConfig;
	// Overrides are the options applied to specific fields overriding the defaults.
	overrides: FieldConfigSourceOverride[];
}

export interface FieldConfigSourceOverride {
	matcher: MatcherConfig;
	properties: DynamicConfigValue[];
}

// A library panel is a reusable panel that you can use in any dashboard.
// When you make a change to a library panel, that change propagates to all instances of where the panel is used.
// Library panels streamline reuse of panels across multiple dashboards.
export interface LibraryPanelRef {
	// Library panel name
	name: string;
	// Library panel uid
	uid: string;
}

// Matcher is a predicate configuration. Based on the config a set of field(s) or values is filtered in order to apply override / transformation.
// It comes with in id ( to resolve implementation from registry) and a configuration that’s specific to a particular matcher type.
export interface MatcherConfig {
	// The matcher id. This is used to find the matcher implementation from registry.
	id: string;
	// The matcher options. This is specific to the matcher implementation.
	options?: any;
}

export interface DynamicConfigValue {
	id: string;
	value?: any;
}

// The data model used in Grafana, namely the data frame, is a columnar-oriented table structure that unifies both time series and table query results.
// Each column within this structure is called a field. A field can represent a single time series or table column.
// Field options allow you to change how the data is displayed in your visualizations.
export interface FieldConfig {
	// The display value for this field.  This supports template variables blank is auto
	displayName?: string;
	// This can be used by data sources that return and explicit naming structure for values and labels
	// When this property is configured, this value is used rather than the default naming strategy.
	displayNameFromDS?: string;
	// Human readable field metadata
	description?: string;
	// An explicit path to the field in the datasource.  When the frame meta includes a path,
	// This will default to `${frame.meta.path}/${field.name}
	// 
	// When defined, this value can be used as an identifier within the datasource scope, and
	// may be used to update the results
	path?: string;
	// True if data source can write a value to the path. Auth/authz are supported separately
	writeable?: boolean;
	// True if data source field supports ad-hoc filters
	filterable?: boolean;
	// Unit a field should use. The unit you select is applied to all fields except time.
	// You can use the units ID availables in Grafana or a custom unit.
	// Available units in Grafana: https://github.com/grafana/grafana/blob/main/packages/grafana-data/src/valueFormats/categories.ts
	// As custom unit, you can use the following formats:
	// `suffix:<suffix>` for custom unit that should go after value.
	// `prefix:<prefix>` for custom unit that should go before value.
	// `time:<format>` For custom date time formats type for example `time:YYYY-MM-DD`.
	// `si:<base scale><unit characters>` for custom SI units. For example: `si: mF`. This one is a bit more advanced as you can specify both a unit and the source data scale. So if your source data is represented as milli (thousands of) something prefix the unit with that SI scale character.
	// `count:<unit>` for a custom count unit.
	// `currency:<unit>` for custom a currency unit.
	unit?: string;
	// Specify the number of decimals Grafana includes in the rendered value.
	// If you leave this field blank, Grafana automatically truncates the number of decimals based on the value.
	// For example 1.1234 will display as 1.12 and 100.456 will display as 100.
	// To display all decimals, set the unit to `String`.
	decimals?: number;
	// The minimum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.
	min?: number;
	// The maximum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.
	max?: number;
	// Convert input values into a display string
	mappings?: ValueMapOrRangeMapOrRegexMapOrSpecialValueMap[];
	// Map numeric values to states
	thresholds?: ThresholdsConfig;
	// Panel color configuration
	color?: FieldColor;
	// The behavior when clicking on a result
	links?: any[];
	// Alternative to empty string
	noValue?: string;
	// custom is specified by the FieldConfig field
	// in panel plugin schemas.
	custom?: any;
}

// Row panel
export interface RowPanel {
	// The panel type
	type: string;
	// Whether this row should be collapsed or not.
	collapsed: boolean;
	// Row title
	title?: string;
	// Name of default datasource for the row
	datasource?: DataSourceRef;
	// Row grid position
	gridPos?: GridPos;
	// Unique identifier of the panel. Generated by Grafana when creating a new panel. It must be unique within a dashboard, but not globally.
	id: number;
	// List of panels in the row
	panels: Panel[];
	// Name of template variable to repeat for.
	repeat?: string;
}

