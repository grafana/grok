package cuetf

import (
	"bytes"
	"fmt"
	"go/format"
	"strings"

	"cuelang.org/go/cue"
	"github.com/grafana/grok/gen/terraform/cuetf/internal"
	"github.com/grafana/grok/gen/terraform/cuetf/internal/utils"
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

	schemaAttributes, err := GenerateSchemaAttributes(nodes)
	if err != nil {
		return nil, err
	}

	modelFields, err := GenerateModelFields(nodes)
	if err != nil {
		return nil, err
	}

	defaults, err := GenerateDefaults(nodes, []string{})
	if err != nil {
		return nil, err
	}

	vars := TVarsDataSource{
		Name:             strings.Title(schema.Lineage().Name()),
		Description:      "TODO description",
		ModelFields:      modelFields,
		SchemaAttributes: strings.Join(schemaAttributes, "\n"),
		Defaults:         defaults,
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

func GenerateSchemaAttributes(nodes []types.Node) ([]string, error) {
	attributes := make([]string, 0)
	for _, node := range nodes {
		vars := TVarsSchemaAttribute{
			Name:          utils.ToSnakeCase(node.Name),
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
					return nil, fmt.Errorf("error trying to generate nested attributes in list: %s", err)
				}
				vars.NestedObjectAttributes = strings.Join(nestedObjectAttributes, "")
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
				return nil, fmt.Errorf("error trying to generate nested attributes in struct: %w", err)
			}
			vars.NestedAttributes = strings.Join(nestedAttributes, "")
		}

		// TODO: fixme
		if vars.AttributeType == "" {
			continue
		}

		buf := new(bytes.Buffer)
		if err := tmpls.Lookup("schema_attribute.tmpl").Execute(buf, vars); err != nil {
			return nil, fmt.Errorf("failed executing datasource template: %w", err)
		}

		attributes = append(attributes, string(buf.Bytes()))
	}

	return attributes, nil
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
			typeStr = "*struct{\n"
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

		fields = append(fields, fmt.Sprintf("%s %s `tfsdk:\"%s\" json:\"%s\"`", utils.ToCamelCase(node.Name), typeStr, utils.ToSnakeCase(node.Name), node.Name))
	}

	return strings.Join(fields, "\n"), nil
}

func GenerateDefaults(nodes []types.Node, parents []string) (string, error) {
	defaults := make([]string, 0)
	for _, node := range nodes {
		kind := TypeMappings[node.Kind]

		if kind != "" && node.Default != "" {
			path := utils.ToCamelCase(node.Name)
			if len(parents) > 0 {
				path = strings.Join(parents, ".") + "." + path
			}

			// TODO: We check if all parent structs are not nil but maybe we should initialise them if they are
			nullFieldConditions := make([]string, 0)
			for i := range parents {
				fields := strings.Join(parents[:i+1], ".")
				nullFieldConditions = append(nullFieldConditions, fmt.Sprintf("data.%s != nil", fields))
			}
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

		// TODO: handle need separately, by adding builder functions?
		if node.Kind != cue.ListKind && len(node.Children) != 0 {
			parentsCopy := parents
			parentsCopy = append(parentsCopy, utils.ToCamelCase(node.Name))
			nestedDefaults, err := GenerateDefaults(node.Children, parentsCopy)
			if err != nil {
				return "", fmt.Errorf("error generating nested defaults: %w", err)
			}
			defaults = append(defaults, nestedDefaults)
		}
	}

	return strings.Join(defaults, ""), nil
}
