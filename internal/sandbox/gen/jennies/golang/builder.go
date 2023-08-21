package golang

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/grafana/codejen"
	"github.com/grafana/grok/internal/sandbox/gen/ast"
)

type GoBuilder struct {
	defaults []string
	file     *ast.File
}

func (jenny *GoBuilder) JennyName() string {
	return "GoRawTypes"
}

func (jenny *GoBuilder) Generate(file *ast.File) (*codejen.File, error) {
	jenny.file = file
	jenny.defaults = nil

	output, err := jenny.generateFile(file)
	if err != nil {
		return nil, err
	}

	return codejen.NewFile(file.Package+"_builder_gen.go", output, jenny), nil
}

func (jenny *GoBuilder) generateFile(file *ast.File) ([]byte, error) {
	var buffer strings.Builder
	tr := newPreprocessor()
	entryPointType, ok := file.EntryPointType()
	if !ok {
		return nil, fmt.Errorf("coult not find entrypoint type")
	}

	tr.translateDefinitions(file.Types)

	buffer.WriteString(fmt.Sprintf("package %s\n\n", file.Package))

	// import generated types
	buffer.WriteString(fmt.Sprintf("import \"github.com/grafana/grok/newgen/%s/types\"\n\n", file.Package))

	// Option type declaration
	buffer.WriteString("type Option func(builder *Builder) error\n\n")

	// Builder type declaration
	buffer.WriteString(fmt.Sprintf(`type Builder struct {
	internal *types.%s
}
`, entryPointType.Name))

	// Include veneers if any
	templateFile := fmt.Sprintf("%s.builder.go.tmpl", strings.ToLower(entryPointType.Name))
	tmpl := templates.Lookup(templateFile)
	if tmpl != nil {
		buf := bytes.Buffer{}
		if err := tmpl.Execute(&buf, nil); err != nil {
			return nil, fmt.Errorf("failed executing veneer template: %w", err)
		}

		buffer.WriteString(buf.String())
	}

	// Define options from types
	for _, typeDef := range tr.sortedTypes() {
		typeDefGen, err := jenny.formatTypeDef(typeDef)
		if err != nil {
			return nil, err
		}
		if typeDefGen == nil {
			continue
		}

		buffer.Write(typeDefGen)
		buffer.WriteString("\n")
	}

	// add calls to set default values
	buffer.WriteString("func defaults() []Option {\n")
	buffer.WriteString("return []Option{\n")
	for _, defaultCall := range jenny.defaults {
		buffer.WriteString(defaultCall + ",\n")
	}
	buffer.WriteString("}\n")
	buffer.WriteString("}\n")

	return []byte(buffer.String()), nil
}

func (jenny *GoBuilder) formatTypeDef(def ast.Definition) ([]byte, error) {
	// nothing to do for enums & other non-struct types
	if def.Kind != ast.KindStruct {
		return nil, nil
	}

	// No options if not main/entrypoint type
	if !def.IsEntryPoint {
		return nil, nil
	}

	return jenny.formatMainTypeOptions(def)
}

func (jenny *GoBuilder) formatMainTypeOptions(def ast.Definition) ([]byte, error) {
	var buffer strings.Builder

	for _, fieldDef := range def.Fields {
		buffer.WriteString(jenny.fieldToOption(fieldDef))
	}

	return []byte(buffer.String()), nil
}

func (jenny *GoBuilder) fieldToOption(def ast.FieldDefinition) string {
	var buffer strings.Builder

	fieldName := strings.Title(def.Name)
	typeName := strings.TrimPrefix(formatType(def.Type, def.Required, "types"), "*")

	generatedConstraints := strings.Join(jenny.constraints(def.Name, def.Type.Constraints), "\n")
	asPointer := ""
	// FIXME: this condition is probably wrong
	if def.Type.Nullable || (def.Type.Kind != ast.KindArray && def.Type.Kind != ast.KindStruct && !def.Required) {
		asPointer = "&"
	}

	defaultValue := def.Type.Default
	if def.Type.IsReference() {
		referredType := jenny.file.LocateDefinition(string(def.Type.Kind))
		defaultValue = referredType.Default
	}
	if defaultValue != nil {
		jenny.defaults = append(jenny.defaults, fmt.Sprintf("%[1]s(%#[2]v)", fieldName, defaultValue))
	}

	buffer.WriteString(fmt.Sprintf(`
func %[1]s(%[2]s %[3]s) Option {
	return func(builder *Builder) error {
		%[4]s
		builder.internal.%[1]s = %[5]s%[2]s

		return nil
	}
}
`, fieldName, def.Name, typeName, generatedConstraints, asPointer))

	return buffer.String()
}

func (jenny *GoBuilder) constraints(argumentName string, constraints []ast.TypeConstraint) []string {
	output := make([]string, 0, len(constraints))

	for _, constraint := range constraints {
		output = append(output, jenny.constraint(argumentName, constraint))
	}

	return output
}

func (jenny *GoBuilder) constraint(argumentName string, constraint ast.TypeConstraint) string {
	var buffer strings.Builder

	buffer.WriteString(fmt.Sprintf("if !(%s) {\n", jenny.constraintComparison(argumentName, constraint)))
	buffer.WriteString(fmt.Sprintf("return errors.New(\"%[1]s must be %[2]s %[3]v\")\n", argumentName, constraint.Op, constraint.Args[0]))
	buffer.WriteString("}\n")

	return buffer.String()
}

func (jenny *GoBuilder) constraintComparison(argumentName string, constraint ast.TypeConstraint) string {
	if constraint.Op == "minLength" {
		return fmt.Sprintf("len([]rune(%[1]s)) >= %[2]v", argumentName, constraint.Args[0])
	}
	if constraint.Op == "maxLength" {
		return fmt.Sprintf("len([]rune(%[1]s)) <= %[2]v", argumentName, constraint.Args[0])
	}

	return fmt.Sprintf("%[1]s %[2]s %[3]v", argumentName, constraint.Op, constraint.Args[0])
}
