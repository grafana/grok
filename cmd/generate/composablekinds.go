package generate

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/grafana/kindsys"
	"github.com/grafana/thema"
	"github.com/yalue/merged_fs"
)

func moduleToComposableKind(themaRuntime *thema.Runtime, commonFS fs.FS, modulePath string) (kindsys.Composable, error) {
	fmt.Printf(" → Loading %s\n", modulePath)

	moduleFS, err := dirToPrefixedFS(modulePath, "")
	if err != nil {
		return nil, fmt.Errorf("could not open module '%s' from registry: %w", modulePath, err)
	}

	overlayFS := merged_fs.NewMergedFS(commonFS, moduleFS)

	cueInstance, err := kindsys.BuildInstance(themaRuntime.Context(), ".", "composable", overlayFS)
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
	return mapDir[kindsys.Composable](composableKindsPath, func(file os.DirEntry) (kindsys.Composable, error) {
		if !file.IsDir() {
			return nil, nil
		}

		return moduleToComposableKind(themaRuntime, commonFS, filepath.Join(composableKindsPath, file.Name()))
	})
}
