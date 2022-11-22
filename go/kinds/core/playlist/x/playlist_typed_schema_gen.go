package playlist

import (
	"fmt"
	"sync"

	"github.com/grafana/grafana/pkg/registry/corekind"
	"github.com/grafana/thema"
)

// KindVersion is the syntactic version of the canonical schema from which this
// package's '[Playlist] type was generated.
//
// This value will always be the same as calling Schema().Version().
var KindVersion thema.SyntacticVersion = [2]uint{0, 0}

var schemaOnce sync.Once
var tschema thema.TypedSchema[*Playlist]

// Schema returns the [thema.TypedSchema] representing the schema from which
// the [Playlist] type and its referenced subtypes (if any) were generated.
func Schema() thema.TypedSchema[*Playlist] {
	reg := corekind.NewBase(nil)
	sch := thema.SchemaP(reg.Playlist().Lineage(), thema.SV(0, 0))
	schemaOnce.Do(func() {
		var err error
		tschema, err = thema.BindType[*Playlist](sch, new(Playlist))
		if err != nil {
			panic(fmt.Sprintf("playlist@0.0 is not assignable to *playlist: %s", err))
		}
	})
	return tschema
}

// NewWithDefaults returns a new [Playlist] that has been initialized with
// schema-specified default values.
func NewWithDefaults() *Playlist {
	return Schema().NewT()
}
