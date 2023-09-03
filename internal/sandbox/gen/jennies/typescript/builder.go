package typescript

import (
	"fmt"
	"strings"

	"github.com/grafana/codejen"
	"github.com/grafana/grok/internal/sandbox/gen/ast"
	"github.com/grafana/grok/internal/sandbox/gen/ast/compiler"
	"github.com/grafana/grok/internal/sandbox/gen/jennies/tools"
)

type TypescriptBuilder struct {
	defaults []string
	file     *ast.File
}

func (jenny *TypescriptBuilder) JennyName() string {
	return "TypescriptBuilder"
}

func (jenny *TypescriptBuilder) Generate(file *ast.File) (codejen.Files, error) {
	preprocessedFile, err := compiler.RewriteEngine().Process([]*ast.File{file})
	if err != nil {
		return nil, err
	}

	jenny.file = preprocessedFile[0]

	var files []codejen.File
	for _, definition := range jenny.file.Definitions {
		// No need for a builder if the object isn't a struct
		if definition.Type.Kind() != ast.KindStruct {
			continue
		}

		output, err := jenny.generateDefinition(definition)
		if err != nil {
			return nil, err
		}

		files = append(files, *codejen.NewFile(strings.ToLower(definition.Name)+"/builder_gen.ts", output, jenny))
	}

	return files, nil
}

func (jenny *TypescriptBuilder) generateDefinition(def ast.Object) ([]byte, error) {
	var buffer strings.Builder
	jenny.defaults = nil
	structType := def.Type.(ast.StructType)
	objectName := tools.UpperCamelCase(def.Name)

	// imports
	buffer.WriteString(fmt.Sprintf("import * as types from \"../%s_types_gen\";\n", strings.ToLower(objectName)))
	buffer.WriteString(fmt.Sprintf("import { OptionsBuilder } from \"../options_builder_gen\";\n\n"))

	// Builder class declaration
	buffer.WriteString(fmt.Sprintf("export class %[1]sBuilder implements OptionsBuilder<types.%[1]s> {\n", objectName))

	// internal property, representing the object being built
	buffer.WriteString(fmt.Sprintf("\tinternal: types.%[1]s;\n", objectName))

	// Allow builders to expose the resource they're building
	buffer.WriteString(fmt.Sprintf(`
	build(): types.%s {
		return this.internal;
	}

`, objectName))

	// Define options from fields
	for _, fieldDef := range structType.Fields {
		opt, err := jenny.fieldToOption(fieldDef)
		if err != nil {
			return nil, err
		}
		buffer.WriteString(opt)
	}

	// End builder class declaration
	buffer.WriteString("}\n")

	return []byte(buffer.String()), nil
}

func (jenny *TypescriptBuilder) fieldToOption(def ast.StructField) (string, error) {
	var buffer strings.Builder

	for _, commentLine := range def.Comments {
		buffer.WriteString(fmt.Sprintf("\t// %s\n", commentLine))
	}

	// references to objects get their own builder
	if def.Type.Kind() == ast.KindRef {
		referredDef := jenny.file.LocateDefinition(def.Type.(ast.RefType).ReferredType)
		if referredDef.Type.Kind() == ast.KindStruct {
			return jenny.referenceFieldToOption(def), nil
		}
	}

	// literal options get their own simplified builder
	if def.Type.Kind() == ast.KindLiteral {
		return jenny.literalFieldToOption(def), nil
	}

	optionName := tools.UpperCamelCase(def.DisplayName)
	argumentName := tools.LowerCamelCase(def.DisplayName)
	typeName, err := formatType(def.Type, "types")
	if err != nil {
		return "", err
	}

	generatedConstraints := ""
	if scalarType, ok := def.Type.(ast.ScalarType); ok {
		generatedConstraints = strings.Join(jenny.constraints(argumentName, scalarType.Constraints), "\n")
	}

	buffer.WriteString(fmt.Sprintf(`	with%[1]s(%[3]s: %[4]s): this {
		%[5]s
		this.internal.%[2]s = %[3]s;

		return this;
	}

`, optionName, def.Name, argumentName, typeName, generatedConstraints))

	return buffer.String(), nil
}

func (jenny *TypescriptBuilder) literalFieldToOption(def ast.StructField) string {
	var buffer strings.Builder

	optionName := tools.UpperCamelCase(def.DisplayName)

	literalDef := def.Type.(ast.Literal)
	value := jenny.formatScalar(literalDef.Value)

	buffer.WriteString(fmt.Sprintf(`	with%[1]s(): this {
		this.internal.%[2]s = %[3]s;

		return this;
	}

`, optionName, def.Name, value))

	return buffer.String()
}

func (jenny *TypescriptBuilder) formatScalar(val any) string {
	if list, ok := val.([]any); ok {
		items := make([]string, 0, len(list))

		for _, item := range list {
			items = append(items, jenny.formatScalar(item))
		}

		// TODO: we can't assume a list of strings
		return fmt.Sprintf("[%s]", strings.Join(items, ", "))
	}

	return fmt.Sprintf("%#v", val)
}

func (jenny *TypescriptBuilder) referenceFieldToOption(def ast.StructField) string {
	var buffer strings.Builder

	referredType := tools.UpperCamelCase(def.Type.(ast.RefType).ReferredType)
	optionName := tools.UpperCamelCase(def.DisplayName)

	buffer.WriteString(fmt.Sprintf(`	with%[1]s(builder: OptionsBuilder<types.%[2]s>): this {
		this.internal.%[3]s = builder.build();

		return this;
	}

`, optionName, referredType, def.Name))

	return buffer.String()
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
