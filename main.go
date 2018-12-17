package main

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	projectsDir = "/home/andrzej.rehmann/projects"
)

func main() {
	projects := findProjects(projectsDir)
	for _, project := range projects {
		fmt.Println(project)
	}
}

func findProjects(dir string) []string {
	var fileList []string
	err := filepath.Walk(dir,
		func(path string, file os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if file.IsDir() && file.Name() == ".git" {
				fileList = append(fileList, path)
			}
			return nil
		})
	if err != nil {
		fmt.Println(err)
	}
	return fileList
}
