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
	b := strings.Builder{}

	fmt.Fprintf(&b, "func (m %s) MarshalJSON() ([]byte, error) {\n", s.Name)
	b.WriteString(s.jsonModel() + "\n")

	structLines := make([]string, 0)
	for _, node := range s.Nodes {
		if !node.IsGenerated() {
			continue
		}

		identifier := "attr_" + strings.ToLower(node.Name)
		funcString := node.TerraformFunc()

		if node.Kind == cue.ListKind {
			subType := node.SubTerraformType()
			subTypeGolang := node.SubGolangType()
			subTypeFunc := node.SubTerraformFunc()
			if subType != "" {
				fmt.Fprintf(&b, "	%s := []%s{}\n", identifier, subTypeGolang)
				fmt.Fprintf(&b, "	for _, v := range m.%s.Elements() {\n", utils.ToCamelCase(node.Name))
				fmt.Fprintf(&b, "		%s = append(%s, v.(types.%s).%s)\n", identifier, identifier, subType, subTypeFunc)
				b.WriteString("	}\n")
			} else if node.SubKind == cue.StructKind {
				fmt.Fprintf(&b, "	%s := []interface{}{}\n", identifier)
				fmt.Fprintf(&b, "	for _, v := range m.%s {\n", utils.ToCamelCase(node.Name))
				fmt.Fprintf(&b, "		%s = append(%s, v)\n", identifier, identifier)
				b.WriteString("	}\n")
			}
		} else if node.Kind == cue.StructKind {
			if node.Optional {
				fmt.Fprintf(&b, "	var %s interface{}\n", identifier)
				fmt.Fprintf(&b, "	if m.%s != nil {\n", utils.ToCamelCase(node.Name))
				fmt.Fprintf(&b, "		%s = m.%s\n", identifier, utils.ToCamelCase(node.Name))
				b.WriteString("	}\n")
			} else {
				fmt.Fprintf(&b, "	var %s interface{} = m.%s\n", identifier, utils.ToCamelCase(node.Name))
			}
		} else if funcString != "" {
			fmt.Fprintf(&b, "	%s := m.%s.%s\n", identifier, utils.ToCamelCase(node.Name), funcString)
			if node.Optional {
				identifier = "&" + identifier
			}
		}

		structLines = append(structLines, fmt.Sprintf("		%s: %s,\n", utils.ToCamelCase(node.Name), identifier))
	}

	fmt.Fprintf(&b, `
	model := &json%s {
%s
	}
	return json.Marshal(model)
}
`, s.Name, strings.Join(structLines, ""))

	return b.String()
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
