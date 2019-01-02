package main

import (
	"github.com/hoto/fuzzy-repo-finder/pkg/config"
	"github.com/hoto/fuzzy-repo-finder/pkg/io"
	"github.com/hoto/fuzzy-repo-finder/pkg/proj"
	"github.com/hoto/fuzzy-repo-finder/pkg/term"
	"os"
)

func main() {
	config.InitConfig()
	projects := readProjectsFromDisk()
	os.Exit(loop(projects))
}

func readProjectsFromDisk() proj.Projects {
	filesystem := io.NewFilesystem(io.Disk{})
	allProjects := proj.NewProjects()
	for _, root := range config.Roots {
		projects := filesystem.FindGitProjects(root)
		allProjects.AddAll(projects.List())
	}
	return allProjects
}

func loop(projects proj.Projects) int {
	terminal := term.NewTerminal(projects, config.Query)
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
