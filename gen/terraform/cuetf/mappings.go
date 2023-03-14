package cuetf

import "cuelang.org/go/cue"

var TypeMappings = map[cue.Kind]string{
	cue.BoolKind:   "Bool",
	cue.IntKind:    "Int64",
	cue.FloatKind:  "Float64",
	cue.NumberKind: "Float64",
	cue.StringKind: "String",
}

var GolangTypeMappings = map[cue.Kind]string{
	cue.BoolKind:   "bool",
	cue.IntKind:    "int64",
	cue.FloatKind:  "float64",
	cue.NumberKind: "float64",
	cue.StringKind: "string",
}

var TerraformFuncTypeMappings = map[cue.Kind]string{
	cue.BoolKind:   "ValueBool()",
	cue.IntKind:    "ValueInt64()",
	cue.FloatKind:  "ValueFloat64()",
	cue.NumberKind: "ValueFloat64()",
	cue.StringKind: "ValueString()",
}
