package annotationpanelfilter

import "github.com/grafana/grok/newgen/dashboard/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.AnnotationPanelFilter
}
func New(options ...Option) (Builder, error) {
	resource := &types.AnnotationPanelFilter{}
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

func (builder *Builder) Internal() *types.AnnotationPanelFilter {
	return builder.internal
}
// Should the specified panels be included or excluded
func Exclude(exclude bool) Option {
	return func(builder *Builder) error {
		
		builder.internal.Exclude = exclude

		return nil
	}
}
// Panel IDs that should be included or excluded
func Ids(ids []uint8) Option {
	return func(builder *Builder) error {
		
		builder.internal.Ids = ids

		return nil
	}
}

func defaults() []Option {
return []Option{
}
}
