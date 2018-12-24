package main

import (
	"github.com/hoto/fuzzy-repo-finder/internal/io"
	"github.com/hoto/fuzzy-repo-finder/internal/project"
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
	allProjects := append(projects, goProjects...)
	os.Exit(runTerminal(allProjects))
}

func runTerminal(projects []project.Project) int {
	terminal := io.NewTerminal(projects)
	terminal.Init()
	defer terminal.Close()

	for {
		rc := terminal.Cycle()
		if rc != 0 {
			return rc
		}
	}
}
