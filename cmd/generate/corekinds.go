package generate

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/grafana/kindsys"
	"github.com/grafana/thema"
)

func moduleToCoreKind(themaRuntime *thema.Runtime, modulePath string) (kindsys.Core, error) {
	fmt.Printf(" → Loading %s\n", modulePath)

	overlayFS, err := dirToPrefixedFS(modulePath, "")
	if err != nil {
		return nil, err
	}

	cueInstance, err := kindsys.BuildInstance(themaRuntime.Context(), ".", "kind", overlayFS)
	if err != nil {
		return nil, fmt.Errorf("could not load kindsys instance: %w", err)
	}

	props, err := kindsys.ToKindProps[kindsys.CoreProperties](cueInstance)
	if err != nil {
		return nil, fmt.Errorf("could not convert cue value to kindsys props: %w", err)
	}

	kindDefinition := kindsys.Def[kindsys.CoreProperties]{
		V:          cueInstance,
		Properties: props,
	}

	boundKind, err := kindsys.BindCore(themaRuntime, kindDefinition)
	if err != nil {
		return nil, fmt.Errorf("could not bind kind definition to kind: %w", err)
	}

	return boundKind, nil
}

func loadCoreKinds(themaRuntime *thema.Runtime, coreKindsPath string) ([]kindsys.Kind, error) {
	return mapDir[kindsys.Kind](coreKindsPath, func(file os.DirEntry) (kindsys.Kind, error) {
		if !file.IsDir() {
			return nil, nil
		}

		return moduleToCoreKind(themaRuntime, filepath.Join(coreKindsPath, file.Name()))
	})
}
