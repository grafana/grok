package newspanelcfg

// PanelOptions defines model for PanelOptions.
type PanelOptions struct {
	// empty/missing will default to grafana blog
	FeedUrl   *string `json:"feedUrl,omitempty"`
	ShowImage *bool   `json:"showImage,omitempty"`
}
