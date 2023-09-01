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
	preprocessedFile, err := compiler.RewriteEngine().Process(file)
	if err != nil {
		return nil, err
	}

	jenny.file = preprocessedFile

	var files []codejen.File
	for _, definition := range preprocessedFile.Definitions {
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
	structType := def.Type.(*ast.StructType)
	objectName := tools.UpperCamelCase(def.Name)

	// Builder class declaration
	buffer.WriteString(fmt.Sprintf("export class %[1]sBuilder extends OptionsBuilder<%[1]s> {\n", objectName))

	// internal property, representing the object being built
	buffer.WriteString(fmt.Sprintf("\tinternal: %[1]s;\n", def.Name))

	// Allow builders to expose the resource they're building
	buffer.WriteString(fmt.Sprintf(`
	build(): %s {
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

	// literal options get their own simplified builder
	if def.Type.Kind() == ast.KindLiteral {
		return jenny.literalFieldToOption(def), nil
	}

	optionName := tools.UpperCamelCase(def.DisplayName)
	typeName, err := formatType(def.Type)
	if err != nil {
		return "", err
	}
	argumentName := tools.LowerCamelCase(def.DisplayName)
	if isReservedGoKeyword(argumentName) {
		argumentName = argumentName + "Arg"
	}

	generatedConstraints := ""
	if scalarType, ok := def.Type.(*ast.ScalarType); ok {
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

	literalDef := def.Type.(*ast.Literal)
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
		return fmt.Sprintf("[]string{%s}", strings.Join(items, ", "))
	}

	return fmt.Sprintf("%#v", val)
}

func (jenny *TypescriptBuilder) referenceFieldToOption(def ast.StructField) string {
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

func isReservedGoKeyword(input string) bool {
	// TODO
	if input == "type" {
		return true
	}

	return false
}
