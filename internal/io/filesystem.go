package io

import (
	"github.com/hoto/fuzzy-repo-finder/internal/project"
	"strings"
)

type Filesystem struct {
	disk IDisk
}

func NewFilesystem(disk IDisk) Filesystem {
	return Filesystem{disk}
}

func (fs Filesystem) FindProjects(root string) []project.Project {
	gitDirs := fs.disk.FindDirs(root, ".git")
	var projects []project.Project
	for _, projectPath := range gitDirs {
		tokens := strings.Split(projectPath, "/")
		group := tokens[len(tokens)-1]
		projectName := tokens[len(tokens)-2]
		newProject := project.Project{FullPath: projectPath, Group: group, Name: projectName}
		projects = append(projects, newProject)
	}
	return projects
}
