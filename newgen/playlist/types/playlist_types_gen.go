package types

type PlaylistItem struct {
	Type  PlaylistItemType `json:"type"`
	Value string           `json:"value"`
}

type PlaylistItemType string

const (
	DashboardByTag PlaylistItemType = "dashboard_by_tag"
	DashboardByUid PlaylistItemType = "dashboard_by_uid"
)

type Playlist struct {
	Interval string         `json:"interval"`
	Items    []PlaylistItem `json:"items"`
	Name     string         `json:"name"`
	Xxx      string         `json:"xxx"`
}
