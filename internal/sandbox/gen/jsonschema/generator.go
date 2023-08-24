package jsonschema

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/grafana/grok/internal/sandbox/gen/ast"
	schemaparser "github.com/santhosh-tekuri/jsonschema"
)

var errUndescriptiveSchema = fmt.Errorf("the schema does not appear to be describing anything")

const (
	typeNull    = "null"
	typeBoolean = "boolean"
	typeObject  = "object"
	typeArray   = "array"
	typeString  = "string"
	typeNumber  = "number"
	typeInteger = "integer"
)

type Config struct {
	// Package name used to generate code into.
	Package string
}

type newGenerator struct {
	file *ast.File
}

func GenerateAST(schemaReader io.Reader, c Config) (*ast.File, error) {
	g := &newGenerator{
		file: &ast.File{
			Package: c.Package,
		},
	}

	compiler := schemaparser.NewCompiler()
	compiler.ExtractAnnotations = true
	if err := compiler.AddResource("schema", schemaReader); err != nil {
		return nil, err
	}

	schema, err := compiler.Compile("schema")
	if err != nil {
		return nil, err
	}

	// The root of the schema is an actual type/object
	if schema.Ref == nil {
		if err := g.declareDefinition(c.Package, schema); err != nil {
			return nil, err
		}
	} else {
		// The root of the schema contains definitions, and a reference to the "main" object
		if err := g.declareDefinition(c.Package, schema.Ref); err != nil {
			return nil, err
		}
	}

	return g.file, nil
}

func (g *newGenerator) declareDefinition(definitionName string, schema *schemaparser.Schema) error {
	def, err := g.walkDefinition(schema)
	if err != nil {
		return fmt.Errorf("%s: %w", definitionName, err)
	}

	def.Name = definitionName

	g.file.Definitions = append(g.file.Definitions, *def)

	return nil
}

func (g *newGenerator) walkDefinition(schema *schemaparser.Schema) (*ast.Definition, error) {
	var def *ast.Definition
	var err error

	if len(schema.Types) == 0 {
		if schema.Ref != nil {
			return g.walkRef(schema)
		}

		if schema.OneOf != nil {
			return g.walkOneOf(schema)
		}

		if schema.AnyOf != nil {
			return g.walkAnyOf(schema)
		}

		if schema.AllOf != nil {
			return g.walkOneOf(schema)
		}

		if schema.Properties != nil || schema.PatternProperties != nil {
			return g.walkObject(schema)
		}

		if schema.Enum != nil {
			return g.walkEnum(schema)
		}

		return nil, errUndescriptiveSchema
	}

	if len(schema.Types) > 1 {
		def, err = g.walkDisjunction(schema)
	} else if schema.Enum != nil {
		def, err = g.walkEnum(schema)
	} else {
		switch schema.Types[0] {
		case typeNull:
			def = &ast.Definition{Kind: ast.KindNull, Comments: schemaComments(schema)}
		case typeBoolean:
			def = &ast.Definition{Kind: ast.KindBool, Comments: schemaComments(schema)}
		case typeString:
			def, err = g.walkString(schema)
		case typeObject:
			def, err = g.walkObject(schema)
		case typeNumber, typeInteger:
			def, err = g.walkNumber(schema)
		case typeArray:
			def, err = g.walkList(schema)

		default:
			return nil, fmt.Errorf("unexpected schema with type '%s'", schema.Types[0])
		}
	}

	return def, err
}

func (g *newGenerator) walkDisjunction(schema *schemaparser.Schema) (*ast.Definition, error) {
	// TODO: finish implementation
	return &ast.Definition{Kind: ast.KindDisjunction, Comments: schemaComments(schema)}, nil
}

func (g *newGenerator) walkDisjunctionBranches(branches []*schemaparser.Schema) (ast.Definitions, error) {
	definitions := make([]ast.Definition, 0, len(branches))
	for _, oneOf := range branches {
		branch, err := g.walkDefinition(oneOf)
		if err != nil {
			return nil, err
		}

		definitions = append(definitions, *branch)
	}

	return definitions, nil
}

func (g *newGenerator) walkOneOf(schema *schemaparser.Schema) (*ast.Definition, error) {
	if len(schema.OneOf) == 0 {
		return nil, fmt.Errorf("oneOf with no branches")
	}

	branches, err := g.walkDisjunctionBranches(schema.OneOf)
	if err != nil {
		return nil, err
	}

	return &ast.Definition{
		Kind:     ast.KindDisjunction,
		Comments: schemaComments(schema),
		Branches: branches,
	}, nil
}

