// Code generated - EDITING IS FUTILE. DO NOT EDIT.
//
// Generated by pipeline:
//     go
// Using jennies:
//     GoTypesJenny
//     ComposableLatestMajorsOrXJenny
//
// Run 'go generate ./' from repository root to regenerate.

package newspanelcfg

// PanelOptions defines model for PanelOptions.
type PanelOptions struct {
	// FeedUrl empty/missing will default to grafana blog
	FeedUrl   *string `json:"feedUrl,omitempty"`
	ShowImage *bool   `json:"showImage,omitempty"`
}
