// Code generated - EDITING IS FUTILE. DO NOT EDIT.
//
// Generated by pipeline:
//     go
// Using jennies:
//     GoTypesJenny
//     ComposableLatestMajorsOrXJenny
//
// Run 'go generate ./' from repository root to regenerate.

package tablepanelcfg

// Defines values for TableCellHeight.
const (
	TableCellHeightLg TableCellHeight = "lg"
	TableCellHeightMd TableCellHeight = "md"
	TableCellHeightSm TableCellHeight = "sm"
)

// PanelOptions defines model for PanelOptions.
type PanelOptions struct {
	// Height of a table cell
	CellHeight *TableCellHeight `json:"cellHeight,omitempty"`

	// Controls footer options
	Footer *interface{} `json:"footer,omitempty"`

	// Represents the index of the selected frame
	FrameIndex float32 `json:"frameIndex"`

	// Controls whether the panel should show the header
	ShowHeader bool `json:"showHeader"`

	// Controls whether the columns should be numbered
	ShowRowNums *bool `json:"showRowNums,omitempty"`

	// Controls whether the header should show icons for the column types
	ShowTypeIcons *bool `json:"showTypeIcons,omitempty"`

	// Used to control row sorting
	SortBy []TableSortByFieldState `json:"sortBy,omitempty"`
}

// Height of a table cell
type TableCellHeight string

// Footer options
type TableFooterOptions struct {
	CountRows        *bool    `json:"countRows,omitempty"`
	EnablePagination *bool    `json:"enablePagination,omitempty"`
	Fields           []string `json:"fields,omitempty"`
	Reducer          []string `json:"reducer"`
	Show             bool     `json:"show"`
}

// Sort by field state
type TableSortByFieldState struct {
	// Flag used to indicate descending sort order
	Desc *bool `json:"desc,omitempty"`

	// Sets the display name of the field to sort by
	DisplayName string `json:"displayName"`
}
