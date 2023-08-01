package jen

import (
	"github.com/grafana/codejen"
	"github.com/grafana/kindsys"
)

// TargetJennies is a set of jennies for a particular target language or
// tool that perform all necessary code generation steps.
type TargetJennies struct {
	Core       *codejen.JennyList[kindsys.Kind]
	Composable *codejen.JennyList[kindsys.Composable]
}

// NewTargetJennies initializes a new TargetJennies with appropriate namers for
// each JennyList.
func NewTargetJennies() TargetJennies {
	return TargetJennies{
		Core: codejen.JennyListWithNamer[kindsys.Kind](func(k kindsys.Kind) string {
			return k.Props().Common().MachineName
		}),
		Composable: codejen.JennyListWithNamer[kindsys.Composable](func(k kindsys.Composable) string {
			return k.Name()
		}),
	}
}
