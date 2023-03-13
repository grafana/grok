package cuetf

import (
	"bytes"
	"errors"
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
		Name:             strings.Title(schema.Lineage().Name()),
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
	if name == "panels" {
		return "", nil
	}

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
	kind := value.IncompleteKind()
	vars.AttributeType = TypeMappings[kind]
	switch kind {
	case cue.ListKind:
		defv, _ := value.Default()
		if !defv.Equals(value) {
			_, v := value.Expr()
			value = v[0]
		}

		e := value.LookupPath(cue.MakePath(cue.AnyIndex))
		if e.Exists() {
			subType := TypeMappings[e.IncompleteKind()]

			if subType != "" {
				// "example_attribute": schema.ListAttribute{
				// 		ElementType: types.StringType,
				// 	    // ... other fields ...
				// },
				vars.AttributeType = "List"
				vars.ElementType = fmt.Sprintf("types.%sType", subType)
			} else {
				// "nested_attribute": schema.ListNestedAttribute{
				//     NestedObject: schema.NestedAttributeObject{
				//         Attributes: map[string]schema.Attribute{
				//             "hello": schema.StringAttribute{
				//                 /* ... */
				//             },
				//         },
				//     },
				// },
				vars.AttributeType = "ListNested"
				nestedObjectAttributes, err := GenerateSchemaAttributes(e)
				if err != nil {
					return "", err
				}
				vars.NestedObjectAttributes = nestedObjectAttributes
			}
		} else {
			return "", errors.New("unreachable - open list must have a type")
		}
	case cue.StructKind:
		// "nested_attribute": schema.SingleNestedAttribute{
		//     Attributes: map[string]schema.Attribute{
		//         "hello": schema.StringAttribute{
		//             /* ... */
		//         },
		//     },
		// },
		vars.AttributeType = "SingleNested"
		nestedAttributes, err := GenerateSchemaAttributes(value.Value())
		if err != nil {
			return "", err
		}
		vars.NestedAttributes = nestedAttributes
	}

	// TODO Remove
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

		field, err := genSingleModelField(iter.Selector().String(), iter.Value())
		if err != nil {
			return "", err
		}

		if field == "" {
			continue
		}

		fields = append(fields, field)
	}

	return strings.Join(fields, "\n"), nil
}

func genSingleModelField(name string, value cue.Value) (string, error) {
	if name == "panels" {
		return "", nil
	}

	goName := strings.Title(name)

	kind := value.IncompleteKind()
	typeStr := "types." + TypeMappings[kind]
	switch kind {
	case cue.ListKind:
		defv, _ := value.Default()
		if !defv.Equals(value) {
			_, v := value.Expr()
			value = v[0]
		}

		e := value.LookupPath(cue.MakePath(cue.AnyIndex))
		if e.Exists() {
			subType := TypeMappings[e.IncompleteKind()]
			if subType != "" {
				typeStr = "types.List"
			} else {
				typeStr = "[]struct{\n"
				nestedAttributes, err := GenerateModelFields(e)
				if err != nil {
					return "", err
				}
				typeStr += nestedAttributes + "\n}"
			}
		} else {
			return "", errors.New("unreachable - open list must have a type")
		}
	case cue.StructKind:
		typeStr = "*struct{\n"
		nestedAttributes, err := GenerateModelFields(value.Value())
		if err != nil {
			return "", err
		}
		typeStr += nestedAttributes + "\n}"
	}

	// TODO remove
	if typeStr == "" || typeStr == "types." {
		return "", nil
	}

	return fmt.Sprintf("%s %s `tfsdk:\"%s\" json:\"%s\"`", goName, typeStr, ToSnakeCase(name), name), nil
}
