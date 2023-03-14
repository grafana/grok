package cuetf

import (
	"bytes"
	"errors"
	"fmt"
	"go/format"
	"strings"

	"cuelang.org/go/cue"
	"github.com/grafana/grok/gen/terraform/cuetf/internal"
	"github.com/grafana/grok/gen/terraform/cuetf/types"
	"github.com/grafana/thema"
	"golang.org/x/tools/imports"
)

// GenerateDataSource takes a cue.Value and generates the corresponding Terraform data source
func GenerateDataSource(schema thema.Schema) (b []byte, err error) {
	nodes, err := internal.GetAllNodes(schema.Underlying())
	if err != nil {
		return nil, err
	}

	if err := extractPanelNodes(schema); err != nil {
		return nil, err
	}

	linName := schema.Lineage().Name()
	if strings.HasPrefix(GetKindName(linName), "Panel") {
		if len(panelNodes) == 0 {
			return nil, errors.New("panel schema not found")
		}
		nodes = append(nodes, panelNodes...)
	}

	schemaAttributes, err := GenerateSchemaAttributes(nodes)
	if err != nil {
		return nil, err
	}

	modelFields, err := GenerateModelFields(nodes)
	if err != nil {
		return nil, err
	}

	jsonModelFields, err := GenerateJSONModelFields(nodes)
	if err != nil {
		return nil, err
	}

	defaults, err := GenerateDefaults(nodes, []types.Node{})
	if err != nil {
		return nil, err
	}

	structName := GetStructName(linName)
	vars := TVarsDataSource{
		Name:              GetResourceName(linName),
		StructName:        structName,
		Description:       "TODO description",
		ModelFields:       modelFields,
		JSONModelFields:   jsonModelFields,
		TFModelToJSONFunc: GenerateToJSONFunction(structName, nodes),
		SchemaAttributes:  schemaAttributes,
		Defaults:          defaults,
	}

	buf := new(bytes.Buffer)
	if err := tmpls.Lookup("datasource.tmpl").Execute(buf, vars); err != nil {
		return nil, fmt.Errorf("failed executing datasource template: %w", err)
	}

	// Add import if needed - for now it should only add "math/big"
	// if there is number attributes with defaults
	byt, err := imports.Process("", buf.Bytes(), nil)
	if err != nil {
		return nil, fmt.Errorf("goimports processing of generated file failed: %w", err)
	}

	return format.Source(byt)
}

func GenerateSchemaAttributes(nodes []types.Node) (string, error) {
	attributes := make([]string, 0)
	for _, node := range nodes {
		vars := TVarsSchemaAttribute{
			Name:          ToSnakeCase(node.Name),
			Description:   node.Doc,
			AttributeType: TypeMappings[node.Kind],
			Computed:      false,
			Optional:      node.Optional,
		}

		if node.Default != "" {
			vars.Optional = true
			vars.Computed = true
		}

		vars.Required = !vars.Optional

		switch node.Kind {
		case cue.ListKind:
			subType := TypeMappings[node.SubKind]
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
				nestedObjectAttributes, err := GenerateSchemaAttributes(node.Children)
				if err != nil {
					return "", fmt.Errorf("error trying to generate nested attributes in list: %s", err)
				}
				vars.NestedObjectAttributes = nestedObjectAttributes
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
			nestedAttributes, err := GenerateSchemaAttributes(node.Children)
			if err != nil {
				return "", fmt.Errorf("error trying to generate nested attributes in struct: %w", err)
			}
			vars.NestedAttributes = nestedAttributes
		}

		// TODO: fixme
		if vars.AttributeType == "" {
			continue
		}

		buf := new(bytes.Buffer)
		if err := tmpls.Lookup("schema_attribute.tmpl").Execute(buf, vars); err != nil {
			return "", fmt.Errorf("failed executing datasource template: %w", err)
		}

		attributes = append(attributes, string(buf.Bytes()))
	}

	return strings.Join(attributes, ""), nil
}

func GenerateModelFields(nodes []types.Node) (string, error) {
	fields := make([]string, 0)
	for _, node := range nodes {
		typeStr := "types." + TypeMappings[node.Kind]
		switch node.Kind {
		case cue.ListKind:
			subType := TypeMappings[node.SubKind]
			if subType != "" {
				typeStr = "types.List"
			} else {
				typeStr = "[]struct{\n"
				nestedAttributes, err := GenerateModelFields(node.Children)
				if err != nil {
					return "", err
				}
				typeStr += nestedAttributes + "\n}"
			}
		case cue.StructKind:
			// If not optional, no need to be a pointer
			typeStr = "struct{\n"
			if node.Optional {
				typeStr = "*" + typeStr
			}
			nestedAttributes, err := GenerateModelFields(node.Children)
			if err != nil {
				return "", err
			}
			typeStr += nestedAttributes + "\n}"
		}

		// TODO: fixme
		if typeStr == "types." {
			continue
		}

		fields = append(fields, fmt.Sprintf("%s %s `tfsdk:\"%s\"`", ToCamelCase(node.Name), typeStr, ToSnakeCase(node.Name)))
	}

	return strings.Join(fields, "\n"), nil
}

