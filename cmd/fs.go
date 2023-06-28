package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"testing/fstest"
)

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
