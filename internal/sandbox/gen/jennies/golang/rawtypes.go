package golang

import (
	"fmt"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/grafana/codejen"
	"github.com/grafana/grok/internal/sandbox/gen/ast"
	"github.com/grafana/grok/internal/sandbox/gen/jennies/tools"
)

type GoRawTypes struct {
}

func (jenny GoRawTypes) JennyName() string {
	return "GoRawTypes"
}

func (jenny GoRawTypes) Generate(file *ast.File) (*codejen.File, error) {
	output, err := jenny.generateFile(file)
	if err != nil {
		return nil, err
	}

	return codejen.NewFile("types/"+file.Package+"_types_gen.go", output, jenny), nil
}

func (jenny GoRawTypes) generateFile(file *ast.File) ([]byte, error) {
	var buffer strings.Builder

	buffer.WriteString("package types\n\n")

	for _, typeDef := range file.Definitions {
		typeDefGen, err := jenny.formatTypeDef(typeDef)
		if err != nil {
			return nil, err
		}

		buffer.Write(typeDefGen)
		buffer.WriteString("\n")
	}

	return []byte(buffer.String()), nil
}

func (jenny GoRawTypes) formatTypeDef(def ast.Object) ([]byte, error) {
	if def.Type == nil {
		spew.Dump(def)
		panic("lala")
	}
	switch def.Type.Kind() {
	case ast.KindStruct:
		return jenny.formatStructDef(def)
	case ast.KindEnum:
		return jenny.formatEnumDef(def)
	case ast.KindRef:
		return []byte(fmt.Sprintf("type %s %s", tools.UpperCamelCase(def.Name), def.Type.(*ast.RefType).ReferredType)), nil
	default:
		return nil, fmt.Errorf("unhandled type def kind: %s", def.Type.Kind())
	}
}

func (jenny GoRawTypes) formatEnumDef(def ast.Object) ([]byte, error) {
	var buffer strings.Builder

	for _, commentLine := range def.Comments {
		buffer.WriteString(fmt.Sprintf("// %s\n", commentLine))
	}

	enumName := tools.UpperCamelCase(def.Name)
	enumType := def.Type.(*ast.EnumType)

	buffer.WriteString(fmt.Sprintf("type %s %s\n", enumName, enumType.Values[0].Type.Kind()))

	buffer.WriteString("const (\n")
	for _, val := range enumType.Values {
		buffer.WriteString(fmt.Sprintf("\t%s %s = %#v\n", tools.UpperCamelCase(val.Name), enumName, val.Value))
	}
	buffer.WriteString(")\n")

	return []byte(buffer.String()), nil
}

func (jenny GoRawTypes) formatStructDef(def ast.Object) ([]byte, error) {
	var buffer strings.Builder

	for _, commentLine := range def.Comments {
		buffer.WriteString(fmt.Sprintf("// %s\n", commentLine))
	}

	buffer.WriteString(fmt.Sprintf("type %s ", tools.UpperCamelCase(def.Name)))
	buffer.WriteString(formatStructBody(def.Type.(*ast.StructType), ""))
	buffer.WriteString("\n")

	return []byte(buffer.String()), nil
}

func formatStructBody(def *ast.StructType, typesPkg string) string {
	var buffer strings.Builder

	buffer.WriteString("struct {\n")

	for _, fieldDef := range def.Fields {
		buffer.WriteString("\t" + formatField(fieldDef, typesPkg))
	}

	buffer.WriteString("}")

	return buffer.String()
}

func formatField(def ast.StructField, typesPkg string) string {
	var buffer strings.Builder

	for _, commentLine := range def.Comments {
		buffer.WriteString(fmt.Sprintf("// %s\n", commentLine))
	}

	// ToDo: this doesn't follow references to other types like the builder jenny does
	/*
		if def.Type.Default != nil {
			buffer.WriteString(fmt.Sprintf("// Default: %#v\n", def.Type.Default))
		}
	*/

	jsonOmitEmpty := ""
	if !def.Required {
		jsonOmitEmpty = ",omitempty"
	}

	buffer.WriteString(fmt.Sprintf(
		"%s %s `json:\"%s%s\"`\n",
		tools.UpperCamelCase(def.Name),
		formatType(def.Type, def.Required, typesPkg),
		def.Name,
		jsonOmitEmpty,
	))

	return buffer.String()
}
func formatType(def ast.Type, fieldIsRequired bool, typesPkg string) string {
	if def.Kind() == ast.KindAny {
		return "any"
	}

	if def.Kind() == ast.KindDisjunction {
		return formatDisjunction(def.(*ast.DisjunctionType), typesPkg)
	}

	if def.Kind() == ast.KindArray {
		return formatArray(def.(*ast.ArrayType), typesPkg)
	}

	if def.Kind() == ast.KindMap {
		return formatMap(def.(*ast.MapType), typesPkg)
	}

	if def.Kind() == ast.KindRef {
		typeName := def.(*ast.RefType).ReferredType

		if typesPkg != "" {
			typeName = typesPkg + "." + typeName
		}

		if !fieldIsRequired {
			typeName = "*" + typeName
		}

		return typeName
	}

	if def.Kind() == ast.KindEnum {
		return "enum here"
	}

	// anonymous struct
	if def.Kind() == ast.KindStruct {
		return formatStructBody(def.(*ast.StructType), typesPkg)
	}

	// TODO: there should be an ast.KindScalar with a matching type
	typeName := string(def.(*ast.ScalarType).ScalarKind)

	if !fieldIsRequired {
		typeName = "*" + typeName
	}
	/*
		if def.Nullable || !fieldIsRequired {
			typeName = "*" + typeName
		}
	*/

	return typeName
}

func formatArray(def *ast.ArrayType, typesPkg string) string {
	subTypeString := formatType(def.ValueType, true, typesPkg)

	return fmt.Sprintf("[]%s", subTypeString)
}

func formatMap(def *ast.MapType, typesPkg string) string {
	keyTypeString := def.IndexType.Kind()
	valueTypeString := formatType(def.ValueType, true, typesPkg)

	return fmt.Sprintf("map[%s]%s", keyTypeString, valueTypeString)
}

func formatDisjunction(def *ast.DisjunctionType, typesPkg string) string {
	subTypes := make([]string, 0, len(def.Branches))
	for _, subType := range def.Branches {
		subTypes = append(subTypes, formatType(subType, true, typesPkg))
	}

	return fmt.Sprintf("disjunction<%s>", strings.Join(subTypes, " | "))
}
