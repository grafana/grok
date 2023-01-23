package dashboardlistpanelcfg

// Defines values for PanelLayout.
const (
	PanelLayoutList PanelLayout = "list"

	PanelLayoutPreviews PanelLayout = "previews"
)

// PanelLayout defines model for PanelLayout.
type PanelLayout string

// PanelOptions defines model for PanelOptions.
type PanelOptions struct {
	FolderId           *int         `json:"folderId,omitempty"`
	Layout             *PanelLayout `json:"layout,omitempty"`
	MaxItems           int          `json:"maxItems"`
	Query              string       `json:"query"`
	ShowHeadings       bool         `json:"showHeadings"`
	ShowRecentlyViewed bool         `json:"showRecentlyViewed"`
	ShowSearch         bool         `json:"showSearch"`
	ShowStarred        bool         `json:"showStarred"`
	Tags               []string     `json:"tags"`
}
