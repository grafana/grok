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

func (jenny TypescriptRawTypes) Generate(file *ast.File) (codejen.Files, error) {
	output, err := jenny.generateFile(file)
	if err != nil {
		return nil, err
	}

	return codejen.Files{
		*codejen.NewFile(file.Package+"_types_gen.ts", output, jenny),
	}, nil
}

func (jenny TypescriptRawTypes) generateFile(file *ast.File) ([]byte, error) {
	var buffer strings.Builder

	for _, typeDef := range file.Definitions {
		typeDefGen, err := jenny.formatObject(typeDef, "")
		if err != nil {
			return nil, err
		}

		buffer.Write(typeDefGen)
		buffer.WriteString("\n")
	}

	return []byte(buffer.String()), nil
}

func (jenny TypescriptRawTypes) formatObject(def ast.Object, typesPkg string) ([]byte, error) {
	switch def.Type.Kind() {
	case ast.KindStruct:
		return jenny.formatStructDef(def, typesPkg)
	case ast.KindEnum:
		return jenny.formatEnumDef(def)
	case ast.KindDisjunction:
		disj := formatDisjunction(def.Type.(ast.DisjunctionType), typesPkg)

		return []byte(fmt.Sprintf("type %s = %s;\n", def.Name, disj)), nil
	case ast.KindAny:
		return []byte(fmt.Sprintf("type %s = any;\n", def.Name)), nil
	default:
		return nil, fmt.Errorf("unhandled type def kind: %s", def.Type.Kind())
	}
}

func (jenny TypescriptRawTypes) formatEnumDef(def ast.Object) ([]byte, error) {
	var buffer strings.Builder

	for _, commentLine := range def.Comments {
		buffer.WriteString(fmt.Sprintf("// %s\n", commentLine))
	}

	enumType := def.Type.(ast.EnumType)

	buffer.WriteString(fmt.Sprintf("export enum %s {\n", def.Name))
	for _, val := range enumType.Values {
		buffer.WriteString(fmt.Sprintf("\t%s = %#v,\n", strings.Title(val.Name), val.Value))
	}
	buffer.WriteString("}\n")

	return []byte(buffer.String()), nil
}

func (jenny TypescriptRawTypes) formatStructDef(def ast.Object, typesPkg string) ([]byte, error) {
	var buffer strings.Builder

	for _, commentLine := range def.Comments {
		buffer.WriteString(fmt.Sprintf("// %s\n", commentLine))
	}

	buffer.WriteString(fmt.Sprintf("export interface %s ", tools.UpperCamelCase(def.Name)))

	structType := def.Type.(ast.StructType)

	body := formatStructFields(structType.Fields, typesPkg)

	buffer.WriteString(body + "\n")

	return []byte(buffer.String()), nil
}

func formatStructFields(fields []ast.StructField, typesPkg string) string {
	var buffer strings.Builder

	buffer.WriteString("{\n")

	for i, fieldDef := range fields {
		fieldDefGen := formatField(fieldDef, typesPkg)

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

	return buffer.String()
}

func formatField(def ast.StructField, typesPkg string) []byte {
	var buffer strings.Builder

	for _, commentLine := range def.Comments {
		buffer.WriteString(fmt.Sprintf("// %s\n", commentLine))
	}

	required := ""
	if !def.Required {
		required = "?"
	}

	formattedType := formatType(def.Type, typesPkg)

	buffer.WriteString(fmt.Sprintf(
		"%s%s: %s;\n",
		tools.LowerCamelCase(def.Name),
		required,
		formattedType,
	))

	return []byte(buffer.String())
}

func formatType(def ast.Type, typesPkg string) string {
	// todo: handle nullable
	// maybe if nullable, append | null to the type?
	switch def.Kind() {
	case ast.KindDisjunction:
		return formatDisjunction(def.(ast.DisjunctionType), typesPkg)
	case ast.KindRef:
		if typesPkg != "" {
			return typesPkg + "." + (def.(ast.RefType)).ReferredType
		}

		return (def.(ast.RefType)).ReferredType
	case ast.KindArray:
		return formatArray(def.(ast.ArrayType), typesPkg)
	case ast.KindStruct:
		return formatStructFields(def.(ast.StructType).Fields, typesPkg)
	case ast.KindMap:
		return formatMap(def.(ast.MapType), typesPkg)
	case ast.KindEnum:
		return formatAnonymousEnum(def.(ast.EnumType))

	case ast.KindLiteral:
		return formatLiteral(def.(ast.Literal))

	case ast.KindNull:
		return "null"
	case ast.KindAny:
		return "any"

	case ast.KindBytes, ast.KindString:
		return "string"

	case ast.KindFloat32, ast.KindFloat64:
		return "number"
	case ast.KindUint8, ast.KindUint16, ast.KindUint32, ast.KindUint64:
		return "number"
	case ast.KindInt8, ast.KintInt16, ast.KindInt32, ast.KindInt64:
		return "number"

	case ast.KindBool:
		return "boolean"

	default:
		return string(def.Kind())
	}
}

func formatArray(def ast.ArrayType, typesPkg string) string {
	subTypeString := formatType(def.ValueType, typesPkg)

	return fmt.Sprintf("%s[]", subTypeString)
}

func formatDisjunction(def ast.DisjunctionType, typesPkg string) string {
	subTypes := make([]string, 0, len(def.Branches))
	for _, subType := range def.Branches {
		subTypes = append(subTypes, formatType(subType, typesPkg))
	}

	return strings.Join(subTypes, " | ")
}

func formatMap(def ast.MapType, typesPkg string) string {
	keyTypeString := formatType(def.IndexType, typesPkg)
	valueTypeString := formatType(def.ValueType, typesPkg)

	return fmt.Sprintf("Record<%s, %s>", keyTypeString, valueTypeString)
}

func formatAnonymousEnum(def ast.EnumType) string {
	values := make([]string, 0, len(def.Values))
	for _, value := range def.Values {
		values = append(values, fmt.Sprintf("%#v", value.Value))
	}

	enumeration := strings.Join(values, " | ")

	return enumeration
}

func formatLiteral(def ast.Literal) string {
	return fmt.Sprintf("%#v", def.Value)
}

func prefixLinesWith(input string, prefix string) string {
	lines := strings.Split(input, "\n")
	prefixed := make([]string, 0, len(lines))

	for _, line := range lines {
		prefixed = append(prefixed, prefix+line)
	}

	return strings.Join(prefixed, "\n")
}
