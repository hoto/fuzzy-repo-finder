package main

import (
	"fmt"
	"github.com/hoto/fuzzy-repo-finder/internal/io"
	"os"
)

var (
	projectsRoot   = os.Getenv("HOME") + "/projects"
	goProjectsRoot = os.Getenv("HOME") + "/go/src"
)

func main() {
	filesystem := io.NewFilesystem(io.Disk{})
	projects := filesystem.FindGitProjects(projectsRoot)
	goProjects := filesystem.FindGitProjects(goProjectsRoot)
	for _, project := range projects {
		fmt.Println(project)
	}
	for _, project := range goProjects {
		fmt.Println(project)
	}
}
