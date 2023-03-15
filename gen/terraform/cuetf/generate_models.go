package cuetf

import (
	"fmt"
	"strings"

	"cuelang.org/go/cue"
	"github.com/grafana/grok/gen/terraform/cuetf/types"
)

func generateModelFields(structName string, fieldNodes []types.Node) (string, error) {
	fields := make([]string, 0)
	for _, node := range fieldNodes {
		typeStr := "types." + TypeMappings[node.Kind]
		switch node.Kind {
		case cue.ListKind:
			subType := TypeMappings[node.SubKind]
			if subType != "" {
				typeStr = "types.List"
			} else if node.SubKind == cue.StructKind {
				typeStr = "[]" + structName + "_" + strings.Title(node.Name)
			}
		case cue.StructKind:
			// If not optional, no need to be a pointer
			typeStr = structName + "_" + strings.Title(node.Name)
			if node.Optional {
				typeStr = "*" + typeStr
			}
		}

		// TODO: fixme
		if typeStr == "types." {
			continue
		}

		fields = append(fields, fmt.Sprintf("%s %s `tfsdk:\"%s\"`", ToCamelCase(node.Name), typeStr, ToSnakeCase(node.Name)))
	}

	return strings.Join(fields, "\n"), nil
}

func generateJSONModelFields(structName string, fieldNodes []types.Node) (string, error) {
	fields := make([]string, 0)
	for _, node := range fieldNodes {
		typeStr := GolangTypeMappings[node.Kind]
		switch node.Kind {
		case cue.ListKind:
			subType := GolangTypeMappings[node.SubKind]
			if subType != "" {
				typeStr = "[]" + subType
			} else if node.SubKind == cue.StructKind {
				typeStr = "[]interface{}"
			}
		case cue.StructKind:
			typeStr = "interface{}"
		}

		// TODO: fixme
		if typeStr == "" {
			continue
		}

		omitStr := ""
		if node.Optional {
			if !strings.HasPrefix(typeStr, "[]") && typeStr != "interface{}" {
				typeStr = "*" + typeStr
			}
			omitStr = ",omitempty"
		}

		fields = append(fields, fmt.Sprintf("		%s %s `json:\"%s%s\"`", ToCamelCase(node.Name), typeStr, node.Name, omitStr))
	}

	return strings.Join(fields, "\n"), nil
}

// generateToJSONFunction generates a function that converts the Terraform SDK model to the JSON model representation
func generateToJSONFunction(structName string, nodes []types.Node) (string, error) {
	jsonModelFields, err := generateJSONModelFields(structName, nodes)
	if err != nil {
		return "", err
	}

	content := fmt.Sprintf(`func (m %[1]s) MarshalJSON() ([]byte, error) {
	type json%[1]s struct {
%[2]s
	}
		`, structName, jsonModelFields)

	structContent := "\n	model := &json" + structName + "{\n"

	for _, node := range nodes {
		identifier := "attr_" + strings.ToLower(node.Name)
		funcString := TerraformFuncTypeMappings[node.Kind]
		generated := false

		if node.Kind == cue.ListKind {
			subType := TypeMappings[node.SubKind]
			subTypeGolang := GolangTypeMappings[node.SubKind]
			subTypeFunc := TerraformFuncTypeMappings[node.SubKind]
			if subType != "" {
				content += fmt.Sprintf("	%s := []%s{}\n", identifier, subTypeGolang)
				content += fmt.Sprintf("	for _, v := range m.%s.Elements() {\n", ToCamelCase(node.Name))
				content += fmt.Sprintf("		%s = append(%s, v.(types.%s).%s)\n", identifier, identifier, subType, subTypeFunc)
				content += "	}\n"
				generated = true
			} else if node.SubKind == cue.StructKind {
				content += fmt.Sprintf("	%s := []interface{}{}\n", identifier)
				content += fmt.Sprintf("	for _, v := range m.%s {\n", ToCamelCase(node.Name))
				content += fmt.Sprintf("		%s = append(%s, v)\n", identifier, identifier)
				content += "	}\n"
				generated = true
			}
		} else if node.Kind == cue.StructKind {
			if node.Optional {
				content += fmt.Sprintf("	var %s interface{}\n", identifier)
				content += fmt.Sprintf("	if m.%s != nil {\n", ToCamelCase(node.Name))
				content += fmt.Sprintf("		%s = m.%s\n", identifier, ToCamelCase(node.Name))
				content += "	}\n"
			} else {
				content += fmt.Sprintf("	var %s interface{} = m.%s\n", identifier, ToCamelCase(node.Name))
			}
			generated = true
		} else if funcString != "" {
			content += fmt.Sprintf("	%s := m.%s.%s\n", identifier, ToCamelCase(node.Name), funcString)
			if node.Optional {
				identifier = "&" + identifier
			}
			generated = true
		}

		if generated {
			structContent += fmt.Sprintf("		%s: %s,\n", ToCamelCase(node.Name), identifier)
		}
	}

	content += structContent + `	}
	return json.Marshal(model)
}

`

	return content, nil
}

func GenerateModels(structName string, fieldNodes []types.Node, top bool) (string, error) {
	b := strings.Builder{}
	for _, node := range fieldNodes {
		if node.Kind == cue.StructKind || node.Kind == cue.ListKind && node.SubKind == cue.StructKind {
			nestedModel, err := GenerateModels(structName+"_"+strings.Title(node.Name), node.Children, false)
			if err != nil {
				return "", err
			}
			b.WriteString(nestedModel)
		}
	}

	fmt.Fprintf(&b, "type %s struct {\n", structName)
	fields, err := generateModelFields(structName, fieldNodes)
	if err != nil {
		return "", err
	}
	if top {
		b.WriteString("	ToJSON types.String `tfsdk:\"to_json\"`\n")
	}
	fmt.Fprintf(&b, "%s\n}\n\n", fields)

	tfToJSON, err := generateToJSONFunction(structName, fieldNodes)
	if err != nil {
		return "", err
	}
	b.WriteString(tfToJSON)

	return b.String(), nil
}
