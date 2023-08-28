package golang

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/grafana/codejen"
	"github.com/grafana/grok/internal/sandbox/gen/ast"
	"github.com/grafana/grok/internal/sandbox/gen/jennies/tools"
)

type GoBuilder struct {
	defaults []string
	file     *ast.File
}

func (jenny *GoBuilder) JennyName() string {
	return "GoBuilder"
}

func (jenny *GoBuilder) Generate(file *ast.File) (codejen.Files, error) {
	jenny.file = file

	tr := newPreprocessor()
	tr.translateDefinitions(file.Definitions)

	var files []codejen.File
	for _, definition := range tr.sortedDefinitions() {
		if definition.Type.Kind() != ast.KindStruct {
			continue
		}

		output, err := jenny.generateDefinition(definition)
		if err != nil {
			return nil, err
		}

		files = append(files, *codejen.NewFile(strings.ToLower(definition.Name)+"/builder_gen.go", output, jenny))
	}

	return files, nil
}

func (jenny *GoBuilder) generateDefinition(def ast.Object) ([]byte, error) {
	var buffer strings.Builder
	jenny.defaults = nil
	structType := def.Type.(*ast.StructType)

	buffer.WriteString(fmt.Sprintf("package %s\n\n", strings.ToLower(def.Name)))

	// import generated types
	buffer.WriteString(fmt.Sprintf("import \"github.com/grafana/grok/newgen/%s/types\"\n\n", jenny.file.Package))

	// Option type declaration
	buffer.WriteString("type Option func(builder *Builder) error\n\n")

	// Builder type declaration
	buffer.WriteString(fmt.Sprintf(`type Builder struct {
	internal *types.%s
}
`, tools.UpperCamelCase(def.Name)))

	// Add a constructor for the builder
	constructorCode, err := jenny.veneer("constructor", def)
	if err != nil {
		return nil, err
	}
	buffer.WriteString(constructorCode)

	// Add JSON (un)marshaling shortcuts
	jsonMarshal, err := jenny.veneer("json_marshal", def)
	if err != nil {
		return nil, err
	}
	buffer.WriteString(jsonMarshal)

	// Allow builders to expose the resource they're building
	// TODO: do we want to do this?
	// TODO: better name, with less conflict chance
	buffer.WriteString(fmt.Sprintf(`
func (builder *Builder) Internal() *types.%s {
	return builder.internal
}
`, tools.UpperCamelCase(def.Name)))

	// Define options from fields
	for _, fieldDef := range structType.Fields {
		buffer.WriteString(jenny.fieldToOption(fieldDef))
	}

	// add calls to set default values
	buffer.WriteString("\n")
	buffer.WriteString("func defaults() []Option {\n")
	buffer.WriteString("return []Option{\n")
	for _, defaultCall := range jenny.defaults {
		buffer.WriteString(defaultCall + ",\n")
	}
	buffer.WriteString("}\n")
	buffer.WriteString("}\n")

	return []byte(buffer.String()), nil
}

func (jenny *GoBuilder) veneer(veneerType string, def ast.Object) (string, error) {
	// First, see if there is a definition-specific veneer
	templateFile := fmt.Sprintf("%s.builder.%s.go.tmpl", strings.ToLower(def.Name), veneerType)
	tmpl := templates.Lookup(templateFile)

	// If not, get the generic one
	if tmpl == nil {
		tmpl = templates.Lookup(fmt.Sprintf("builder.%s.go.tmpl", veneerType))
	}
	// If not, something went wrong.
	if tmpl == nil {
		return "", fmt.Errorf("veneer '%s' not found", veneerType)
	}

	buf := bytes.Buffer{}
	if err := tmpl.Execute(&buf, map[string]any{
		"def": def,
	}); err != nil {
		return "", fmt.Errorf("failed executing veneer template: %w", err)
	}

	return buf.String(), nil
}

