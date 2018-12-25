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
	os.Exit(run(allProjects))
}

func run(projects []project.Project) int {
	terminal := io.NewTerminal(projects)
	terminal.Init()
	defer terminal.Close()

	for {
		rc := terminal.Cycle()
		switch rc {
		case io.CONTINUE:
			continue
		case io.NORMAL_EXIT:
			return 0
		case io.ABNORMAL_EXIT:
			return 1
		}
	}
}
