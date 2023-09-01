package compiler

import (
	"github.com/grafana/grok/internal/sandbox/gen/ast"
	"github.com/grafana/grok/internal/sandbox/gen/rewrite"
)

var _ Pass = (*Rewrite)(nil)

type Rewrite struct {
	fieldRules []rewrite.FieldRewriteRule
}

func NewRewrite(fieldRules []rewrite.FieldRewriteRule) *Rewrite {
	return &Rewrite{
		fieldRules: fieldRules,
	}
}

func (pass *Rewrite) Process(file *ast.File) (*ast.File, error) {
	processedObjects := make([]ast.Object, 0, len(file.Definitions))
	for _, object := range file.Definitions {
		processedObjects = append(processedObjects, pass.processObject(object))
	}

	return &ast.File{
		Package:     file.Package,
		Definitions: processedObjects,
	}, nil
}

func (pass *Rewrite) processObject(object ast.Object) ast.Object {
	newObject := object
	newObject.Type = pass.processType(object, object.Type)

	return newObject
}

func (pass *Rewrite) processType(parentObject ast.Object, def ast.Type) ast.Type {
	if def.Kind() == ast.KindStruct {
		return pass.processStruct(parentObject, def.(*ast.StructType))
	}

	return def
}

func (pass *Rewrite) processStruct(parentObject ast.Object, def *ast.StructType) *ast.StructType {
	newDef := def

	processedFields := make([]ast.StructField, 0, len(def.Fields))
	for _, field := range def.Fields {
		newFields := []ast.StructField{field}

		for _, rule := range pass.fieldRules {
			if rule.Selector(parentObject, field) {
				// FIXME: what am I even doing here?
				var moarFields []ast.StructField

				for _, modifiedField := range newFields {
					moarFields = append(moarFields, rule.Action(modifiedField)...)
				}

				newFields = moarFields
			}
		}

		processedFields = append(processedFields, newFields...)
	}

	newDef.Fields = processedFields

	return newDef
}
