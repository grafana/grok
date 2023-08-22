package dashboard

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

func Title(title string) Option {
	return func(builder *Builder) error {

		builder.internal.Title = title

		return nil
	}
}

func Type(typeArg types.DashboardLinkType) Option {
	return func(builder *Builder) error {

		builder.internal.Type = typeArg

		return nil
	}
}

func Icon(icon string) Option {
	return func(builder *Builder) error {

		builder.internal.Icon = icon

		return nil
	}
}

func Tooltip(tooltip string) Option {
	return func(builder *Builder) error {

		builder.internal.Tooltip = tooltip

		return nil
	}
}

func Url(url string) Option {
	return func(builder *Builder) error {

		builder.internal.Url = url

		return nil
	}
}

func Tags(tags []string) Option {
	return func(builder *Builder) error {

		builder.internal.Tags = tags

		return nil
	}
}

func AsDropdown(asDropdown bool) Option {
	return func(builder *Builder) error {

		builder.internal.AsDropdown = asDropdown

		return nil
	}
}

func TargetBlank(targetBlank bool) Option {
	return func(builder *Builder) error {

		builder.internal.TargetBlank = targetBlank

		return nil
	}
}

func IncludeVars(includeVars bool) Option {
	return func(builder *Builder) error {

		builder.internal.IncludeVars = includeVars

		return nil
	}
}

func KeepTime(keepTime bool) Option {
	return func(builder *Builder) error {

		builder.internal.KeepTime = keepTime

		return nil
	}
}

func defaults() []Option {
	return []Option{
		AsDropdown(false),
		TargetBlank(false),
		IncludeVars(false),
		KeepTime(false),
	}
}