// TODO: what's the difference between oneOf and anyOf?
func (g *newGenerator) walkAnyOf(schema *schemaparser.Schema) (*ast.Definition, error) {
	if len(schema.AnyOf) == 0 {
		return nil, fmt.Errorf("anyOf with no branches")
	}

	branches, err := g.walkDisjunctionBranches(schema.AnyOf)
	if err != nil {
		return nil, err
	}

	return &ast.Definition{
		Kind:     ast.KindDisjunction,
		Comments: schemaComments(schema),
		Branches: branches,
	}, nil
}

func (g *newGenerator) walkAllOf(schema *schemaparser.Schema) (*ast.Definition, error) {
	// TODO: finish implementation and use correct type
	return &ast.Definition{Kind: ast.KindDisjunction, Comments: schemaComments(schema)}, nil
}

func (g *newGenerator) walkRef(schema *schemaparser.Schema) (*ast.Definition, error) {
	parts := strings.Split(schema.Ref.Ptr, "/")
	referredKindName := parts[len(parts)-1] // Very naive

	if err := g.declareDefinition(referredKindName, schema.Ref); err != nil {
		return nil, err
	}

	return &ast.Definition{
		Kind:     ast.Kind(referredKindName),
		Comments: schemaComments(schema),
	}, nil
}

func (g *newGenerator) walkString(schema *schemaparser.Schema) (*ast.Definition, error) {
	def := &ast.Definition{Kind: ast.KindString, Comments: schemaComments(schema)}

	if len(schema.Enum) != 0 {
		def.Constraints = append(def.Constraints, ast.TypeConstraint{
			Op:   "in",
			Args: []any{schema.Enum},
		})
	}

	return def, nil
}

func (g *newGenerator) walkNumber(schema *schemaparser.Schema) (*ast.Definition, error) {
	// TODO: finish implementation
	return &ast.Definition{Kind: ast.KindInt64, Comments: schemaComments(schema)}, nil
}

func (g *newGenerator) walkList(schema *schemaparser.Schema) (*ast.Definition, error) {
	var itemsDef *ast.Definition
	var err error

	if schema.Items == nil {
		itemsDef = &ast.Definition{
			Kind: ast.KindAny,
		}
	} else {
		// TODO: schema.Items might not be a schema?
		itemsDef, err = g.walkDefinition(schema.Items.(*schemaparser.Schema))
		// items contains an empty schema: `{}`
		if errors.Is(err, errUndescriptiveSchema) {
			itemsDef = &ast.Definition{
				Kind: ast.KindAny,
			}
		} else if err != nil {
			return nil, err
		}
	}

	return &ast.Definition{
		Kind:      ast.KindArray,
		IndexType: ast.KindInt64,
		ValueType: itemsDef,

		Comments: schemaComments(schema),
	}, nil
}

func (g *newGenerator) walkEnum(schema *schemaparser.Schema) (*ast.Definition, error) {
	if len(schema.Enum) == 0 {
		return nil, fmt.Errorf("enum with no values")
	}

	values := make([]ast.EnumValue, 0, len(schema.Enum))
	for _, enumValue := range schema.Enum {
		values = append(values, ast.EnumValue{
			Type: ast.KindString, // TODO: identify that correctly

			// Simple mapping of all enum values (which we are assuming are in
			// lowerCamelCase) to corresponding CamelCase
			Name:  enumValue.(string),
			Value: enumValue.(string),
		})
	}

	return &ast.Definition{
		Kind:     ast.KindEnum,
		Comments: schemaComments(schema),
		Values:   values,
		// TODO: default value?
	}, nil
}

func (g *newGenerator) walkObject(schema *schemaparser.Schema) (*ast.Definition, error) {
	// TODO: finish implementation
	def := &ast.Definition{Kind: ast.KindStruct, Comments: schemaComments(schema)}

	for name, property := range schema.Properties {
		fieldDef, err := g.walkDefinition(property)
		if err != nil {
			return nil, err
		}

		def.Fields = append(def.Fields, ast.FieldDefinition{
			Name:     name,
			Comments: schemaComments(schema),
			Required: stringInList(schema.Required, name),
			Type:     *fieldDef,
		})
	}

	return def, nil
}
