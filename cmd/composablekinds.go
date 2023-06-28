package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing/fstest"

	"github.com/grafana/kindsys"
	"github.com/grafana/thema"
)

func moduleToComposableKind(themaRuntime *thema.Runtime, modulePath string) (kindsys.Composable, error) {
	fmt.Printf(" → Loading %s\n", modulePath)

	moduleHandle, err := os.ReadDir(modulePath)
	if err != nil {
		return nil, fmt.Errorf("could not open module '%s' from registry: %w", modulePath, err)
	}

	overlayFS := fstest.MapFS{}
	for _, file := range moduleHandle {
		if file.IsDir() {
			continue
		}

		content, err := os.ReadFile(filepath.Join(modulePath, file.Name()))
		if err != nil {
			return nil, err
		}

		overlayFS[file.Name()] = &fstest.MapFile{
			Data: []byte(fmt.Sprintf("package grafanaplugin\n%s", content)),
		}
	}

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

func loadComposableKinds(themaRuntime *thema.Runtime, composableKindsPath string) ([]kindsys.Composable, error) {
	files, err := os.ReadDir(composableKindsPath)
	if err != nil {
		return nil, fmt.Errorf("could not open registry: %w", err)
	}

	composableKinds := make([]kindsys.Composable, 0, len(files))
	for _, file := range files {
		if !file.IsDir() {
			continue
		}

		// TODO: remove
		if file.Name() != "alertgroups" {
			continue
		}

		composableKind, err := moduleToComposableKind(themaRuntime, filepath.Join(composableKindsPath, file.Name()))
		if err != nil {
			return nil, err
		}

		composableKinds = append(composableKinds, composableKind)
	}

	return composableKinds, nil
}
