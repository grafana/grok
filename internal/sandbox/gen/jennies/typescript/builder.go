package typescript

import (
	"fmt"
	"strings"

	"github.com/grafana/codejen"
	"github.com/grafana/grok/internal/sandbox/gen/ast"
	"github.com/grafana/grok/internal/sandbox/gen/jennies/tools"
)

type TypescriptBuilder struct {
	defaults []string
	file     *ast.File
}

func (jenny *TypescriptBuilder) JennyName() string {
	return "TypescriptBuilder"
}

func (jenny *TypescriptBuilder) Generate(builder ast.Builder) (codejen.Files, error) {
	output, err := jenny.generateBuilder(builder)
	if err != nil {
		return nil, err
	}

	return codejen.Files{
		*codejen.NewFile(strings.ToLower(builder.For.Name)+"/builder_gen.ts", output, jenny),
	}, nil
}

func (jenny *TypescriptBuilder) generateBuilder(builder ast.Builder) ([]byte, error) {
	var buffer strings.Builder

	objectName := tools.UpperCamelCase(builder.For.Name)

	// imports
	buffer.WriteString(fmt.Sprintf("import * as types from \"../%s_types_gen\";\n", strings.ToLower(objectName)))
	buffer.WriteString(fmt.Sprintf("import { OptionsBuilder } from \"../options_builder_gen\";\n\n"))

	// Builder class declaration
	buffer.WriteString(fmt.Sprintf("export class %[1]sBuilder implements OptionsBuilder<types.%[1]s> {\n", objectName))

	// internal property, representing the object being built
	buffer.WriteString(fmt.Sprintf("\tinternal: types.%[1]s;\n", objectName))

	// Add a constructor for the builder
	constructorCode := jenny.generateConstructor(builder)
	buffer.WriteString(constructorCode)

	// Allow builders to expose the resource they're building
	buffer.WriteString(fmt.Sprintf(`
	build(): types.%s {
		return this.internal;
	}

`, objectName))

	// Define options
	for _, option := range builder.Options {
		opt, err := jenny.generateOption(option)
		if err != nil {
			return nil, err
		}
		buffer.WriteString(opt)
	}

	// End builder class declaration
	buffer.WriteString("}\n")

	return []byte(buffer.String()), nil
}

func (jenny *TypescriptBuilder) generateConstructor(builder ast.Builder) string {
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

	args = strings.Join(argsList, ", ")
	fieldsInit = strings.Join(fieldsInitList, "\n")

	buffer.WriteString(fmt.Sprintf(`
	constructor(%[2]s) {
		%[3]s
	}
`, typeName, args, fieldsInit))

	return buffer.String()
}

func (jenny *TypescriptBuilder) generateInitAssignment(assignment ast.Assignment) string {
	fieldPath := assignment.Path

	if assignment.ValueHasBuilder {
		return "constructor init assignment with type that has a builder is not supported yet"
	}

	if assignment.ArgumentName == "" {
		return fmt.Sprintf("this.internal.%[1]s = %[2]s;", fieldPath, formatScalar(assignment.Value))
	}

	argName := tools.LowerCamelCase(assignment.ArgumentName)

	generatedConstraints := strings.Join(jenny.constraints(argName, assignment.Constraints), "\n")
	if generatedConstraints != "" {
		generatedConstraints = generatedConstraints + "\n\n"
	}

	return generatedConstraints + fmt.Sprintf("this.internal.%[1]s = %[2]s;", fieldPath, argName)
}

func (jenny *TypescriptBuilder) generateOption(def ast.Option) (string, error) {
	var buffer strings.Builder

	for _, commentLine := range def.Comments {
		buffer.WriteString(fmt.Sprintf("\t// %s\n", commentLine))
	}

	// Option name
	optionName := tools.UpperCamelCase(def.Title)

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

	// Option body
	buffer.WriteString(fmt.Sprintf(`	with%[1]s(%[2]s): this {
		%[3]s

		return this;
	}

`, optionName, arguments, assignments))

	return buffer.String(), nil
}

func (jenny *TypescriptBuilder) generateArgument(arg ast.Argument) string {
	typeName := formatType(arg.Type, "types")

	if arg.TypeHasBuilder {
		referredTypeName := arg.Type.(ast.RefType).ReferredType
		referredTypePkg := tools.UpperCamelCase(referredTypeName)

		return fmt.Sprintf(`%[1]s: OptionsBuilder<types.%[2]s>`, arg.Name, referredTypePkg)
	}

	name := tools.LowerCamelCase(arg.Name)

	return fmt.Sprintf("%s: %s", name, typeName)
}

func (jenny *TypescriptBuilder) generateAssignment(assignment ast.Assignment) string {
	fieldPath := assignment.Path

	if assignment.ValueHasBuilder {
		return fmt.Sprintf(`this.internal.%[1]s = %[2]s.build();
`, fieldPath, assignment.ArgumentName)
	}

	if assignment.ArgumentName == "" {
		return fmt.Sprintf("this.internal.%[1]s = %[2]s;", fieldPath, formatScalar(assignment.Value))
	}

	argName := tools.LowerCamelCase(assignment.ArgumentName)

	generatedConstraints := strings.Join(jenny.constraints(argName, assignment.Constraints), "\n")
	if generatedConstraints != "" {
		generatedConstraints = generatedConstraints + "\n\n"
	}

	return generatedConstraints + fmt.Sprintf("this.internal.%[1]s = %[2]s;", fieldPath, argName)

}

func (jenny *TypescriptBuilder) constraints(argumentName string, constraints []ast.TypeConstraint) []string {
	output := make([]string, 0, len(constraints))

	for _, constraint := range constraints {
		output = append(output, jenny.constraint(argumentName, constraint))
	}

	return output
}

func (jenny *TypescriptBuilder) constraint(argumentName string, constraint ast.TypeConstraint) string {
	var buffer strings.Builder

	buffer.WriteString(fmt.Sprintf("if (!(%s)) {\n", jenny.constraintComparison(argumentName, constraint)))
	buffer.WriteString(fmt.Sprintf("\t\t\tthrow new Error(\"%[1]s must be %[2]s %[3]v\");\n", argumentName, constraint.Op, constraint.Args[0]))
	buffer.WriteString("\t\t}\n")

	return buffer.String()
}

func (jenny *TypescriptBuilder) constraintComparison(argumentName string, constraint ast.TypeConstraint) string {
	if constraint.Op == "minLength" {
		return fmt.Sprintf("%[1]s.length >= %[2]v", argumentName, constraint.Args[0])
	}
	if constraint.Op == "maxLength" {
		return fmt.Sprintf("%[1]s.length <= %[2]v", argumentName, constraint.Args[0])
	}

	return fmt.Sprintf("%[1]s %[2]s %#[3]v", argumentName, constraint.Op, constraint.Args[0])
}
