// Code generated - EDITING IS FUTILE. DO NOT EDIT.
//
// Generated by pipeline:
//     go
// Using jennies:
//     GoTypesJenny
//     ComposableLatestMajorsOrXJenny
//
// Run 'go generate ./' from repository root to regenerate.

package nodegraphpanelcfg

// ArcOption defines model for ArcOption.
type ArcOption struct {
	// The color of the arc.
	Color *string `json:"color,omitempty"`

	// Field from which to get the value. Values should be less than 1, representing fraction of a circle.
	Field *string `json:"field,omitempty"`
}

// EdgeOptions defines model for EdgeOptions.
type EdgeOptions struct {
	// Unit for the main stat to override what ever is set in the data frame.
	MainStatUnit *string `json:"mainStatUnit,omitempty"`

	// Unit for the secondary stat to override what ever is set in the data frame.
	SecondaryStatUnit *string `json:"secondaryStatUnit,omitempty"`
}

// NodeOptions defines model for NodeOptions.
type NodeOptions struct {
	// Define which fields are shown as part of the node arc (colored circle around the node).
	Arcs *[]ArcOption `json:"arcs,omitempty"`

	// Unit for the main stat to override what ever is set in the data frame.
	MainStatUnit *string `json:"mainStatUnit,omitempty"`

	// Unit for the secondary stat to override what ever is set in the data frame.
	SecondaryStatUnit *string `json:"secondaryStatUnit,omitempty"`
}

// PanelOptions defines model for PanelOptions.
type PanelOptions struct {
	Edges *EdgeOptions `json:"edges,omitempty"`
	Nodes *NodeOptions `json:"nodes,omitempty"`
}