package cuetf

import (
	"bytes"
	"fmt"

	"cuelang.org/go/cue"
	"github.com/grafana/thema"
)

// GenerateDataSource takes a cue.Value and generates the corresponding Terraform data source
func GenerateDataSource(schema thema.Schema) (b []byte, err error) {
	schemaAttributes, err := GenerateSchemaAttributes(schema.Underlying())

	vars := TVarsDataSource{
		Name:             schema.Lineage().Name(),
		Description:      "TODO description",
		ModelFields:      "TODO modelFields",
		SchemaAttributes: string(schemaAttributes),
	}

	buf := new(bytes.Buffer)
	if err := tmpls.Lookup("datasource.tmpl").Execute(buf, vars); err != nil {
		return nil, fmt.Errorf("failed executing datasource template: %w", err)
	}

	return buf.Bytes(), nil
}

func GenerateSchemaAttributes(val cue.Value) ([]byte, error) {
	if err := val.Validate(); err != nil {
		return nil, err
	}

	iter, err := val.Fields(
		cue.Definitions(true),
		cue.Optional(true),
	)
	if err != nil {
		return nil, err
	}

	attributes := make([]byte, 0)
	for iter.Next() {
		attr, err := genSingleSchemaAttribute(iter.Selector().String(), iter.Value())
		if err != nil {
			return nil, err
		}

		attributes = append(attributes, attr...)
		attributes = append(attributes, []byte("\n")...)
	}

	return attributes, nil
}

func genSingleSchemaAttribute(name string, value cue.Value) ([]byte, error) {
	vars := TVarsSchemaAttribute{
		Name:     name,
		Computed: true,
	}

	for _, comment := range value.Doc() {
		vars.Description += comment.Text()
	}

	// ListAttribute
	// MapAttribute
	// ObjectAttribute
	// SetAttribute

	switch value.Kind() {
	case cue.BoolKind:
		vars.AttributeType = "Bool"
	case cue.IntKind:
		vars.AttributeType = "Int64"
	case cue.FloatKind:
		vars.AttributeType = "Float64"
	case cue.NumberKind:
		vars.AttributeType = "Number"
	case cue.StringKind, cue.BytesKind:
		vars.AttributeType = "String"
	case cue.StructKind:
		// TODO
	case cue.ListKind:
		// TODO
	case cue.BottomKind, cue.NullKind, cue.TopKind:
		// TODO
	}

	buf := new(bytes.Buffer)
	if err := tmpls.Lookup("datasource.tmpl").Execute(buf, vars); err != nil {
		return nil, fmt.Errorf("failed executing datasource template: %w", err)
	}

	return buf.Bytes(), nil
}
