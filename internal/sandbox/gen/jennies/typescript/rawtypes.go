package typescript

import (
	"fmt"
	"strings"

	"github.com/grafana/codejen"
	"github.com/grafana/grok/internal/sandbox/gen/ast"
	"github.com/grafana/grok/internal/sandbox/gen/jennies/tools"
)

type TypescriptRawTypes struct {
}

func (jenny TypescriptRawTypes) JennyName() string {
	return "TypescriptRawTypes"
}

func (jenny TypescriptRawTypes) Generate(file *ast.File) (*codejen.File, error) {
	output, err := jenny.generateFile(file)
	if err != nil {
		return nil, err
	}

	return codejen.NewFile(file.Package+"_types_gen.ts", output, jenny), nil
}

func (jenny TypescriptRawTypes) generateFile(file *ast.File) ([]byte, error) {
	var buffer strings.Builder

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

func (jenny TypescriptRawTypes) formatTypeDef(def ast.Object) ([]byte, error) {
	switch def.Type.Kind() {
	case ast.KindStruct:
		return jenny.formatStructDef(def)
	case ast.KindEnum:
		return jenny.formatEnumDef(def)
	default:
		return nil, fmt.Errorf("unhandled type def kind: %s", def.Type.Kind())
	}
}

func (jenny TypescriptRawTypes) formatEnumDef(def ast.Object) ([]byte, error) {
	var buffer strings.Builder

	for _, commentLine := range def.Comments {
		buffer.WriteString(fmt.Sprintf("// %s\n", commentLine))
	}

	enumType := def.Type.(*ast.EnumType)

	buffer.WriteString(fmt.Sprintf("export enum %s {\n", def.Name))
	for _, val := range enumType.Values {
		buffer.WriteString(fmt.Sprintf("\t%s = %#v,\n", strings.Title(val.Name), val.Value))
	}
	buffer.WriteString("}\n")

	return []byte(buffer.String()), nil
}

func (jenny TypescriptRawTypes) formatStructDef(def ast.Object) ([]byte, error) {
	var buffer strings.Builder

	for _, commentLine := range def.Comments {
		buffer.WriteString(fmt.Sprintf("// %s\n", commentLine))
	}

	buffer.WriteString(fmt.Sprintf("export interface %s ", tools.UpperCamelCase(def.Name)))

	structType := def.Type.(*ast.StructType)

	body, err := jenny.formatStructFields(structType.Fields)
	if err != nil {
		return nil, nil
	}

	buffer.WriteString(body + "\n")

	return []byte(buffer.String()), nil
}

func (jenny TypescriptRawTypes) formatStructFields(fields []ast.StructField) (string, error) {
	var buffer strings.Builder

	buffer.WriteString("{\n")

	for i, fieldDef := range fields {
		fieldDefGen, err := jenny.formatField(fieldDef)
		if err != nil {
			return "", err
		}

		buffer.WriteString(
			strings.TrimSuffix(
				prefixLinesWith(string(fieldDefGen), "\t"),
				"\n\t",
			),
		)

		if i != len(fields)-1 {
			buffer.WriteString("\n")
		}
	}

	buffer.WriteString("\n}")

	return buffer.String(), nil
}

func (jenny TypescriptRawTypes) formatField(def ast.StructField) ([]byte, error) {
	var buffer strings.Builder

	for _, commentLine := range def.Comments {
		buffer.WriteString(fmt.Sprintf("// %s\n", commentLine))
	}

	required := ""
	if !def.Required {
		required = "?"
	}

	formattedType, err := jenny.formatType(def.Type)
	if err != nil {
		return nil, err
	}

	buffer.WriteString(fmt.Sprintf(
		"%s%s: %s;\n",
		tools.LowerCamelCase(def.Name),
		required,
		formattedType,
	))

	return []byte(buffer.String()), nil
}

func (jenny TypescriptRawTypes) formatType(def ast.Type) (string, error) {
	// todo: handle nullable
	// maybe if nullable, append | null to the type?
	switch def.Kind() {
	case ast.KindDisjunction:
		return jenny.formatDisjunction(def.(*ast.DisjunctionType))
	case ast.KindRef:
		return (def.(*ast.RefType)).ReferredType, nil
	case ast.KindArray:
		return jenny.formatArray(def.(*ast.ArrayType))
	case ast.KindStruct:
		return jenny.formatStructFields(def.(*ast.StructType).Fields)
	case ast.KindMap:
		return jenny.formatMap(def.(*ast.MapType))
	case ast.KindEnum:
		return jenny.formatAnonymousEnum(def.(*ast.EnumType))

	case ast.KindNull:
		return "null", nil
	case ast.KindAny:
		return "any", nil

	case ast.KindBytes, ast.KindString:
		return "string", nil

	case ast.KindFloat32, ast.KindFloat64:
		return "number", nil
	case ast.KindUint8, ast.KindUint16, ast.KindUint32, ast.KindUint64:
		return "number", nil
	case ast.KindInt8, ast.KintInt16, ast.KindInt32, ast.KindInt64:
		return "number", nil

	case ast.KindBool:
		return "boolean", nil

	default:
		return "", fmt.Errorf("unhandled type: %s", def.Kind())
	}
}

func (jenny TypescriptRawTypes) formatArray(def *ast.ArrayType) (string, error) {
	// we don't know what to do here (yet)
	subTypeString, err := jenny.formatType(def.ValueType)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s[]", subTypeString), nil
}

func (jenny TypescriptRawTypes) formatDisjunction(def *ast.DisjunctionType) (string, error) {
	subTypes := make([]string, 0, len(def.Branches))
	for _, subType := range def.Branches {
		formatted, err := jenny.formatType(subType)
		if err != nil {
			return "", err
		}

		subTypes = append(subTypes, formatted)
	}

	return strings.Join(subTypes, " | "), nil
}

func (jenny TypescriptRawTypes) formatMap(def *ast.MapType) (string, error) {
	keyTypeString := def.IndexType
	valueTypeString, err := jenny.formatType(def.ValueType)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Record<%s, %s>", keyTypeString, valueTypeString), nil
}

func (jenny TypescriptRawTypes) formatAnonymousEnum(def *ast.EnumType) (string, error) {
	values := make([]string, 0, len(def.Values))
	for _, value := range def.Values {
		values = append(values, fmt.Sprintf("%#v", value.Value))
	}

	enumeration := strings.Join(values, " | ")

	return enumeration, nil
}

func prefixLinesWith(input string, prefix string) string {
	lines := strings.Split(input, "\n")
	prefixed := make([]string, 0, len(lines))

	for _, line := range lines {
		prefixed = append(prefixed, prefix+line)
	}

	return strings.Join(prefixed, "\n")
}
