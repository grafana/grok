package rewrite

import (
	"github.com/grafana/grok/internal/sandbox/gen/ast"
)

type ObjectSelector func(obj ast.Object) bool

type FieldSelector func(object ast.Object, field ast.StructField) bool

func ExactField(objectName string, fieldName string) FieldSelector {
	return func(object ast.Object, field ast.StructField) bool {
		return object.Name == objectName && field.Name == fieldName
	}
}

func EveryField() FieldSelector {
	return func(object ast.Object, field ast.StructField) bool {
		return true
	}
}

type FieldRewriteAction func(field ast.StructField) []ast.StructField

func Rename(newName string) FieldRewriteAction {
	return func(field ast.StructField) []ast.StructField {
		newField := field
		newField.DisplayName = newName

		return []ast.StructField{newField}
	}
}

func Omit() FieldRewriteAction {
	return func(field ast.StructField) []ast.StructField {
		return nil
	}
}

type FieldRewriteRule struct {
	Selector FieldSelector
	Action   FieldRewriteAction
}

func RenameField(fieldSelector FieldSelector, newName string) FieldRewriteRule {
	return FieldRewriteRule{
		Selector: fieldSelector,
		Action:   Rename(newName),
	}
}

func OmitField(fieldSelector FieldSelector) FieldRewriteRule {
	return FieldRewriteRule{
		Selector: fieldSelector,
		Action:   Omit(),
	}
}

func SetDefaultDisplayName() FieldRewriteRule {
	return FieldRewriteRule{
		Selector: EveryField(),
		Action: func(field ast.StructField) []ast.StructField {
			newField := field
			newField.DisplayName = field.Name

			return []ast.StructField{newField}
		},
	}
}

type BooleanUnfold struct {
	OptionTrue  string
	OptionFalse string
}

func UnfoldBoolean(fieldSelector FieldSelector, unfoldOpts BooleanUnfold) FieldRewriteRule {
	return FieldRewriteRule{
		Selector: fieldSelector,
		Action: func(field ast.StructField) []ast.StructField {
			return []ast.StructField{
				{
					Name:        field.Name,
					DisplayName: unfoldOpts.OptionTrue,
					Comments:    field.Comments,
					Type: &ast.Literal{
						ScalarType: ast.ScalarType{ScalarKind: ast.KindBool},
						Value:      true,
					},
				},
				{
					Name:        field.Name,
					DisplayName: unfoldOpts.OptionFalse,
					Comments:    field.Comments,
					Type: &ast.Literal{
						ScalarType: ast.ScalarType{ScalarKind: ast.KindBool},
						Value:      false,
					},
				},
			}
		},
	}
}
