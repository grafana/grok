package cuetf

import (
	"bytes"
	"errors"
	"fmt"
	"go/format"
	"regexp"
	"strings"

	"cuelang.org/go/cue"
	"github.com/grafana/grok/gen/terraform/cuetf/internal"
	"github.com/grafana/grok/gen/terraform/cuetf/internal/utils"
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
	model := types.Model{
		Name:  structName + "Model",
		Nodes: nodes,
	}

	vars := TVarsDataSource{
		Name:             GetResourceName(linName),
		StructName:       structName,
		Description:      "TODO description",
		Models:           model.Generate(),
		SchemaAttributes: schemaAttributes,
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
		if !node.IsGenerated() {
			continue
		}

		description := node.Doc
		if node.Default != "" {
			if description != "" && !strings.HasSuffix(description, ".") {
				description += "."
			}
			description += " Defaults to " + strings.ReplaceAll(node.Default, "`", `"`) + "."
		}

		deprecated := deprecationMessage(node.Doc)

		vars := TVarsSchemaAttribute{
			Name:               utils.ToSnakeCase(node.Name),
			Description:        description,
			DeprecationMessage: deprecated,
			AttributeType:      node.TerraformType(),
			Computed:           false,
			Optional:           node.Optional,
		}

		if node.Default != "" {
			vars.Optional = true
			vars.Computed = true
		}

		switch node.Kind {
		case cue.ListKind:
			subType := node.SubTerraformType()
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

		buf := new(bytes.Buffer)
		if err := tmpls.Lookup("schema_attribute.tmpl").Execute(buf, vars); err != nil {
			return "", fmt.Errorf("failed executing datasource template: %w", err)
		}

		attributes = append(attributes, buf.String())
	}

	return strings.Join(attributes, ""), nil
}

var deprecatedMatch = regexp.MustCompile(`^\W*@deprecated\W*`)

func deprecationMessage(str string) string {
	if !strings.Contains(str, "@deprecated") {
		return str
	}

	deprecated := deprecatedMatch.ReplaceAllString(str, "")
	return utils.CapitalizeFirstLetter(deprecated)
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

func GetKindName(rawName string) string {
	name := rawName
	if strings.HasSuffix(name, "PanelCfg") {
		name = "Panel" + strings.TrimSuffix(name, "PanelCfg")
	} else if strings.HasSuffix(name, "DataQuery") {
		name = "Query" + strings.TrimSuffix(name, "DataQuery")
	} else {
		switch name {
		case "dashboard", "playlist", "preferences", "team":
			name = strings.ToUpper(name[:1]) + name[1:]
		case "publicdashboard":
			name = "PublicDashboard"
		case "librarypanel":
			name = "LibraryPanel"
		case "serviceaccount":
			name = "ServiceAccount"
		}
		name = "Core" + name
	}

	return name
}

func GetStructName(rawName string) string {
	return strings.Title(GetKindName(rawName)) + "DataSource"
}

func GetResourceName(rawName string) string {
	return utils.ToSnakeCase(GetKindName(rawName))
}
