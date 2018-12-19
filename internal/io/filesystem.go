package io

import (
	"github.com/hoto/fuzzy-repo-finder/internal/project"
	"os"
	"strings"
)

var (
	separator = string(os.PathSeparator)
)

type Filesystem struct {
	disk IDisk
}

func NewFilesystem(disk IDisk) Filesystem {
	return Filesystem{disk}
}

func (fs Filesystem) FindGitProjects(root string) []project.Project {
	gitDirs := fs.disk.FindDirs(root, ".git")
	var projects []project.Project
	for _, path := range gitDirs {
		tokens := strings.Split(path, separator)
		fullPath := strings.Join(tokens[0:len(tokens)-1], separator)
		group := diffPath(root, fullPath)
		name := tokens[len(tokens)-2]
		projects = append(projects, project.Project{FullPath: fullPath, Group: group, Name: name})
	}
	return projects
}

func diffPath(root string, fullPath string) string {
	fullPathTokens := strings.Split(fullPath, separator)
	pathToProject := strings.Join(fullPathTokens[0:len(fullPathTokens)-1], separator)
	if len(root) == len(pathToProject) {
		return ""
	}
	return pathToProject[len(root)+1:]
}
