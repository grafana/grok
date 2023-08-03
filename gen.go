// //go:build ignore
// // +build ignore

//go:generate go run gen.go

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/grafana/codejen"
	"github.com/grafana/grafana/pkg/plugins/pfs/corelist"
	"github.com/grafana/grafana/pkg/registry/corekind"
	"github.com/grafana/grok/gen/jsonnet"
	"github.com/grafana/grok/gen/jsonschema"
	"github.com/grafana/grok/internal/jen"
	"github.com/grafana/kindsys"
	"github.com/grafana/kindsys/pkg/codegen"
)

func main() {
	grafanaVersion := os.Getenv("GRAFANA_VERSION")
	if grafanaVersion == "" {
		panic("GRAFANA_VERSION environment variable must be set")
	}

	jfs := codejen.NewFS()

	// i've got a lovely bunch of coconuts
	// there they are all standing in a row
	coco := lineUpJennies(grafanaVersion)

	var corek []kindsys.Kind
	var compok []kindsys.Composable

	minMaturity :=
		func() kindsys.Maturity {
			switch os.Getenv("MIN_MATURITY") {
			case "merged":
				return kindsys.MaturityMerged
			case "experimental":
				return kindsys.MaturityExperimental
			case "stable":
				return kindsys.MaturityStable
			case "mature":
				return kindsys.MaturityMature
			default:
				return kindsys.MaturityExperimental
			}
		}()

	for _, kind := range corekind.NewBase(nil).All() {
		if kind.Maturity().Less(minMaturity) {
			continue
		}
		corek = append(corek, kind)
	}
	for _, pp := range corelist.New(nil) {
		for _, kind := range pp.ComposableKinds {
			if kind.Maturity().Less(minMaturity) {
				continue
			}
			compok = append(compok, kind)
		}
	}

	ckfs, err := coco.Core.GenerateFS(corek...)
	die(err)
	die(jfs.Merge(ckfs))
	ckfs, err = coco.Composable.GenerateFS(compok...)
	die(err)
	die(jfs.Merge(ckfs))

	if _, set := os.LookupEnv("CODEGEN_VERIFY"); set {
		if err = jfs.Verify(context.Background(), ""); err != nil {
			die(fmt.Errorf("generated code is out of sync with inputs:\n%s\nrun `make gen-cue` to regenerate", err))
		}
	} else if err = jfs.Write(context.Background(), ""); err != nil {
		die(fmt.Errorf("error while writing generated code to disk:\n%s", err))
	}
}

// Line up all the jennies from all the language targets, prefixing them with
// their lang target subpaths.
func lineUpJennies(grafanaVersion string) jen.TargetJennies {
	tgt := jen.NewTargetJennies()

	tgtmap := map[string]jen.TargetJennies{
		//"go":         _go.JenniesForGo(grafanaVersion), // This is not ready yet
		"jsonschema": jsonschema.JenniesForJsonSchema(grafanaVersion),
		"jsonnet":    jsonnet.JenniesForJsonnet(grafanaVersion),
	}

	for path, tj := range tgtmap {
		tj.Core.AddPostprocessors(codegen.Prefixer(path), codegen.SlashHeaderMapper(path))
		tj.Composable.AddPostprocessors(codegen.Prefixer(path), codegen.SlashHeaderMapper(path))
		tgt.Core.AppendManyToMany(tj.Core)
		tgt.Composable.AppendManyToMany(tj.Composable)
	}

	return tgt
}

func die(err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, err, "\n")
		os.Exit(1)
	}
}
