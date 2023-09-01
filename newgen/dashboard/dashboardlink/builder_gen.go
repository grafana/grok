package dashboardlink

import "github.com/grafana/grok/newgen/dashboard/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.DashboardLink
}
func New(options ...Option) (Builder, error) {
	resource := &types.DashboardLink{}
	builder := &Builder{internal: resource}

	for _, opt := range append(defaults(), options...) {
		if err := opt(builder); err != nil {
			return *builder, err
		}
	}

	return *builder, nil
}
// MarshalJSON implements the encoding/json.Marshaler interface.
//
// This method can be used to render the resource as JSON
// which your configuration management tool of choice can then feed into
// Grafana.
func (builder *Builder) MarshalJSON() ([]byte, error) {
	return json.Marshal(builder.internal)
}

// MarshalIndentJSON renders the resource as indented JSON
// which your configuration management tool of choice can then feed into
// Grafana.
func (builder *Builder) MarshalIndentJSON() ([]byte, error) {
	return json.MarshalIndent(builder.internal, "", "  ")
}

func (builder *Builder) Internal() *types.DashboardLink {
	return builder.internal
}
// Title to display with the link
func Title(title string) Option {
	return func(builder *Builder) error {
		
		builder.internal.Title = title

		return nil
	}
}
// Link type. Accepted values are dashboards (to refer to another dashboard) and link (to refer to an external resource)
func Type(typeArg types.DashboardLinkType) Option {
	return func(builder *Builder) error {
		
		builder.internal.Type = typeArg

		return nil
	}
}
// Icon name to be displayed with the link
func Icon(icon string) Option {
	return func(builder *Builder) error {
		
		builder.internal.Icon = icon

		return nil
	}
}
// Tooltip to display when the user hovers their mouse over it
func Tooltip(tooltip string) Option {
	return func(builder *Builder) error {
		
		builder.internal.Tooltip = tooltip

		return nil
	}
}
// Link URL. Only required/valid if the type is link
func Url(url string) Option {
	return func(builder *Builder) error {
		
		builder.internal.Url = url

		return nil
	}
}
// List of tags to limit the linked dashboards. If empty, all dashboards will be displayed. Only valid if the type is dashboards
func Tags(tags []string) Option {
	return func(builder *Builder) error {
		
		builder.internal.Tags = tags

		return nil
	}
}
// If true, all dashboards links will be displayed in a dropdown. If false, all dashboards links will be displayed side by side. Only valid if the type is dashboards
func AsDropdown(asDropdown bool) Option {
	return func(builder *Builder) error {
		
		builder.internal.AsDropdown = asDropdown

		return nil
	}
}
// If true, the link will be opened in a new tab
func TargetBlank(targetBlank bool) Option {
	return func(builder *Builder) error {
		
		builder.internal.TargetBlank = targetBlank

		return nil
	}
}
// If true, includes current template variables values in the link as query params
func IncludeVars(includeVars bool) Option {
	return func(builder *Builder) error {
		
		builder.internal.IncludeVars = includeVars

		return nil
	}
}
// If true, includes current time range in the link as query params
func KeepTime(keepTime bool) Option {
	return func(builder *Builder) error {
		
		builder.internal.KeepTime = keepTime

		return nil
	}
}

func defaults() []Option {
return []Option{
}
}
