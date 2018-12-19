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
		group := parseGroup(root, fullPath)
		name := tokens[len(tokens)-2]
		projects = append(projects, project.Project{FullPath: fullPath, Group: group, Name: name})
	}
	return projects
}

// tdd made me do it
func parseGroup(root string, fullPath string) string {
	numberOfRootPathTokens := len(strings.Split(root, separator))
	fullPathSplitByRootPathTokens := strings.SplitAfterN(fullPath, separator, numberOfRootPathTokens+1)
	pathWithoutRoot := fullPathSplitByRootPathTokens[2]
	pathWithoutRootTokens := strings.Split(pathWithoutRoot, separator)
	groupTokens := pathWithoutRootTokens[0 : len(pathWithoutRootTokens)-1]
	group := strings.Join(groupTokens, separator)
	return group
}
