package dashboard

import "github.com/grafana/grok/newgen/dashboard/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.RegexMap
}

func Type(type string) Option {
	return func(builder *Builder) error {
		if !(type == regex) {
return errors.New("type must be == regex")
}

		builder.internal.Type = type

		return nil
	}
}

func Options(options struct {
	// Regular expression to match against
Pattern string `json:"pattern"`
	// Config to apply when the value matches the regex
Result types.ValueMappingResult `json:"result"`
}) Option {
	return func(builder *Builder) error {
		
		builder.internal.Options = options

		return nil
	}
}

func defaults() []Option {
return []Option{
}
}
