package finder

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func FindProjects(dir string) []Project {
	gitDirs := findGitDirectories(dir)

	var projects []Project
	for _, projectPath := range gitDirs {
		tokens := strings.Split(projectPath, "/")
		group := tokens[len(tokens)-1]
		projectName := tokens[len(tokens)-2]
		project := Project{projectPath, group, projectName}
		projects = append(projects, project)
	}
	return projects
}

func findGitDirectories(dir string) []string {
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
		fmt.Println(err)
	}
	return gitDirs
}

func isGitDirectory(file os.FileInfo) bool {
	return file.IsDir() && file.Name() == ".git"
}

type Project struct {
	FullPath string
	Group    string
	Name     string
}

func (p Project) String() string {
	return fmt.Sprintf("Name=[%s], Group=[%s], FullPath=[%s]", p.Name, p.Group, p.FullPath)
}
