package cuetf

import (
	"bytes"
	"errors"
	"fmt"
	"go/format"
	"strings"

	"cuelang.org/go/cue"
	"github.com/grafana/thema"
)

type cueField struct {
	Name       string
	Value      cue.Value
	IsOptional bool
}

func schemaToCueFields(schema cue.Value) ([]cueField, error) {
	if !schema.IsConcrete() {
		return nil, nil
	}

	fields := []cueField{}
	iter, err := schema.Fields(
		cue.Definitions(false),
		cue.Optional(true),
	)
	if err != nil {
		return nil, fmt.Errorf("error retrieving value fields: %w", err)
	}
	for iter.Next() {
		fields = append(fields, cueField{
			Name:       iter.Selector().String(),
			Value:      iter.Value(),
			IsOptional: iter.IsOptional(),
		})
	}
	return fields, nil
}

func GetStructName(kindName string) string {
	return strings.Title(kindName) + "DataSource"
}

// GenerateDataSource takes a cue.Value and generates the corresponding Terraform data source
func GenerateDataSource(schema thema.Schema) (b []byte, err error) {
	kindName := schema.Lineage().Name()
	if schema.Underlying().Validate() != nil {
		return nil, fmt.Errorf("error validating schema: %w", err)
	}

	fields, err := schemaToCueFields(schema.Underlying())
	if err != nil {
		return nil, err
	}

	extractPanelSchema(schema)

	if strings.HasSuffix(kindName, "PanelCfg") {
		if !panelSchema.Exists() {
			return nil, errors.New("panel schema not found")
		}
		panelFields, err := schemaToCueFields(panelSchema)
		if err != nil {
			return nil, err
		}
		fields = append(fields, panelFields...)
	}

	schemaAttributes, err := GenerateSchemaAttributes(fields)
	if err != nil {
		return nil, err
	}

	modelFields, err := GenerateModelFields(fields)
	if err != nil {
		return nil, err
	}

	vars := TVarsDataSource{
		Name:             kindName,
		StructName:       GetStructName(kindName),
		Description:      "TODO description",
		ModelFields:      modelFields,
		SchemaAttributes: string(schemaAttributes),
	}

	buf := new(bytes.Buffer)
	if err := tmpls.Lookup("datasource.tmpl").Execute(buf, vars); err != nil {
		return nil, fmt.Errorf("failed executing datasource template: %w", err)
	}

	// if err := os.MkdirAll("/Users/julienduchesne/Repos/terraform-provider-schemas/tools/grok/terraform/debug", 0755); err != nil {
	// 	return nil, fmt.Errorf("failed creating debug directory: %w", err)
	// }
	// if err := os.WriteFile("/Users/julienduchesne/Repos/terraform-provider-schemas/tools/grok/terraform/debug/"+kindName+".go", buf.Bytes(), 0644); err != nil {
	// 	return nil, fmt.Errorf("failed writing debug file: %w", err)
	// }

	return format.Source(buf.Bytes())
}

func GenerateSchemaAttributesFromSchema(val cue.Value) (string, error) {
	fields, err := schemaToCueFields(val)
	if err != nil {
		return "", err
	}

	return GenerateSchemaAttributes(fields)
}

func GenerateSchemaAttributes(cueFields []cueField) (string, error) {
	attributes := make([]string, 0)
	for _, cueField := range cueFields {
		attr, err := genSingleSchemaAttribute(cueField.Name, cueField.Value, cueField.IsOptional)
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
	if name == "panels" || name == "points" || name == "bucketAggs" || name == "metrics" {
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
	vars.Description = strings.ReplaceAll(strings.Trim(vars.Description, "\n "), "`", "")

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
				nestedObjectAttributes, err := GenerateSchemaAttributesFromSchema(e)
				if err != nil {
					return "", fmt.Errorf("error trying to generate nested attributes in list: %s", err)
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
		nestedAttributes, err := GenerateSchemaAttributesFromSchema(value.Value())
		if err != nil {
			return "", fmt.Errorf("error trying to generate nested attributes in struct: %w", err)
		}
		vars.NestedAttributes = nestedAttributes
	}

	// TODO Remove
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

func GenerateModelFieldsFromSchema(val cue.Value) (string, error) {
	fields, err := schemaToCueFields(val)
	if err != nil {
		return "", err
	}

	return GenerateModelFields(fields)
}

func GenerateModelFields(cueFields []cueField) (string, error) {
	fields := make([]string, 0)
	for _, cueField := range cueFields {
		field, err := genSingleModelField(cueField.Name, cueField.Value)
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
	if name == "panels" || name == "points" || name == "bucketAggs" || name == "metrics" {
		return "", nil
	}

	goName := ToCamelCase(name)

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
				nestedAttributes, err := GenerateModelFieldsFromSchema(e)
				if err != nil {
					return "", err
				}
				typeStr += nestedAttributes + "\n}"
			}
		} else {
			return "", errors.New("unreachable - open list must have a type")
		}
	case cue.StructKind:
		// If not optional, no need to be a pointer
		typeStr = "*struct{\n"
		nestedAttributes, err := GenerateModelFieldsFromSchema(value.Value())
		if err != nil {
			return "", err
		}
		typeStr += nestedAttributes + "\n}"
	}

	// TODO: jduchesne, empty attribute type fails
	if typeStr == "" || typeStr == "types." {
		return "", nil
	}

	return fmt.Sprintf("%s %s `tfsdk:\"%s\" json:\"%s\"`", goName, typeStr, ToSnakeCase(name), name), nil
}

var panelSchema cue.Value

func extractPanelSchema(schema thema.Schema) {
	if schema.Lineage().Name() == "dashboard" {
		iter, _ := schema.Underlying().Fields(
			cue.Definitions(true),
			cue.Optional(false),
			cue.Attributes(false),
		)
		for iter.Next() {
			if iter.Selector().String() == "#Panel" {
				panelSchema = iter.Value()
				break
			}
		}
	}

}
