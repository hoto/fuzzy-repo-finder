package main

import (
	"github.com/hoto/fuzzy-repo-finder/pkg/config"
	"github.com/hoto/fuzzy-repo-finder/pkg/io"
	"github.com/hoto/fuzzy-repo-finder/pkg/proj"
	"github.com/hoto/fuzzy-repo-finder/pkg/term"
	"os"
)

func main() {
	config.ParseArguments()

	filesystem := io.NewFilesystem(io.Disk{})
	allProjects := proj.NewProjects()
	for _, root := range config.Roots {
		projects := filesystem.FindGitProjects(root)
		allProjects.AddAll(projects.List())
	}

	os.Exit(run(allProjects, config.Query))
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
