package io

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type IDisk interface {
	FindDirs(root string, dirName string) []string
}

type Disk struct{}

func (Disk) FindDirs(root string, needle string) []string {
	matchingPaths := make([]string, 0)
	scan(root, needle, &matchingPaths)
	return matchingPaths
}

func scan(dir string, needle string, matchingPaths *[]string) {
	haystack, err := ioutil.ReadDir(dir)
	check(err)
	if containsNeedle(haystack, needle) {
		gitPath := dir + "/.git" // TODO return just the dir without the .git
		*matchingPaths = append(*matchingPaths, gitPath)
		return
	}
	for _, file := range haystack {
		if file.IsDir() {
			dir := fmt.Sprintf("%s/%s", dir, file.Name())
			scan(dir, needle, matchingPaths)
		}
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func containsNeedle(files []os.FileInfo, needle string) bool {
	for _, file := range files {
		if file.Name() == needle {
			return true
		}
	}
	return false
}
