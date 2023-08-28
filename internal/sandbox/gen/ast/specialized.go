package ast

// interface for every type that we can represent:
// struct, enum, array, string, int, ...
type Type interface {
	Kind() Kind
}

// named declaration of a type
type Object struct {
	Name     string
	Comments []string
	Type     Type
}

type File struct {
	Package     string
	Definitions []Object
}

func (file *File) LocateDefinition(name string) Object {
	for _, def := range file.Definitions {
		if def.Name == name {
			return def
		}
	}

	return Object{}
}

var _ Type = (*DisjunctionType)(nil)

type DisjunctionType struct {
	Branches []Type
}

func (disjunctionType DisjunctionType) Kind() Kind {
	return KindDisjunction
}

var _ Type = (*ArrayType)(nil)

type ArrayType struct {
	ValueType Type
}

func (arrayType ArrayType) Kind() Kind {
	return KindArray
}

var _ Type = (*EnumType)(nil)

type EnumType struct {
	Values []EnumValue // possible values. Value types might be different
}

type EnumValue struct {
	Type  Type
	Name  string
	Value any
}

func (arrayType EnumType) Kind() Kind {
	return KindEnum
}

var _ Type = (*MapType)(nil)

type MapType struct {
	IndexType Type
	ValueType Type
}

func (arrayType MapType) Kind() Kind {
	return KindMap
}

var _ Type = (*StructType)(nil)

type StructType struct {
	Fields []StructField
}

func (structType StructType) Kind() Kind {
	return KindStruct
}

type StructField struct {
	Name     string
	Comments []string
	Type     Type
	Required bool
}

var _ Type = (*RefType)(nil)

type RefType struct {
	ReferredType string
}

func (refType RefType) Kind() Kind {
	return KindRef
}

var _ Type = (*ScalarType)(nil)

type ScalarType struct {
	ScalarKind Kind // bool, bytes, string, int, float, ...
}

func (scalarType ScalarType) Kind() Kind {
	return scalarType.ScalarKind
}

var _ Type = (*Literal)(nil)

type Literal struct {
	ScalarType ScalarType
	Value      any
}

func (literal Literal) Kind() Kind {
	return literal.ScalarType.Kind()
}
