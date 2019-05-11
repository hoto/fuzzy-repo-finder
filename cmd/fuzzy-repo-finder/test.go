package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	level := 0
	//dir := "/home/andrzej.rehmann/projects"
	dir := "/home/andrzej.rehmann/go"
	scan(level, dir, "")
}

func scan(level int, dir string, folderName string) {
	files, err := ioutil.ReadDir(dir)
	check(err)
	if isAGitRepo(files) {
		return
	}
	for i, file := range files {
		fmt.Printf("%d-%d %s/%s\n", level, i, dir, file.Name())
		if file.IsDir() {
			nextLevel := level + 1
			scan(nextLevel, fmt.Sprintf("%s/%s", dir, file.Name()), file.Name())
		}
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func isAGitRepo(files []os.FileInfo) bool {
	for _, file := range files {
		if file.Name() == ".git" {
			return true
		}
	}
	return false
}
