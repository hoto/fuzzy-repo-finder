package finder

import (
	"fmt"
	"strings"
)

type IO struct {
	filesystem Filesystem
}

func (io IO) FindProjects(dir string) []Project {
	gitDirs := io.filesystem.FindGitDirectories(dir)

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

type Project struct {
	FullPath string
	Group    string
	Name     string
}

func (p Project) String() string {
	return fmt.Sprintf("Name=[%s], Group=[%s], FullPath=[%s]", p.Name, p.Group, p.FullPath)
}
