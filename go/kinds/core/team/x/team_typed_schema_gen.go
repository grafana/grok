package team

import (
	"fmt"
	"sync"

	"github.com/grafana/grafana/pkg/registry/corekind"
	"github.com/grafana/thema"
)

// KindVersion is the syntactic version of the canonical schema from which this
// package's '[Team] type was generated.
//
// This value will always be the same as calling Schema().Version().
var KindVersion thema.SyntacticVersion = [2]uint{0, 0}

var schemaOnce sync.Once
var tschema thema.TypedSchema[*Team]

// Schema returns the [thema.TypedSchema] representing the schema from which
// the [Team] type and its referenced subtypes (if any) were generated.
func Schema() thema.TypedSchema[*Team] {
	reg := corekind.NewBase(nil)
	sch := thema.SchemaP(reg.Team().Lineage(), thema.SV(0, 0))
	schemaOnce.Do(func() {
		var err error
		tschema, err = thema.BindType[*Team](sch, new(Team))
		if err != nil {
			panic(fmt.Sprintf("team@0.0 is not assignable to *team: %s", err))
		}
	})
	return tschema
}

// NewWithDefaults returns a new [Team] that has been initialized with
// schema-specified default values.
func NewWithDefaults() *Team {
	return Schema().NewT()
}