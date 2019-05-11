package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//root := "/home/andrzej.rehmann/projects"
	root := "/home/andrzej.rehmann/go/src"
	matchingPaths := make([]string, 0)
	scan(root, ".git", &matchingPaths)
	for i, e := range matchingPaths {
		fmt.Printf("%d %s\n", i, e)
	}
}

func scan(root string, needle string, matchingPaths *[]string) {
	haystack, err := ioutil.ReadDir(root)
	check(err)
	if containsNeedle(haystack, needle) {
		*matchingPaths = append(*matchingPaths, root)
		return
	}
	for _, file := range haystack {
		if file.IsDir() {
			dir := fmt.Sprintf("%s/%s", root, file.Name())
			scan(dir, needle, matchingPaths)
		}
	}
}

func check(err error) {
	if err != nil {
		panic(err)
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
