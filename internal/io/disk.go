package io

import (
	"log"
	"os"
	"path/filepath"
)

type IDisk interface {
	FindDirs(root string, dirName string) []string
}

type Disk struct{}

func (Disk) FindDirs(dir string, matchDir string) []string {
	var matchingDirs []string
	appendMatchingDirs := func(path string, file os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if matchesDirName(file, matchDir) {
			matchingDirs = append(matchingDirs, path)
		}
		return nil
	}
	err := filepath.Walk(dir, appendMatchingDirs)
	if err != nil {
		log.Fatal(err)
	}
	return matchingDirs
}

func matchesDirName(file os.FileInfo, dirName string) bool {
	return file.IsDir() && file.Name() == dirName
}
