package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"testing/fstest"

	"github.com/grafana/kindsys"
	"github.com/grafana/thema"
	"github.com/yalue/merged_fs"
)

func moduleToComposableKind(themaRuntime *thema.Runtime, commonFS fs.FS, modulePath string) (kindsys.Composable, error) {
	fmt.Printf(" → Loading %s\n", modulePath)

	moduleHandle, err := os.ReadDir(modulePath)
	if err != nil {
		return nil, fmt.Errorf("could not open module '%s' from registry: %w", modulePath, err)
	}

	moduleFS := fstest.MapFS{}
	for _, file := range moduleHandle {
		if file.IsDir() {
			continue
		}

		content, err := os.ReadFile(filepath.Join(modulePath, file.Name()))
		if err != nil {
			return nil, err
		}

		// weird CUE errors if we don't add this
		//panic: schemaInterface: cannot convert non-concrete value string:
		//	/github.com/grafana/kindsys/kindcat_composable.cue:18:2
		//schif: invalid non-ground value string (must be concrete string):
		//	/github.com/grafana/kindsys/kindcat_composable.cue:18:19
		//name: cannot convert non-concrete value =~"^([A-Z][a-zA-Z0-9-]{0,61}[a-zA-Z0-9])$":
		//	/github.com/grafana/kindsys/kindcats.cue:38:2
		//machineName: error in call to strings.Replace: non-concrete value string:
		//	/github.com/grafana/kindsys/kindcats.cue:46:31
		//	/github.com/grafana/kindsys/kindcats.cue:38:8
		//pluralName: cannot convert non-concrete value =~"^([A-Z][a-zA-Z0-9-]{0,61}[a-zA-Z])$":
		//	/github.com/grafana/kindsys/kindcats.cue:49:2
		//pluralMachineName: error in call to strings.Replace: non-concrete value string:
		//	/github.com/grafana/kindsys/kindcats.cue:54:37
		//	/github.com/grafana/kindsys/kindcats.cue:49:14
		moduleFS[file.Name()] = &fstest.MapFile{
			Data: []byte(fmt.Sprintf("package grafanaplugin\n%s", content)),
		}
	}

	overlayFS := merged_fs.NewMergedFS(commonFS, moduleFS)

	cueInstance, err := kindsys.BuildInstance(themaRuntime.Context(), ".", "grafanaplugin", overlayFS)
	if err != nil {
		return nil, fmt.Errorf("could not load kindsys instance: %w", err)
	}

	props, err := kindsys.ToKindProps[kindsys.ComposableProperties](cueInstance)
	if err != nil {
		return nil, fmt.Errorf("could not convert cue value to kindsys props: %w", err)
	}

	kindDefinition := kindsys.Def[kindsys.ComposableProperties]{
		V:          cueInstance,
		Properties: props,
	}

	boundKind, err := kindsys.BindComposable(themaRuntime, kindDefinition)
	if err != nil {
		return nil, fmt.Errorf("could not bind kind definition to kind: %w", err)
	}

	return boundKind, nil
}

func loadComposableKinds(themaRuntime *thema.Runtime, commonFS fs.FS, composableKindsPath string) ([]kindsys.Composable, error) {
	files, err := os.ReadDir(composableKindsPath)
	if err != nil {
		return nil, fmt.Errorf("could not open registry: %w", err)
	}

	composableKinds := make([]kindsys.Composable, 0, len(files))
	for _, file := range files {
		if !file.IsDir() {
			continue
		}

		// TODO: figure out why loading this one hangs and remove
		if file.Name() == "elasticsearch" {
			continue
		}

		composableKind, err := moduleToComposableKind(themaRuntime, commonFS, filepath.Join(composableKindsPath, file.Name()))
		if err != nil {
			return nil, err
		}

		composableKinds = append(composableKinds, composableKind)
	}

	return composableKinds, nil
}
