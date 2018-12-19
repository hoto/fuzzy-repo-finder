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
	fs := io.NewFilesystem(io.Disk{})
	projects := fs.FindGitProjects(projectsRoot)
	goProjects := fs.FindGitProjects(goProjectsRoot)
	for _, project := range projects {
		fmt.Println(project)
	}
	for _, project := range goProjects {
		fmt.Println(project)
	}
}
