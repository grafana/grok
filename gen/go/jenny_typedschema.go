package _go

import (
	"bytes"
	"text/template"

	"github.com/grafana/codejen"
	"github.com/grafana/grafana/pkg/codegen"
)

// TypedSchemaJenny generates a func that returns a [thema.TypedLineage] for a
// given schema. Implicitly depends on the output of [codegen.GoTypesJenny].
//
// TODO this could probably be upstreamed to thema eventually
type TypedSchemaJenny struct{}

func (j TypedSchemaJenny) JennyName() string {
	return "TypedSchemaJenny"
}

func (j TypedSchemaJenny) Generate(sfg codegen.SchemaForGen) (*codejen.File, error) {
	lin := sfg.Schema.Lineage()
	vars := map[string]any{
		"PackageName": lin.Name(),
		"Majv":        sfg.Schema.Version()[0],
		"Minv":        sfg.Schema.Version()[1],
		"TitleName":   sfg.Name,
		"Name":        lin.Name(),
	}

	buf := new(bytes.Buffer)
	err := tmplTypedSchema.Execute(buf, vars)
	if err != nil {
		return nil, err
	}

	return codejen.NewFile(lin.Name()+"_typed_schema_gen.go", buf.Bytes(), j), nil
}

var tmplTypedSchema = template.Must(template.New("typedsch").Parse(`package {{ .PackageName }}

import (
	"fmt"
	"sync"

	"github.com/grafana/grafana/pkg/registry/corekind"
	"github.com/grafana/thema"
)

// KindVersion is the syntactic version of the canonical schema from which this
// package's '[{{ .TitleName }}] type was generated.
//
// This value will always be the same as calling Schema().Version().
var KindVersion thema.SyntacticVersion = [2]uint{ {{- .Majv }}, {{ .Minv -}} }

var schemaOnce sync.Once
var tschema thema.TypedSchema[*{{ .TitleName }}]

// Schema returns the [thema.TypedSchema] representing the schema from which
// the [{{ .TitleName }}] type and its referenced subtypes (if any) were generated.
func Schema() thema.TypedSchema[*{{ .TitleName }}] {
	reg := corekind.NewBase(nil)
	sch := thema.SchemaP(reg.{{ .TitleName }}().Lineage(), thema.SV({{ .Majv }}, {{ .Minv }}))
	schemaOnce.Do(func() {
		var err error
		tschema, err = thema.BindType[*{{ .TitleName }}](sch, new({{ .TitleName }}))
		if err != nil {
			panic(fmt.Sprintf("{{ .Name }}@{{ .Majv }}.{{ .Minv }} is not assignable to *{{ .Name }}: %s", err))
		}
	})
	return tschema
}

// NewWithDefaults returns a new [{{ .TitleName }}] that has been initialized with
// schema-specified default values.
func NewWithDefaults() *{{ .TitleName }} {
	return Schema().NewT()
}
`))