func (jenny *GoBuilder) fieldToOption(def ast.StructField) string {
	var buffer strings.Builder

	for _, commentLine := range def.Comments {
		buffer.WriteString(fmt.Sprintf("// %s\n", commentLine))
	}

	// structs get their own builder
	if def.Type.Kind() == ast.KindRef {
		referredDef := jenny.file.LocateDefinition(def.Type.(*ast.RefType).ReferredType)
		if referredDef.Type.Kind() == ast.KindStruct {
			return jenny.referenceFieldToOption(def)
		}
	}

	fieldName := tools.UpperCamelCase(def.Name)
	typeName := strings.TrimPrefix(formatType(def.Type, def.Required, "types"), "*")
	argumentName := tools.LowerCamelCase(def.Name)
	if isReservedGoKeyword(argumentName) {
		argumentName = argumentName + "Arg"
	}

	generatedConstraints := ""
	asPointer := ""
	/*
		generatedConstraints := strings.Join(jenny.constraints(argumentName, def.Type.Constraints), "\n")
		asPointer := ""
		// FIXME: this condition is probably wrong
		if def.Type.Nullable || (def.Type.Kind != ast.KindArray && def.Type.Kind != ast.KindStruct && !def.Required) {
			asPointer = "&"
		}

		if def.HasDefaultValue() {
			jenny.defaults = append(jenny.defaults, jenny.formatDefaultValue(def))
		}
	*/

	buffer.WriteString(fmt.Sprintf(`func %[1]s(%[2]s %[3]s) Option {
	return func(builder *Builder) error {
		%[4]s
		builder.internal.%[1]s = %[5]s%[2]s

		return nil
	}
}
`, fieldName, argumentName, typeName, generatedConstraints, asPointer))

	return buffer.String()
}

/*
func (jenny *GoBuilder) formatDefaultValue(field ast.FieldDefinition) string {
	fieldName := tools.UpperCamelCase(field.Name)

	if field.Type.Kind != ast.KindStruct {
		defaultValue := field.Type.Default
		if field.Type.IsReference() {
			referredType := jenny.file.LocateDefinition(string(field.Type.Kind))
			defaultValue = referredType.Default
		}

		return fmt.Sprintf("%[1]s(%[2]s)", fieldName, jenny.formatScalar(defaultValue))
	}

	return fmt.Sprintf("%[1]s(%[2]s)", fieldName, jenny.formatAnonymousStructDefaultValue(field.Type))
}

// FIXME: this breaks for anonymous structs with anonymous types defined in one or more of their fields
func (jenny *GoBuilder) formatAnonymousStructDefaultValue(structDef ast.DefinitionImpl) string {
	var buffer bytes.Buffer

	buffer.WriteString(formatStructBody(structDef, "types"))
	buffer.WriteString("{\n")
	for _, field := range structDef.Fields {
		if !field.HasDefaultValue() {
			continue
		}

		fieldName := strings.Title(field.Name)

		buffer.WriteString(fmt.Sprintf("%s: %s,\n", fieldName, jenny.formatScalar(field.Type.Default)))
	}
	buffer.WriteString("\n}")

	return buffer.String()
}
*/

func (jenny *GoBuilder) formatScalar(val any) string {
	if list, ok := val.([]any); ok {
		items := make([]string, 0, len(list))

		for _, item := range list {
			items = append(items, jenny.formatScalar(item))
		}

		// TODO: we can't assume a list of strings
		return fmt.Sprintf("[]string{%s}", strings.Join(items, ", "))
	}

	return fmt.Sprintf("%#v", val)
}

func (jenny *GoBuilder) referenceFieldToOption(def ast.StructField) string {
	var buffer strings.Builder

	fieldName := tools.UpperCamelCase(def.Name)
	referredPackage := strings.ToLower(def.Type.(*ast.RefType).ReferredType)

	buffer.WriteString(fmt.Sprintf(`
func %[1]s(opts ...%[2]s.Option) Option {
	return func(builder *Builder) error {
		resource, err := %[2]s.New(opts...)
		if err != nil {
			return err
		}

		builder.internal.%[1]s = resource.Internal()

		return nil
	}
}
`, fieldName, referredPackage))

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

	return fmt.Sprintf("%[1]s %[2]s %#[3]v", argumentName, constraint.Op, constraint.Args[0])
}

func isReservedGoKeyword(input string) bool {
	// TODO
	if input == "type" {
		return true
	}

	return false
}
