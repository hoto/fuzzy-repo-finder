package main

import (
	"github.com/hoto/fuzzy-repo-finder/pkg/config"
	"github.com/hoto/fuzzy-repo-finder/pkg/io"
	"github.com/hoto/fuzzy-repo-finder/pkg/proj"
	"github.com/hoto/fuzzy-repo-finder/pkg/term"
	"os"
)

var (
	projectsRoot   = os.Getenv("HOME") + "/projects"
	goProjectsRoot = os.Getenv("HOME") + "/go/src"
)

func main() {
	query := config.ParseArguments()

	filesystem := io.NewFilesystem(io.Disk{})
	projects := filesystem.FindGitProjects(projectsRoot)
	goProjects := filesystem.FindGitProjects(goProjectsRoot)
	allProjects := proj.NewProjects()
	allProjects.AddAll(projects.List())
	allProjects.AddAll(goProjects.List())

	os.Exit(run(allProjects, query))
}

func run(projects proj.Projects, query string) int {
	terminal := term.NewTerminal(projects, query)
	terminal.Init()
	defer terminal.Close()

	for {
		rc := terminal.Cycle()
		switch rc {
		case term.ContinueRunning:
			continue
		case term.NormalExit:
			return 0
		case term.AbnormalExit:
			return 1
		}
	}
}
