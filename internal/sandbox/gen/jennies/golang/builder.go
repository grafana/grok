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
}

func (jenny *GoBuilder) JennyName() string {
	return "GoBuilder"
}

func (jenny *GoBuilder) Generate(builder ast.Builder) (codejen.Files, error) {
	output, err := jenny.generateBuilder(builder)
	if err != nil {
		return nil, err
	}

	return codejen.Files{
		*codejen.NewFile(builder.Package+"/"+strings.ToLower(builder.For.Name)+"/builder_gen.go", output, jenny),
	}, nil
}

func (jenny *GoBuilder) generateBuilder(builder ast.Builder) ([]byte, error) {
	var buffer strings.Builder

	buffer.WriteString(fmt.Sprintf("package %s\n\n", strings.ToLower(builder.For.Name)))

	// import generated types
	buffer.WriteString(fmt.Sprintf("import \"github.com/grafana/grok/generated/types/%s\"\n\n", builder.Package))

	// Option type declaration
	buffer.WriteString("type Option func(builder *Builder) error\n\n")

	// Builder type declaration
	buffer.WriteString(fmt.Sprintf(`type Builder struct {
	internal *types.%s
}

`, tools.UpperCamelCase(builder.For.Name)))

	// Add a constructor for the builder
	constructorCode := jenny.generateConstructor(builder)
	buffer.WriteString(constructorCode)

	// Add JSON (un)marshaling shortcuts
	jsonMarshal, err := jenny.veneer("json_marshal", builder.For)
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
`, tools.UpperCamelCase(builder.For.Name)))

	// Define options
	for _, option := range builder.Options {
		buffer.WriteString(jenny.generateOption(option) + "\n")
	}

	// add calls to set default values
	buffer.WriteString("\n")
	buffer.WriteString("func defaults() []Option {\n")
	buffer.WriteString("return []Option{\n")
	for _, opt := range builder.Options {
		if opt.Default != nil {
			buffer.WriteString(jenny.generateDefaultCall(opt) + ",\n")
		}
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

func (jenny *GoBuilder) generateConstructor(builder ast.Builder) string {
	var buffer strings.Builder

	typeName := tools.UpperCamelCase(builder.For.Name)
	args := ""
	fieldsInit := ""
	var argsList []string
	var fieldsInitList []string
	for _, opt := range builder.Options {
		if !opt.IsConstructorArg {
			continue
		}

		// FIXME: this is assuming that there's only one argument for that option
		argsList = append(argsList, jenny.generateArgument(opt.Args[0]))
		fieldsInitList = append(
			fieldsInitList,
			jenny.generateInitAssignment(opt.Assignments[0]),
		)
	}

	if len(argsList) != 0 {
		args = strings.Join(argsList, ", ") + ", "
	}
	if len(fieldsInitList) != 0 {
		fieldsInit = strings.Join(fieldsInitList, ",\n") + ",\n"
	}

	buffer.WriteString(fmt.Sprintf(`
func New(%[2]soptions ...Option) (Builder, error) {
	resource := &types.%[1]s{
		%[3]s
	}
	builder := &Builder{internal: resource}

	for _, opt := range append(defaults(), options...) {
		if err := opt(builder); err != nil {
			return *builder, err
		}
	}

	return *builder, nil
}
`, typeName, args, fieldsInit))

	return buffer.String()
}

func (jenny *GoBuilder) formatFieldPath(fieldPath string) string {
	parts := strings.Split(fieldPath, ".")
	formatted := make([]string, 0, len(parts))

	for _, part := range parts {
		formatted = append(formatted, tools.UpperCamelCase(part))
	}

	return strings.Join(formatted, ".")
}

func (jenny *GoBuilder) generateInitAssignment(assignment ast.Assignment) string {
	fieldPath := jenny.formatFieldPath(assignment.Path)
	valueType := assignment.ValueType

	if assignment.ValueHasBuilder {
		return "constructor init assignment with type that has a builder is not supported yet"
	}

	if assignment.ArgumentName == "" {
		return fmt.Sprintf("%[1]s: %[2]s", fieldPath, formatScalar(assignment.Value))
	}

	argName := jenny.escapeVarName(tools.LowerCamelCase(assignment.ArgumentName))

	asPointer := ""
	// FIXME: this condition is probably wrong
	if valueType.Kind() != ast.KindArray && valueType.Kind() != ast.KindStruct && assignment.IntoOptionalField {
		asPointer = "&"
	}

	generatedConstraints := strings.Join(jenny.constraints(argName, assignment.Constraints), "\n")
	if generatedConstraints != "" {
		generatedConstraints = generatedConstraints + "\n\n"
	}

	return generatedConstraints + fmt.Sprintf("%[1]s: %[3]s%[2]s", fieldPath, argName, asPointer)
}

func (jenny *GoBuilder) generateOption(def ast.Option) string {
	var buffer strings.Builder

	for _, commentLine := range def.Comments {
		buffer.WriteString(fmt.Sprintf("// %s\n", commentLine))
	}

	// Option name
	optionName := tools.UpperCamelCase(def.Name)

	// Arguments list
	arguments := ""
	if len(def.Args) != 0 {
		argsList := make([]string, 0, len(def.Args))
		for _, arg := range def.Args {
			argsList = append(argsList, jenny.generateArgument(arg))
		}

		arguments = strings.Join(argsList, ", ")
	}

	// Assignments
	assignmentsList := make([]string, 0, len(def.Assignments))
	for _, assignment := range def.Assignments {
		assignmentsList = append(assignmentsList, jenny.generateAssignment(assignment))
	}
	assignments := strings.Join(assignmentsList, "\n")

	buffer.WriteString(fmt.Sprintf(`func %[1]s(%[2]s) Option {
	return func(builder *Builder) error {
		%[3]s

		return nil
	}
}
`, optionName, arguments, assignments))

	return buffer.String()
}

func (jenny *GoBuilder) generateArgument(arg ast.Argument) string {
	typeName := formatType(arg.Type, true, "types")

	if arg.TypeHasBuilder {
		referredTypeName := arg.Type.(ast.RefType).ReferredType
		referredTypePkg := strings.ToLower(referredTypeName)

		return fmt.Sprintf(`opts ...%[1]s.Option`, referredTypePkg)
	}

	name := jenny.escapeVarName(tools.LowerCamelCase(arg.Name))

	return fmt.Sprintf("%s %s", name, typeName)
}

func (jenny *GoBuilder) generateAssignment(assignment ast.Assignment) string {
	fieldPath := jenny.formatFieldPath(assignment.Path)
	valueType := assignment.ValueType

	if assignment.ValueHasBuilder {
		referredType := valueType.(ast.RefType)
		referredTypePkg := strings.ToLower(referredType.ReferredType)

		return fmt.Sprintf(`resource, err := %[2]s.New(opts...)
		if err != nil {
			return err
		}

		builder.internal.%[1]s = resource.Internal()
`, fieldPath, referredTypePkg)
	}

	if assignment.ArgumentName == "" {
		return fmt.Sprintf("builder.internal.%[1]s = %[2]s", fieldPath, formatScalar(assignment.Value))
	}

	argName := jenny.escapeVarName(tools.LowerCamelCase(assignment.ArgumentName))

	asPointer := ""
	// FIXME: this condition is probably wrong
	if valueType.Kind() != ast.KindArray && valueType.Kind() != ast.KindStruct && assignment.IntoOptionalField {
		asPointer = "&"
	}

	generatedConstraints := strings.Join(jenny.constraints(argName, assignment.Constraints), "\n")
	if generatedConstraints != "" {
		generatedConstraints = generatedConstraints + "\n\n"
	}

	return generatedConstraints + fmt.Sprintf("builder.internal.%[1]s = %[3]s%[2]s", fieldPath, argName, asPointer)
}

func (jenny *GoBuilder) escapeVarName(varName string) string {
	if isReservedGoKeyword(varName) {
		return varName + "Arg"
	}

	return varName
}

func (jenny *GoBuilder) generateDefaultCall(option ast.Option) string {
	args := make([]string, 0, len(option.Default.ArgsValues))
	for _, arg := range option.Default.ArgsValues {
		args = append(args, formatScalar(arg))
	}

	return fmt.Sprintf("%s(%s)", tools.UpperCamelCase(option.Name), strings.Join(args, ", "))
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
