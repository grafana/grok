package dashboard

import (
	"fmt"
	"sync"

	"github.com/grafana/grafana/pkg/registry/corekind"
	"github.com/grafana/thema"
)

// KindVersion is the syntactic version of the canonical schema from which this
// package's '[Dashboard] type was generated.
//
// This value will always be the same as calling Schema().Version().
var KindVersion thema.SyntacticVersion = [2]uint{0, 0}

var schemaOnce sync.Once
var tschema thema.TypedSchema[*Dashboard]

// Schema returns the [thema.TypedSchema] representing the schema from which
// the [Dashboard] type and its referenced subtypes (if any) were generated.
func Schema() thema.TypedSchema[*Dashboard] {
	reg := corekind.NewBase(nil)
	sch := thema.SchemaP(reg.Dashboard().Lineage(), thema.SV(0, 0))
	schemaOnce.Do(func() {
		var err error
		tschema, err = thema.BindType[*Dashboard](sch, new(Dashboard))
		if err != nil {
			panic(fmt.Sprintf("dashboard@0.0 is not assignable to *dashboard: %s", err))
		}
	})
	return tschema
}

// NewWithDefaults returns a new [Dashboard] that has been initialized with
// schema-specified default values.
func NewWithDefaults() *Dashboard {
	return Schema().NewT()
}
