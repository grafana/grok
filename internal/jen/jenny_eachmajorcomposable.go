package jen

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/grafana/codejen"
	"github.com/grafana/grafana/pkg/codegen"
)

// TODO remove this once there's a standard jenny for this...somewhere in core
func ComposableLatestMajorsOrXJenny(parentdir string, inner codejen.OneToOne[codegen.SchemaForGen]) codejen.OneToMany[*ComposableForGen] {
	if inner == nil {
		panic("inner jenny must not be nil")
	}

	return &clmox{
		parentdir: parentdir,
		inner:     inner,
	}
}

type clmox struct {
	parentdir string
	inner     codejen.OneToOne[codegen.SchemaForGen]
}

func (j *clmox) JennyName() string {
	return "ComposableLatestMajorsOrXJenny"
}

func (j *clmox) Generate(decl *ComposableForGen) (codejen.Files, error) {
	decl.Info.Meta()
	sfg := codegen.SchemaForGen{
		Name:    decl.Slot.Name(),
		IsGroup: decl.Slot.IsGroup(),
	}

	nam := fmt.Sprintf("%s-%s", strings.ToLower(decl.Info.Meta().Id), strings.ToLower(sfg.Name))

	do := func(sfg codegen.SchemaForGen, infix string) (codejen.Files, error) {
		f, err := j.inner.Generate(sfg)
		if err != nil {
			return nil, fmt.Errorf("%s jenny failed on %s schema for %s: %w", j.inner.JennyName(), sfg.Schema.Version(), nam, err)
		}
		if f == nil || !f.Exists() {
			return nil, nil
		}

		f.RelativePath = filepath.Join(j.parentdir, strings.ToLower(decl.Info.Meta().Id), strings.ToLower(sfg.Name), infix, f.RelativePath)
		f.From = append(f.From, j)
		return codejen.Files{*f}, nil
	}

	// TODO uncomment this latter half once plugins are fully converted to new system
	// if comm.Maturity.Less(kindsys.MaturityStable) {
	sfg.Schema = decl.Lineage.Latest()
	return do(sfg, "x")
	// }

	// var fl codejen.Files
	// for sch := decl.Lineage.First(); sch != nil; sch.Successor() {
	// 	sfg.Schema = sch.LatestInMajor()
	// 	files, err := do(sfg, fmt.Sprintf("v%v", sch.Version()[0]))
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	fl = append(fl, files...)
	// }
	// if fl.Validate() != nil {
	// 	return nil, fl.Validate()
	// }
	// return fl, nil
}
