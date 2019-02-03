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

func (Disk) FindDirs(root string, dirName string) []string {
	var matchingDirs []string
	appendMatchingDirs := func(path string, file os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if matchesDirName(file, dirName) {
			matchingDirs = append(matchingDirs, path)
			return filepath.SkipDir
		}
		return nil
	}
	err := filepath.Walk(root, appendMatchingDirs)
	if err != nil {
		log.Fatal(err)
	}
	return matchingDirs
}

func matchesDirName(file os.FileInfo, dirName string) bool {
	return file.IsDir() && file.Name() == dirName
}
