// Code generated - EDITING IS FUTILE. DO NOT EDIT.
//
// Generated by pipeline:
//     go
// Using jennies:
//     GoTypesJenny
//     ComposableLatestMajorsOrXJenny
//
// Run 'go generate ./' from repository root to regenerate.

package debugpanelcfg

// Defines values for DebugMode.
const (
	DebugModeCursor     DebugMode = "cursor"
	DebugModeEvents     DebugMode = "events"
	DebugModeRender     DebugMode = "render"
	DebugModeState      DebugMode = "State"
	DebugModeThrowError DebugMode = "ThrowError"
)

// DebugMode defines model for DebugMode.
type DebugMode string

// PanelOptions defines model for PanelOptions.
type PanelOptions struct {
	Counters *UpdateConfig `json:"counters,omitempty"`
	Mode     DebugMode     `json:"mode"`
}

// UpdateConfig defines model for UpdateConfig.
type UpdateConfig struct {
	DataChanged   bool `json:"dataChanged"`
	Render        bool `json:"render"`
	SchemaChanged bool `json:"schemaChanged"`
}