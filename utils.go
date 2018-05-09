package main

import (
	"os"
	"path/filepath"
)

func subDirectoryFiles(rootPath string, actionFunc func(os.FileInfo) error) error {
	return filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err := actionFunc(info); err != nil {
			return err
		}
		return nil
	})
}
