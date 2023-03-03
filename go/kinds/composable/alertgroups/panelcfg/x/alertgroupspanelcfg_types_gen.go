// Code generated - EDITING IS FUTILE. DO NOT EDIT.
//
// Generated by pipeline:
//     go
// Using jennies:
//     GoTypesJenny
//     ComposableLatestMajorsOrXJenny
//
// Run 'go generate ./' from repository root to regenerate.

package alertgroupspanelcfg

// PanelOptions defines model for PanelOptions.
type PanelOptions struct {
	// Name of the alertmanager used as a source for alerts
	Alertmanager string `json:"alertmanager"`

	// Expand all alert groups by default
	ExpandAll bool `json:"expandAll"`

	// Comma-separated list of values used to filter alert results
	Labels string `json:"labels"`
}
