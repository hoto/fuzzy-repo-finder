package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//dir := "/home/andrzej.rehmann/projects"
	dir := "/home/andrzej.rehmann/go"
	scan(dir)
}

func scan(dir string) {
	files, err := ioutil.ReadDir(dir)
	check(err)
	if isAGitRepo(files) {
		return
	}
	for _, file := range files {
		fmt.Printf("%s/%s\n", dir, file.Name())
		if file.IsDir() {
			scan(fmt.Sprintf("%s/%s", dir, file.Name()))
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
