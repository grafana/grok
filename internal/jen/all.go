package jen

import (
	"fmt"
	"path/filepath"

	"github.com/grafana/codejen"
	"github.com/grafana/grafana/pkg/codegen"
	"github.com/grafana/grafana/pkg/kindsys"
	"github.com/grafana/grafana/pkg/plugins/pfs"
	"github.com/grafana/thema"
)

// TargetJennies is a set of jennies for a particular target language or
// tool that perform all necessary code generation steps.
type TargetJennies struct {
	Core *codejen.JennyList[*codegen.DeclForGen]
	// TODO replace pfs.PluginInfo with type from kindsys once implemented
	Composable *codejen.JennyList[*ComposableForGen]
}

// NewTargetJennies initializes a new TargetJennies with appropriate namers for
// each JennyList.
func NewTargetJennies() TargetJennies {
	return TargetJennies{
		Core: codejen.JennyListWithNamer[*codegen.DeclForGen](func(decl *codegen.DeclForGen) string {
			return decl.Meta.Common().MachineName
		}),
		Composable: codejen.JennyListWithNamer[*ComposableForGen](func(cfg *ComposableForGen) string {
			return fmt.Sprintf("%s-%s", cfg.Info.Meta().Id, cfg.Slot.Name())
		}),
	}
}

// ComposablesFromTree gives a ComposableForGen from a pfs.Tree.
//
// Temporary until we have proper types for this in grafana/grafana itself.
func ComposablesFromTree(ptree *pfs.Tree) []*ComposableForGen {
	allslots := kindsys.AllSlots(nil)
	info := ptree.RootPlugin()
	var compok []*ComposableForGen
	for slot, lin := range info.SlotImplementations() {
		compok = append(compok, &ComposableForGen{
			Info:    info,
			Slot:    *allslots[slot],
			Lineage: lin,
		})
	}

	return compok
}

// ComposableForGen is a codegen-friendly representation of a Grafana
// ComposableKind.
type ComposableForGen struct {
	Info    pfs.PluginInfo
	Slot    kindsys.Slot
	Lineage thema.Lineage
}

// Prefixer returns a FileMapper that injects the provided path prefix to files
// passed through it.
func Prefixer(prefix string) codejen.FileMapper {
	return func(f codejen.File) (codejen.File, error) {
		f.RelativePath = filepath.Join(prefix, f.RelativePath)
		return f, nil
	}
}
