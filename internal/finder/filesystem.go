package finder

import (
	"log"
	"os"
	"path/filepath"
)

type Filesystem interface {
	FindGitDirectories(dir string) []string
}

func FindGitDirectories(dir string) []string {
	var gitDirs []string
	err := filepath.Walk(dir,
		func(path string, file os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if isGitDirectory(file) {
				gitDirs = append(gitDirs, path)
			}
			return nil
		})
	if err != nil {
		//TODO: should I throw error to the caller or crash the app?
		log.Fatal(err)
	}
	return gitDirs
}

func isGitDirectory(file os.FileInfo) bool {
	return file.IsDir() && file.Name() == ".git"
}
