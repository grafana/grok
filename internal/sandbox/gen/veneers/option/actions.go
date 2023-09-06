package option

import (
	"github.com/grafana/grok/internal/sandbox/gen/ast"
)

type RewriteAction func(option ast.Option) []ast.Option

func RenameAction(newName string) RewriteAction {
	return func(option ast.Option) []ast.Option {
		newOption := option
		newOption.Name = newName

		return []ast.Option{newOption}
	}
}

func OmitAction() RewriteAction {
	return func(_ ast.Option) []ast.Option {
		return nil
	}
}

func PromoteToConstructorAction() RewriteAction {
	return func(option ast.Option) []ast.Option {
		newOpt := option
		newOpt.IsConstructorArg = true

		return []ast.Option{newOpt}
	}
}

// FIXME: looks at the first arg only, no way to configure that right now
func StructFieldsAsArgumentsAction() RewriteAction {
	return func(option ast.Option) []ast.Option {
		// do nothing if we can't do anything.
		if len(option.Args) < 1 || option.Args[0].Type.Kind() != ast.KindStruct {
			return []ast.Option{option}
		}

		oldArgs := option.Args
		oldAssignments := option.Assignments
		assignmentPathPrefix := oldAssignments[0].Path
		structType := option.Args[0].Type.(ast.StructType)

		newOpt := option
		newOpt.Args = nil
		newOpt.Assignments = nil

		for _, field := range structType.Fields {
			var constraints []ast.TypeConstraint
			if scalarType, ok := field.Type.(ast.ScalarType); ok {
				constraints = scalarType.Constraints
			}

			newOpt.Args = append(newOpt.Args, ast.Argument{
				Name: field.Name,
				Type: field.Type,
			})

			newOpt.Assignments = append(newOpt.Assignments, ast.Assignment{
				Path:              assignmentPathPrefix + "." + field.Name,
				ArgumentName:      field.Name,
				ValueType:         field.Type,
				Constraints:       constraints,
				IntoOptionalField: !field.Required,
			})
		}

		if len(oldArgs) > 1 {
			newOpt.Args = append(newOpt.Args, oldArgs[1:]...)
			newOpt.Assignments = append(newOpt.Assignments, oldAssignments[1:]...)
		}

		return []ast.Option{newOpt}
	}
}

type BooleanUnfold struct {
	OptionTrue  string
	OptionFalse string
}

func UnfoldBooleanAction(unfoldOpts BooleanUnfold) RewriteAction {
	return func(option ast.Option) []ast.Option {
		return []ast.Option{
			{
				Name:     unfoldOpts.OptionTrue,
				Comments: option.Comments,
				Args:     nil,
				Assignments: []ast.Assignment{
					{
						Path:              option.Assignments[0].Path,
						ValueType:         option.Assignments[0].ValueType,
						IntoOptionalField: option.Assignments[0].IntoOptionalField,
						Value:             true,
					},
				},
				// TODO: default
			},

			{
				Name:     unfoldOpts.OptionFalse,
				Comments: option.Comments,
				Args:     nil,
				Assignments: []ast.Assignment{
					{
						Path:              option.Assignments[0].Path,
						ValueType:         option.Assignments[0].ValueType,
						IntoOptionalField: option.Assignments[0].IntoOptionalField,
						Value:             false,
					},
				},
				// TODO: default
			},
		}
	}
}
