package annolist

// PanelOptions defines model for PanelOptions.
type PanelOptions struct {
	Limit                 int      `json:"limit"`
	NavigateAfter         string   `json:"navigateAfter"`
	NavigateBefore        string   `json:"navigateBefore"`
	NavigateToPanel       bool     `json:"navigateToPanel"`
	OnlyFromThisDashboard bool     `json:"onlyFromThisDashboard"`
	OnlyInTimeRange       bool     `json:"onlyInTimeRange"`
	ShowTags              bool     `json:"showTags"`
	ShowTime              bool     `json:"showTime"`
	ShowUser              bool     `json:"showUser"`
	Tags                  []string `json:"tags"`
}
