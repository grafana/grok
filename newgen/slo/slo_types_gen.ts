export interface Objective {
	// is a value between 0 and 1 if the value of the query's output
	// is above the objective, the SLO is met.
	value: number;
	// is a Prometheus-parsable time duration string like 24h, 60m. This is the time
	// window the objective is measured over.
	window: string;
}

export interface Query {

}

export interface ThresholdQuery {
	groupByLabels?: string[];
	thresholdMetric: MetricDef;
	threshold: Threshold;
}

export interface RatioQuery {
	groupByLabels?: string[];
	successMetric: MetricDef;
	totalMetric: MetricDef;
}

export interface HistogramQuery {
	groupByLabels?: string[];
	histogramMetric: MetricDef;
	percentile: number;
	threshold: Threshold;
}

export interface FreeformQuery {
	freeformQuery: string;
}

export interface Threshold {
	value: number;
	operator: OperatorEnum;
}

export interface MetricDef {
	prometheusMetric: string;
	type?: string;
}

export interface GrafanaMetadata {
	organizationId: string;
	userEmail: string;
	userName: string;
	provenance: string;
}

export interface Label {
	key: string;
	value: string;
}

export interface AlertingMetadata {
	labels?: Label[];
	annotations?: Label[];
}

export interface Alerting {
	// will be attached to all alerts generated by any of these rules.
	labels?: Label[];
	// will be attached to all alerts generated by any of these rules.
	annotations?: Label[];
	// Metadata to attach only to fastBurn alerts.
	fastBurn?: AlertingMetadata;
	// Metadata to attach only to slowBurn alerts.
	slowBurn?: AlertingMetadata;
}

// metadata contains embedded CommonMetadata and can be extended with custom string fields
// TODO: use CommonMetadata instead of redefining here; currently needs to be defined here
// without external reference as using the CommonMetadata reference breaks thema codegen.
export interface Metadata {
	updateTimestamp: string;
	createdBy: string;
	updatedBy: string;
	uid: string;
	creationTimestamp: string;
	deletionTimestamp?: string;
	finalizers: string[];
	resourceVersion: string;
	// extraFields is reserved for any fields that are pulled from the API server metadata but do not have concrete fields in the CUE metadata
	extraFields: any;
	labels: Record<&{string}, string>;
}

export interface Spec {
	// This is used internally by the plugin for permission management and
	// similar functions.
	grafanaMetadata?: GrafanaMetadata;
	// A unique, random identifier. This value will also be the name of the
	// resource stored in the API server. Must be set for a PUT.
	uuid: string;
	// should be a short description of your indicator. Consider names like
	// "API Availability"
	name: string;
	// is a free-text field that can provide more context to an
	// SLO. It is shown on SLO drill-down dashboards and in hover text on
	// the SLO summary dashboard.
	description: string;
	// describes the indicator that will be measured against the
	// objective. Four query types are supported:
	// 1. Ratio Queries provide a successMetric and totalMetric whose ratio is the SLI.
	// 2. Threshold Queries provide a thresholdMetric and a threshold. The
	//    SLI is the boolean result of evaluating the threshould.
	// 3. Histogram Queries are similar to threshold queries, but the use a
	//    Prometheus histogram metric, percentile value, and a threshold to
	//    generate the boolean output.
	// 4. Freeform Queries supply a single freeFormQuery string that is
	//    evaluated to produce the SLI output. The value should range beween 0
	//    and 1.0. Freeform queries should include a time variable named
	//    either `$__rate_interval`,`$__interval` or `$__range`. This will be used by the
	//    tool to evaluate the burn rate of an SLO over various time
	//    windows. Queries that don't include this interval will have
	//    sensitive and imprecise alerting.
	// Additionally, "groupByLabels" are used in the first three query types
	// to define how to group series for evaluation. They are discarded for
	// freeform queries.
	query: Query;
	// You can have multiple time windows and objectives associated with an
	// SLO. Over each rolling time window, the remaining error budget will
	// be calculated, and separate alerts can be generated for each time
	// window based on the SLO burn rate or remaining error budget.
	objectives: Objective[];
	// Any additional labels that will be attached to all metrics generated
	// from the query. These labels are useful for grouping SLOs in
	// dashboard views that you create by hand.
	// The key must match the prometheus label requirements regex:
	// "^[a-zA-Z_][a-zA-Z0-9_]*$"
	labels?: Label[];
	// Configures the alerting rules that will be generated for each
	// time window associated with the SLO. Grafana SLOs can generate
	// alerts when the short-term error budget burn is very high, the
	// long-term error budget burn rate is high, or when the remaining
	// error budget is below a certain threshold.
	alerting?: Alerting;
}

export interface OperatorState {
	// lastEvaluation is the ResourceVersion last evaluated
	lastEvaluation: string;
	// state describes the state of the lastEvaluation.
	// It is limited to three possible states for machine evaluation.
	state: StateEnum;
	// descriptiveState is an optional more descriptive state field which has no requirements on format
	descriptiveState?: string;
	// details contains any extra information that is operator-specific
	details?: any;
}

// Status is a common kubernetes subresource that is used to provide
// information about the current state, that isn't a direct part of the
// resource. Here we use it to provide a pointer to the generated
// dashboard.
export interface Status {
	drillDownDashboard: {
		uid: string;
		// The generation of the SLO when this dashboard was last updated.
		reconciledForGeneration: string;
		lastError: string;
	};
	// operatorStates is a map of operator ID to operator state evaluations.
	// Any operator which consumes this kind SHOULD add its state evaluation information to this field.
	operatorStates?: any;
	prometheusRules: {
		// The generation of the SLO when these rules were last updated.
		reconciledForGeneration: string;
		lastError: string;
	};
	// additionalFields is reserved for future use
	additionalFields?: any;
}

