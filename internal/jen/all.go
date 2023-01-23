package jen

import (
	"path/filepath"

	"github.com/grafana/codejen"
	"github.com/grafana/grafana/pkg/codegen"
	"github.com/grafana/grafana/pkg/kindsys"
)

// TargetJennies is a set of jennies for a particular target language or
// tool that perform all necessary code generation steps.
type TargetJennies struct {
	Core *codejen.JennyList[*codegen.DeclForGen]
	// TODO replace pfs.PluginInfo with type from kindsys once implemented
	Composable *codejen.JennyList[kindsys.Composable]
}

// NewTargetJennies initializes a new TargetJennies with appropriate namers for
// each JennyList.
func NewTargetJennies() TargetJennies {
	return TargetJennies{
		Core: codejen.JennyListWithNamer[*codegen.DeclForGen](func(decl *codegen.DeclForGen) string {
			return decl.Properties.Common().MachineName
		}),
		Composable: codejen.JennyListWithNamer[kindsys.Composable](func(k kindsys.Composable) string {
			return k.Name()
		}),
	}
}

// Prefixer returns a FileMapper that injects the provided path prefix to files
// passed through it.
func Prefixer(prefix string) codejen.FileMapper {
	return func(f codejen.File) (codejen.File, error) {
		f.RelativePath = filepath.Join(prefix, f.RelativePath)
		return f, nil
	}
}
