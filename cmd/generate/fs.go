package generate

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"testing/fstest"
)

type FSOption func(file *fstest.MapFile) *fstest.MapFile

func PrependFilesWith(input string) FSOption {
	return func(file *fstest.MapFile) *fstest.MapFile {
		return &fstest.MapFile{
			Data: []byte(fmt.Sprintf("%s%s", input, file.Data)),
		}
	}
}

func fileToFS(filePath string, options ...FSOption) (fs.FS, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	mapFile := &fstest.MapFile{Data: content}
	for _, opt := range options {
		mapFile = opt(mapFile)
	}

	return fstest.MapFS{
		filepath.Base(filePath): mapFile,
	}, nil
}

func dirToPrefixedFS(directory string, prefix string, options ...FSOption) (fs.FS, error) {
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

		mapFile := &fstest.MapFile{Data: content}
		for _, opt := range options {
			mapFile = opt(mapFile)
		}

		commonFS[filepath.Join(prefix, file.Name())] = mapFile
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
