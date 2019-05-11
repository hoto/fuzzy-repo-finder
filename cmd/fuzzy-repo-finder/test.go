package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//dir := "/home/andrzej.rehmann/projects"
	root := "/home/andrzej.rehmann/go"
	scan(root, ".git")
}

func scan(root string, needle string) {
	haystack, err := ioutil.ReadDir(root)
	check(err)
	if containsNeedle(haystack, needle) {
		return
	}
	for _, file := range haystack {
		fmt.Printf("%s/%s\n", root, file.Name())
		if file.IsDir() {
			dir := fmt.Sprintf("%s/%s", root, file.Name())
			scan(dir, needle)
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
