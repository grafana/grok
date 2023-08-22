package dashboard

import "github.com/grafana/grok/newgen/dashboard/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.RowPanel
}

func Type(typeArg string) Option {
	return func(builder *Builder) error {

		builder.internal.Type = typeArg

		return nil
	}
}

func Collapsed(collapsed bool) Option {
	return func(builder *Builder) error {

		builder.internal.Collapsed = collapsed

		return nil
	}
}

func Title(title string) Option {
	return func(builder *Builder) error {

		builder.internal.Title = &title

		return nil
	}
}

func Datasource(datasource types.DataSourceRef) Option {
	return func(builder *Builder) error {

		builder.internal.Datasource = &datasource

		return nil
	}
}

func GridPos(gridPos types.GridPos) Option {
	return func(builder *Builder) error {

		builder.internal.GridPos = &gridPos

		return nil
	}
}

func Id(id uint32) Option {
	return func(builder *Builder) error {

		builder.internal.Id = id

		return nil
	}
}

func Panels(panels []types.Panel) Option {
	return func(builder *Builder) error {

		builder.internal.Panels = panels

		return nil
	}
}

func Repeat(repeat string) Option {
	return func(builder *Builder) error {

		builder.internal.Repeat = &repeat

		return nil
	}
}

func defaults() []Option {
	return []Option{
		Collapsed(false),
	}
}
