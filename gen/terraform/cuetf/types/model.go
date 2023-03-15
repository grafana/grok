package types

import (
	"fmt"
	"strings"

	"cuelang.org/go/cue"
	"github.com/grafana/grok/gen/terraform/cuetf/internal/utils"
)

type Model struct {
	Name   string
	Nodes  []Node
	Nested bool
}

// terraformModel generates the Terraform SDK model
func (s *Model) terraformModel() string {
	fields := make([]string, 0)
	if !s.Nested {
		fields = append(fields, "ToJSON types.String `tfsdk:\"to_json\"`")
	}

	for _, node := range s.Nodes {
		if !node.IsGenerated() {
			continue
		}
		fields = append(fields, node.TerraformModelField(s.Name))
	}

	return fmt.Sprintf(`type %s struct {
	%s
}
`, s.Name, strings.Join(fields, "\n"))
}

// jsonModel generates the JSON model used to convert the Terraform SDK model to JSON
func (s *Model) jsonModel() string {
	fields := make([]string, 0)
	for _, node := range s.Nodes {
		if !node.IsGenerated() {
			continue
		}
		fields = append(fields, node.JSONModelField())
	}

	return fmt.Sprintf(`type json%s struct {
	%s
}
`, s.Name, strings.Join(fields, "\n"))
}

// generateToJSONFunction generates a function that converts the Terraform SDK model to the JSON model representation
func (s *Model) generateToJSONFunction() string {
	content := fmt.Sprintf(`func (m %s) MarshalJSON() ([]byte, error) {
	%s
		`, s.Name, s.jsonModel())

	structContent := "\n	model := &json" + s.Name + "{\n"

	for _, node := range s.Nodes {
		identifier := "attr_" + strings.ToLower(node.Name)
		funcString := node.TerraformFunc()
		generated := false

		if node.Kind == cue.ListKind {
			subType := node.SubTerraformType()
			subTypeGolang := node.SubGolangType()
			subTypeFunc := node.SubTerraformFunc()
			if subType != "" {
				content += fmt.Sprintf("	%s := []%s{}\n", identifier, subTypeGolang)
				content += fmt.Sprintf("	for _, v := range m.%s.Elements() {\n", utils.ToCamelCase(node.Name))
				content += fmt.Sprintf("		%s = append(%s, v.(types.%s).%s)\n", identifier, identifier, subType, subTypeFunc)
				content += "	}\n"
				generated = true
			} else if node.SubKind == cue.StructKind {
				content += fmt.Sprintf("	%s := []interface{}{}\n", identifier)
				content += fmt.Sprintf("	for _, v := range m.%s {\n", utils.ToCamelCase(node.Name))
				content += fmt.Sprintf("		%s = append(%s, v)\n", identifier, identifier)
				content += "	}\n"
				generated = true
			}
		} else if node.Kind == cue.StructKind {
			if node.Optional {
				content += fmt.Sprintf("	var %s interface{}\n", identifier)
				content += fmt.Sprintf("	if m.%s != nil {\n", utils.ToCamelCase(node.Name))
				content += fmt.Sprintf("		%s = m.%s\n", identifier, utils.ToCamelCase(node.Name))
				content += "	}\n"
			} else {
				content += fmt.Sprintf("	var %s interface{} = m.%s\n", identifier, utils.ToCamelCase(node.Name))
			}
			generated = true
		} else if funcString != "" {
			content += fmt.Sprintf("	%s := m.%s.%s\n", identifier, utils.ToCamelCase(node.Name), funcString)
			if node.Optional {
				identifier = "&" + identifier
			}
			generated = true
		}

		if generated {
			structContent += fmt.Sprintf("		%s: %s,\n", utils.ToCamelCase(node.Name), identifier)
		}
	}

	content += structContent + `	}
	return json.Marshal(model)
}

`

	return content
}

func (s *Model) Generate() string {
	b := strings.Builder{}
	for _, node := range s.Nodes {
		if node.Kind == cue.StructKind || node.Kind == cue.ListKind && node.SubKind == cue.StructKind {
			nestedModel := Model{
				Name:   s.Name + "_" + strings.Title(node.Name),
				Nodes:  node.Children,
				Nested: true,
			}
			b.WriteString(nestedModel.Generate())
		}
	}

	b.WriteString(s.terraformModel())
	b.WriteString(s.generateToJSONFunction())

	return b.String()
}
