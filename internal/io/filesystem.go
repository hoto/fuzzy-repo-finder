package io

import (
	"fmt"
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
	fmt.Println("separator=", separator)
	gitDirs := fs.disk.FindDirs(root, ".git")
	var projects []project.Project
	for _, path := range gitDirs {
		tokens := strings.Split(path, separator)
		fullPath := strings.Join(tokens[0:len(tokens)-1], separator)
		group := ""
		name := tokens[len(tokens)-2]
		newProject := project.Project{FullPath: fullPath, Group: group, Name: name}
		projects = append(projects, newProject)
	}
	return projects
}
