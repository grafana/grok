package ast

type Builder struct {
	For     Object
	Options []Option
}

type Option struct {
	Title            string
	Comments         []string
	Args             []Argument
	Assignments      []Assignment
	Default          *OptionDefault
	IsConstructorArg bool
}

type OptionDefault struct {
	ArgsValues []any
}

type Argument struct {
	Name           string
	Type           Type
	TypeHasBuilder bool
}

type Assignment struct {
	// Where
	Path string

	// What
	ValueType    Type   // type of the value being assigned
	ArgumentName string // if empty, then use `Value`
	Value        any

	Constraints []TypeConstraint

	// Some more context on the what
	IntoOptionalField bool
	ValueHasBuilder   bool
}

type BuilderGenerator struct {
}

func (generator *BuilderGenerator) FromAST(files []*File) []Builder {
	builders := make([]Builder, 0, len(files))

	for _, file := range files {
		for _, object := range file.Definitions {
			// we only want builders for structs
			if object.Type.Kind() != KindStruct {
				continue
			}

			builders = append(builders, generator.structObjectToBuilder(file, object))
		}
	}

	return builders
}

func (generator *BuilderGenerator) structObjectToBuilder(file *File, object Object) Builder {
	builder := Builder{
		For:     object,
		Options: nil,
	}
	structType := object.Type.(StructType)

	for _, field := range structType.Fields {
		builder.Options = append(builder.Options, generator.structFieldToOption(file, field))
	}

	return builder
}

func (generator *BuilderGenerator) structFieldToOption(file *File, field StructField) Option {
	valueHasBuilder := false
	if field.Type.Kind() == KindRef {
		referredDef := file.LocateDefinition(field.Type.(RefType).ReferredType)
		valueHasBuilder = referredDef.Type.Kind() == KindStruct
	}

	var constraints []TypeConstraint
	if scalarType, ok := field.Type.(ScalarType); ok {
		constraints = scalarType.Constraints
	}

	opt := Option{
		Title:    field.Name,
		Comments: field.Comments,
		Args: []Argument{
			{
				Name:           field.Name,
				Type:           field.Type,
				TypeHasBuilder: valueHasBuilder,
			},
		},
		Assignments: []Assignment{
			{
				Path:              field.Name,
				ArgumentName:      field.Name,
				ValueType:         field.Type,
				Constraints:       constraints,
				IntoOptionalField: !field.Required,
				ValueHasBuilder:   valueHasBuilder,
			},
		},
	}

	if field.Default != nil {
		opt.Default = &OptionDefault{
			ArgsValues: []any{field.Default},
		}
	}

	return opt
}
