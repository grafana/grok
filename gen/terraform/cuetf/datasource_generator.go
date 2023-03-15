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
		// The common schema has an `options` field that is empty
		// and the panel schema has a `panelOptions` field that is supposed to be used in the `options` json attribute
		for _, node := range panelNodes {
			if node.Name == "options" {
				continue
			}
			nodes = append(nodes, node)
		}
		for i, node := range nodes {
			if node.Name == "PanelOptions" {
				node.Name = "options"
				nodes[i] = node
			}
			if node.Name == "type" {
				panelType := strings.ToLower(strings.TrimPrefix(GetKindName(linName), "Panel")) // TODO: Better way to get panel type?
				node.Default = fmt.Sprintf("`%s`", panelType)
				nodes[i] = node
			}
		}
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
		Defaults:         defaults,
	}

	buf := new(bytes.Buffer)
	if err := tmpls.Lookup("datasource.tmpl").Execute(buf, vars); err != nil {
		return nil, fmt.Errorf("failed executing datasource template: %w", err)
	}

	// return buf.Bytes(), nil

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
