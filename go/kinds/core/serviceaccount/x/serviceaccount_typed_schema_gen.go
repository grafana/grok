// Code generated - EDITING IS FUTILE. DO NOT EDIT.
//
// Generated by pipeline:
//     go
// Using jennies:
//     TypedSchemaJenny
//     LatestMajorsOrXJenny
//
// Run 'go generate ./' from repository root to regenerate.

package serviceaccount

import (
	"fmt"
	"sync"

	"github.com/grafana/grafana/pkg/registry/corekind"
	"github.com/grafana/thema"
)

// KindVersion is the syntactic version of the canonical schema from which this
// package's '[ServiceAccount] type was generated.
//
// This value will always be the same as calling Schema().Version().
var KindVersion thema.SyntacticVersion = [2]uint{0, 0}

var schemaOnce sync.Once
var tschema thema.TypedSchema[*ServiceAccount]

// Schema returns the [thema.TypedSchema] representing the schema from which
// the [ServiceAccount] type and its referenced subtypes (if any) were generated.
func Schema() thema.TypedSchema[*ServiceAccount] {
	reg := corekind.NewBase(nil)
	sch := thema.SchemaP(reg.ServiceAccount().Lineage(), thema.SV(0, 0))
	schemaOnce.Do(func() {
		var err error
		tschema, err = thema.BindType[*ServiceAccount](sch, new(ServiceAccount))
		if err != nil {
			panic(fmt.Sprintf("serviceaccount@0.0 is not assignable to *serviceaccount: %s", err))
		}
	})
	return tschema
}

// NewWithDefaults returns a new [ServiceAccount] that has been initialized with
// schema-specified default values.
func NewWithDefaults() *ServiceAccount {
	return Schema().NewT()
}