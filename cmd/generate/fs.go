package generate

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"testing/fstest"
)

func dirToPrefixedFS(directory string, prefix string) (fs.FS, error) {
	dirHandle, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	commonFS := fstest.MapFS{}
	for _, file := range dirHandle {
		if file.IsDir() {
			continue
		}

		content, err := os.ReadFile(filepath.Join(directory, file.Name()))
		if err != nil {
			return nil, err
		}

		commonFS[filepath.Join(prefix, file.Name())] = &fstest.MapFile{Data: content}
	}

	return commonFS, nil
}

func mapDir[T any](directory string, readFunc func(file os.DirEntry) (T, error)) ([]T, error) {
	files, err := os.ReadDir(directory)
	if err != nil {
		return nil, fmt.Errorf("could not open directory '%s': %w", directory, err)
	}

	results := make([]T, 0, len(files))
	for _, file := range files {
		result, err := readFunc(file)
		if err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}
