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
		// The common schema has an `options` field that is empty and overriden by `panelOptions` in the panel schema
		// and a `fieldConfig` field that should contain a `custom` field that contains the panel schema `panelFieldConfig` nodes
		// It seems like all other fields in the panel schema should be definitions
		var panelOptions *types.Node
		var panelFieldConfig *types.Node
		for i, node := range nodes {
			if node.Name == "PanelOptions" {
				nodes[i].Name = "options"
				panelOptions = &nodes[i]
			} else if node.Name == "PanelFieldConfig" {
				nodes[i].Name = "custom"
				panelFieldConfig = &nodes[i]
			}
		}

		if len(panelNodes) == 0 {
			return nil, errors.New("panel schema not found")
		}

		for i, node := range panelNodes {
			if node.Name == "options" && panelOptions != nil {
				panelNodes[i] = *panelOptions
			}

			if node.Name == "fieldConfig" && panelFieldConfig != nil {
				for _, n1 := range node.Children {
					if n1.Name != "defaults" {
						continue
					}

					for j, n2 := range n1.Children {
						if n2.Name == "custom" {
							n1.Children[j] = *panelFieldConfig
						}
					}
				}
			}

			// TODO: set it as read-only?
			if node.Name == "type" {
				panelType := strings.ToLower(strings.TrimPrefix(GetKindName(linName), "Panel")) // TODO: Better way to get panel type?
				node.Default = fmt.Sprintf("`%s`", panelType)
				panelNodes[i] = node
			}
		}

		nodes = panelNodes
	}

	schemaAttributes, err := GenerateSchemaAttributes(nodes)
	if err != nil {
		return nil, err
	}

	structName := GetStructName(linName)
	models, err := GenerateModels(structName+"Model", nodes, true)
	if err != nil {
		return nil, err
	}

	initStructs := InitStructs(structName+"Model", nodes, []types.Node{})
	defaults, err := GenerateDefaults(nodes, []types.Node{})
	if err != nil {
		return nil, err
	}

	vars := TVarsDataSource{
		Name:             GetResourceName(linName),
		StructName:       structName,
		Description:      "TODO description",
		Models:           models,
		SchemaAttributes: schemaAttributes,
		Defaults:         initStructs + defaults,
	}

	buf := new(bytes.Buffer)
	if err := tmpls.Lookup("datasource.tmpl").Execute(buf, vars); err != nil {
		return nil, fmt.Errorf("failed executing datasource template: %w", err)
	}

	// return buf.Bytes(), nil
	return format.Source(buf.Bytes())
}

func GenerateSchemaAttributes(nodes []types.Node) (string, error) {
	attributes := make([]string, 0)
	for _, node := range nodes {
		description := node.Doc
		if node.Default != "" {
			if description != "" && !strings.HasSuffix(description, ".") {
				description += "."
			}
			description += " Defaults to " + strings.ReplaceAll(node.Default, "`", `"`) + "."
		}

		vars := TVarsSchemaAttribute{
			Name:          ToSnakeCase(node.Name),
			Description:   description,
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
			} else if node.SubKind == cue.StructKind {
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

			// Structs should be computed if we want to set nested defaults?
			vars.Computed = true
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

func InitStructs(structName string, nodes []types.Node, parents []types.Node) string {
	init := ""
	for _, node := range nodes {
		if node.Kind == cue.StructKind && node.Optional {
			fieldType := structName
			path := ""
			for _, parent := range parents {
				fieldType = fieldType + "_" + ToCamelCase(parent.Name)
				path += ToCamelCase(parent.Name) + "."
			}
			path += ToCamelCase(node.Name)
			fieldType += "_" + ToCamelCase(node.Name)

			init += fmt.Sprintf("if data.%s == nil {\n", path)
			init += fmt.Sprintf("data.%s = &%s{}\n", path, fieldType)
			init += fmt.Sprintln("}")
		}

		if node.Kind != cue.ListKind {
			parentsCopy := parents
			parentsCopy = append(parentsCopy, node)
			init += InitStructs(structName, node.Children, parentsCopy)
		}
	}

	return init
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
