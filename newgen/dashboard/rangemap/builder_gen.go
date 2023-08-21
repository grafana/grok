package dashboard

import "github.com/grafana/grok/newgen/dashboard/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.RangeMap
}

func Type(type string) Option {
	return func(builder *Builder) error {
		if !(type == range) {
return errors.New("type must be == range")
}

		builder.internal.Type = type

		return nil
	}
}

func Options(options struct {
	// Min value of the range. It can be null which means -Infinity
From *float64 `json:"from"`
	// Max value of the range. It can be null which means +Infinity
To *float64 `json:"to"`
	// Config to apply when the value is within the range
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
