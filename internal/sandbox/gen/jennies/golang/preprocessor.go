package golang

import (
	"sort"
	"strings"

	"github.com/grafana/grok/internal/sandbox/gen/ast"
)

type preprocessor struct {
	defs map[string]ast.Definition
}

func newPreprocessor() *preprocessor {
	return &preprocessor{
		defs: make(map[string]ast.Definition),
	}
}

// inefficient, but I'm lazy. It's only used during code generation anyway.
func (preprocessor *preprocessor) sortedDefinitions() []ast.Definition {
	typeNames := make([]string, 0, len(preprocessor.defs))
	for typeName := range preprocessor.defs {
		typeNames = append(typeNames, typeName)
	}

	sort.Strings(typeNames)

	sorted := make([]ast.Definition, 0, len(preprocessor.defs))
	for _, k := range typeNames {
		sorted = append(sorted, preprocessor.defs[k])
	}

	return sorted
}

func (preprocessor *preprocessor) translateDefinitions(definitions []ast.Definition) {
	for _, typeDef := range definitions {
		preprocessor.translate(typeDef)
	}
}

func (preprocessor *preprocessor) translate(def ast.Definition) {
	preprocessor.defs[def.Name] = preprocessor.translateDefinition(def)
}

func (preprocessor *preprocessor) translateDefinition(def ast.Definition) ast.Definition {
	if def.Kind == ast.KindDisjunction {
		return preprocessor.expandDisjunction(def)
	}

	if def.Kind == ast.KindArray {
		translated := preprocessor.translateDefinition(*def.ValueType)
		def.ValueType = &translated

		return def
	}

	if def.Kind != ast.KindStruct {
		return def
	}

	newFields := make([]ast.FieldDefinition, 0, len(def.Fields))
	for _, fieldDef := range def.Fields {
		newFields = append(newFields, preprocessor.translateFieldDefinition(fieldDef))
	}

	newDef := def
	newDef.Fields = newFields

	return newDef
}

func (preprocessor *preprocessor) translateFieldDefinition(def ast.FieldDefinition) ast.FieldDefinition {
	newDef := def

	// anonymous enum
	if def.Type.Kind == ast.KindEnum {
		newEnumType := preprocessor.anonymousEnumToExplicitEnum(def)
		preprocessor.defs[newEnumType.Name] = newEnumType
		newDef.Type = ast.Definition{
			Kind: ast.Kind(newEnumType.Name),
		}
	} else {
		newDef.Type = preprocessor.translateDefinition(def.Type)
	}

	return newDef
}

func (preprocessor *preprocessor) anonymousEnumToExplicitEnum(def ast.FieldDefinition) ast.Definition {
	if def.Type.Kind != ast.KindEnum {
		return def.Type
	}

	enumTypeName := formatIdentifier(def.Name) + "Enum"
	enumType := def.Type

	values := make([]ast.EnumValue, 0, len(enumType.Values))
	for _, val := range enumType.Values {
		values = append(values, ast.EnumValue{
			Type:  val.Type,
			Name:  def.Name + formatIdentifier(val.Name),
			Value: val.Value,
		})
	}

	newType := ast.Definition{
		Kind:   enumType.Values[0].Type,
		Name:   enumTypeName,
		Values: values,
	}

	return newType
}

// def is either a disjunction or a list of unknown sub-types
func (preprocessor *preprocessor) expandDisjunction(def ast.Definition) ast.Definition {
	if def.Kind == ast.KindArray {
		translated := preprocessor.translateDefinition(*def.ValueType)
		def.ValueType = &translated

		return def
	}

	// Ex: type | null
	if len(def.Branches) == 2 && def.Branches.HasNullType() {
		finalType := def.Branches.NonNullTypes()[0]
		finalType.Nullable = true

		return finalType
	}

	// type | otherType | something (| null)?
	// generate a type with a nullable field for every branch of the disjunction,
	// add it to preprocessor.types, and use it instead.
	newTypeName := preprocessor.disjunctionTypeName(def.Branches)

	if _, ok := preprocessor.defs[newTypeName]; !ok {
		newType := ast.Definition{
			Kind: ast.KindStruct,
			Name: newTypeName,
		}

		for _, branch := range def.Branches {
			if branch.Kind == ast.KindNull {
				continue
			}

			newType.Fields = append(newType.Fields, ast.FieldDefinition{
				Name: "Val" + strings.Title(string(branch.Kind)),
				Type: ast.Definition{
					Nullable: true,
					Kind:     branch.Kind,

					IndexType:   branch.IndexType,
					ValueType:   branch.ValueType,
					Constraints: branch.Constraints,
				},
				Required: false,
			})
		}

		preprocessor.defs[newTypeName] = newType
	}

	return ast.Definition{
		Kind:     ast.Kind(newTypeName),
		Nullable: def.Branches.HasNullType(),
	}
}

func (preprocessor *preprocessor) disjunctionTypeName(disjunctionTypes ast.Definitions) string {
	parts := make([]string, 0, len(disjunctionTypes))

	for _, subType := range disjunctionTypes {
		parts = append(parts, strings.Title(string(subType.Kind)))
	}

	return strings.Title(strings.Join(parts, "Or"))
}
