package cuetf

import "cuelang.org/go/cue"

var TypeMappings = map[cue.Kind]string{
	cue.BoolKind:   "Bool",
	cue.IntKind:    "Int64",
	cue.FloatKind:  "Float64",
	cue.NumberKind: "Number",
	cue.StringKind: "String",
	cue.StructKind: "Object",
	cue.ListKind:   "List",
}
