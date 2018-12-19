package io

import (
	"strings"
)

type Filesystem struct {
	disk IDisk
}

func NewFilesystem(disk IDisk) Filesystem {
	return Filesystem{disk}
}

func (fs Filesystem) FindProjects(root string) []Project {
	gitDirs := fs.disk.FindDirs(root, ".git")

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
