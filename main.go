package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const (
	projectsDir = "/home/andrzej.rehmann/projects"
)

func main() {
	err := filepath.Walk(projectsDir,
		func(path string, file os.FileInfo, err error) error {
			if file.IsDir() && file.Name() == ".git" {
				if err != nil {
					return err
				}
				fmt.Println(path)
			}
			return nil
		})

	if err != nil {
		log.Println(err)
	}
}
