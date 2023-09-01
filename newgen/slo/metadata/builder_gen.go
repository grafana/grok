package metadata

import "github.com/grafana/grok/newgen/slo/types"

type Option func(builder *Builder) error

type Builder struct {
	internal *types.Metadata
}
func New(options ...Option) (Builder, error) {
	resource := &types.Metadata{}
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

func (builder *Builder) Internal() *types.Metadata {
	return builder.internal
}
func UpdateTimestamp(updateTimestamp string) Option {
	return func(builder *Builder) error {
		
		builder.internal.UpdateTimestamp = updateTimestamp

		return nil
	}
}
func CreatedBy(createdBy string) Option {
	return func(builder *Builder) error {
		
		builder.internal.CreatedBy = createdBy

		return nil
	}
}
func UpdatedBy(updatedBy string) Option {
	return func(builder *Builder) error {
		
		builder.internal.UpdatedBy = updatedBy

		return nil
	}
}
func Uid(uid string) Option {
	return func(builder *Builder) error {
		
		builder.internal.Uid = uid

		return nil
	}
}
func CreationTimestamp(creationTimestamp string) Option {
	return func(builder *Builder) error {
		
		builder.internal.CreationTimestamp = creationTimestamp

		return nil
	}
}
func DeletionTimestamp(deletionTimestamp string) Option {
	return func(builder *Builder) error {
		
		builder.internal.DeletionTimestamp = &deletionTimestamp

		return nil
	}
}
func Finalizers(finalizers []string) Option {
	return func(builder *Builder) error {
		
		builder.internal.Finalizers = finalizers

		return nil
	}
}
func ResourceVersion(resourceVersion string) Option {
	return func(builder *Builder) error {
		
		builder.internal.ResourceVersion = resourceVersion

		return nil
	}
}
// extraFields is reserved for any fields that are pulled from the API server metadata but do not have concrete fields in the CUE metadata
func ExtraFields(extraFields any) Option {
	return func(builder *Builder) error {
		
		builder.internal.ExtraFields = extraFields

		return nil
	}
}
func Labels(labels map[string]string) Option {
	return func(builder *Builder) error {
		
		builder.internal.Labels = labels

		return nil
	}
}

func defaults() []Option {
return []Option{
}
}
