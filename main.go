package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const (
	//projectsDir = "/home/andrzej.rehmann/projects"
	projectsDir = "/root"
)

func main() {
	err := filepath.Walk(projectsDir, findGitProjects)

	if err != nil {
		log.Println(err)
	}
}

func findGitProjects(path string, file os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if file.IsDir() && file.Name() == ".git" {
		fmt.Println(path)
	}
	return nil
}
