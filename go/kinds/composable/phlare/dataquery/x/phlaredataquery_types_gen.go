// Code generated - EDITING IS FUTILE. DO NOT EDIT.
//
// Generated by pipeline:
//     go
// Using jennies:
//     GoTypesJenny
//     ComposableLatestMajorsOrXJenny
//
// Run 'go generate ./' from repository root to regenerate.

package phlaredataquery

// Defines values for PhlareQueryType.
const (
	PhlareQueryTypeBoth PhlareQueryType = "both"

	PhlareQueryTypeMetrics PhlareQueryType = "metrics"

	PhlareQueryTypeProfile PhlareQueryType = "profile"
)

// These are the common properties available to all queries in all datasources.
// Specific implementations will *extend* this interface, adding the required
// properties for the given context.
type DataQuery struct {
	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	Datasource *interface{} `json:"datasource,omitempty"`

	// true if query is disabled (ie should not be returned to the dashboard)
	Hide *bool `json:"hide,omitempty"`

	// Unique, guid like, string used in explore mode
	Key *string `json:"key,omitempty"`

	// Specify the query flavor
	// TODO make this required and give it a default
	QueryType *string `json:"queryType,omitempty"`

	// A - Z
	RefId string `json:"refId"`
}

// PhlareDataQuery defines model for PhlareDataQuery.
type PhlareDataQuery struct {
	// Embedded struct due to allOf(#/components/schemas/DataQuery)
	DataQuery `yaml:",inline"`
	// Embedded fields due to inline allOf schema
}

// PhlareQueryType defines model for PhlareQueryType.
type PhlareQueryType string