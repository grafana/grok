package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"testing/fstest"
)

type someFS struct {
	basePath string
}

func (f *someFS) Open(name string) (fs.File, error) {
	fmt.Printf("opening file %s\n", filepath.Join(f.basePath, name))
	return os.Open(filepath.Join(f.basePath, name))
}

func fileToFS(filePath string) (fs.FS, error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// core kinds are all dumped in the same folder, which isn't a valid cue module.
	// to work around that, we create a virtual FS to isolate each files into an in-memory module
	return fstest.MapFS{
		filepath.Base(filePath): &fstest.MapFile{Data: fileContent},
	}, nil
}
