package cuetf

import (
	"bytes"
	"fmt"
	"go/format"
	"regexp"
	"strings"

	"cuelang.org/go/cue"
	"github.com/grafana/thema"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

// GenerateDataSource takes a cue.Value and generates the corresponding Terraform data source
func GenerateDataSource(schema thema.Schema) (b []byte, err error) {
	schemaAttributes, err := GenerateSchemaAttributes(schema.Underlying())
	if err != nil {
		return nil, err
	}

	modelFields, err := GenerateModelFields(schema.Underlying())
	if err != nil {
		return nil, err
	}

	vars := TVarsDataSource{
		Name:             schema.Lineage().Name(),
		Description:      "TODO description",
		ModelFields:      modelFields,
		SchemaAttributes: string(schemaAttributes),
	}

	buf := new(bytes.Buffer)
	if err := tmpls.Lookup("datasource.tmpl").Execute(buf, vars); err != nil {
		return nil, fmt.Errorf("failed executing datasource template: %w", err)
	}

	return format.Source(buf.Bytes())
}

func GenerateSchemaAttributes(val cue.Value) (string, error) {
	if err := val.Validate(); err != nil {
		return "", err
	}

	iter, err := val.Fields(
		cue.Definitions(true),
		cue.Optional(true),
	)
	if err != nil {
		return "", err
	}

	attributes := make([]string, 0)
	for iter.Next() {
		if iter.IsDefinition() {
			continue
		}

		attr, err := genSingleSchemaAttribute(iter.Selector().String(), iter.Value(), iter.IsOptional())
		if err != nil {
			return "", err
		}

		if attr == "" {
			continue
		}

		attributes = append(attributes, attr)
	}

	return strings.Join(attributes, "\n"), nil
}

func genSingleSchemaAttribute(name string, value cue.Value, isOptional bool) (string, error) {
	vars := TVarsSchemaAttribute{
		Name:     ToSnakeCase(name),
		Computed: false,
		Optional: isOptional,
		Required: !isOptional,
	}

	for _, comment := range value.Doc() {
		vars.Description += comment.Text()
	}
	vars.Description = strings.Trim(vars.Description, "\n ")

	// TODO: handle special cases (struct, list, bottom, null, top)
	vars.AttributeType = TypeMappings[value.IncompleteKind()]
	// TODO: jduchesne, empty attribute type fails
	if vars.AttributeType == "" {
		return "", nil
	}

	buf := new(bytes.Buffer)
	if err := tmpls.Lookup("schema_attribute.tmpl").Execute(buf, vars); err != nil {
		return "", fmt.Errorf("failed executing datasource template: %w", err)
	}

	return string(buf.Bytes()), nil
}

func GenerateModelFields(val cue.Value) (string, error) {
	if err := val.Validate(); err != nil {
		return "", err
	}

	iter, err := val.Fields(
		cue.Definitions(true),
		cue.Optional(true),
	)
	if err != nil {
		return "", err
	}

	fields := make([]string, 0)
	for iter.Next() {
		if iter.IsDefinition() {
			continue
		}

		field := genSingleModelField(iter.Selector().String(), iter.Value())

		if field == "" {
			continue
		}

		fields = append(fields, field)
	}

	return strings.Join(fields, "\n"), nil
}

func genSingleModelField(name string, value cue.Value) string {
	goName := strings.Title(name)
	typeStr := TypeMappings[value.IncompleteKind()]

	// TODO: jduchesne, empty attribute type fails
	if typeStr == "" {
		return ""
	}

	return fmt.Sprintf("%s types.%s `tfsdk:\"%s\" json:\"%s\"`", goName, typeStr, ToSnakeCase(name), name)
}