func GenerateJSONModelFields(nodes []types.Node) (string, error) {
	fields := make([]string, 0)
	for _, node := range nodes {
		typeStr := GolangTypeMappings[node.Kind]
		switch node.Kind {
		case cue.ListKind:
			subType := GolangTypeMappings[node.SubKind]
			if subType != "" {
				typeStr = "[]" + subType
			} else {
				typeStr = "[]struct{\n"
				nestedAttributes, err := GenerateJSONModelFields(node.Children)
				if err != nil {
					return "", err
				}
				typeStr += nestedAttributes + "\n}"
			}
		case cue.StructKind:
			// If not optional, no need to be a pointer
			typeStr = "struct{\n"
			nestedAttributes, err := GenerateJSONModelFields(node.Children)
			if err != nil {
				return "", err
			}
			typeStr += nestedAttributes + "\n}"
		default:

		}

		// TODO: fixme
		if typeStr == "" {
			continue
		}

		omitStr := ""
		if node.Optional {
			if !strings.HasPrefix(typeStr, "[]") {
				typeStr = "*" + typeStr
			}
			omitStr = ",omitempty"
		}

		fields = append(fields, fmt.Sprintf("%s %s `json:\"%s%s\"`", ToCamelCase(node.Name), typeStr, node.Name, omitStr))
	}

	return strings.Join(fields, "\n"), nil
}

// GenerateToJSONFunction generates a function that converts the Terraform SDK model to the JSON model representation
func GenerateToJSONFunction(structName string, nodes []types.Node) string {
	content := fmt.Sprintf("func (d %[1]sModel) MarshalJSON() ([]byte, error) {\n", structName)
	returnContent := "\n	model := &" + structName + "ModelJSON{\n"

	for _, node := range nodes {
		funcString := TerraformFuncTypeMappings[node.Kind]
		if funcString == "" {
			continue
		}

		identifier := "attr_" + strings.ToLower(node.Name)
		content += fmt.Sprintf("	%[1]s := d.%[2]s.%[3]s\n", identifier, ToCamelCase(node.Name), funcString)

		ref := ""
		if node.Optional {
			ref = "&"
		}
		returnContent += fmt.Sprintf("		%[1]s: %s%s,\n", ToCamelCase(node.Name), ref, identifier)
	}

	content += returnContent + `	}
	return json.Marshal(model)
}`

	return content
}

func GenerateDefaults(nodes []types.Node, parents []types.Node) (string, error) {
	defaults := make([]string, 0)
	for _, node := range nodes {
		kind := TypeMappings[node.Kind]

		if kind != "" && node.Default != "" {
			path := ""
			// TODO: We check if all parent structs are not nil but maybe we should initialise them if they are
			nullFieldConditions := make([]string, 0)
			for _, parent := range parents {
				path = path + ToCamelCase(parent.Name)
				if parent.Optional {
					nullFieldConditions = append(nullFieldConditions, fmt.Sprintf("data.%s != nil", path))
				}
				path += "."
			}
			path += ToCamelCase(node.Name)
			nullFieldConditions = append(nullFieldConditions, fmt.Sprintf("data.%s.IsNull()", path))

			vars := TVarsDefault{
				Name:               path,
				NullFieldCondition: strings.Join(nullFieldConditions, " && "),
				Type:               kind,
				Default:            node.Default,
			}

			buf := new(bytes.Buffer)
			if err := tmpls.Lookup("default.tmpl").Execute(buf, vars); err != nil {
				return "", fmt.Errorf("failed executing datasource template: %w", err)
			}

			defaults = append(defaults, string(buf.Bytes()))
		}

		// TODO: handle lists separately, by adding builder functions?
		if node.Kind != cue.ListKind && len(node.Children) != 0 {
			parentsCopy := parents
			parentsCopy = append(parentsCopy, node)
			nestedDefaults, err := GenerateDefaults(node.Children, parentsCopy)
			if err != nil {
				return "", fmt.Errorf("error generating nested defaults: %w", err)
			}
			defaults = append(defaults, nestedDefaults)
		}
	}

	return strings.Join(defaults, ""), nil
}

var panelNodes []types.Node

func extractPanelNodes(schema thema.Schema) error {
	if schema.Lineage().Name() == "dashboard" {
		iter, err := schema.Underlying().Fields(
			cue.Definitions(true),
			cue.Optional(false),
			cue.Attributes(false),
		)
		if err != nil {
			return err
		}
		for iter.Next() {
			if iter.Selector().String() == "#Panel" {
				if panelNodes, err = internal.GetAllNodes(iter.Value()); err != nil {
					return err
				}
				break
			}
		}
	}
	return nil
}
