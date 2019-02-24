package main

import (
	"fmt"
	"github.com/hoto/fuzzy-repo-finder/pkg/config"
	"github.com/hoto/fuzzy-repo-finder/pkg/io"
	"github.com/hoto/fuzzy-repo-finder/pkg/proj"
	"github.com/hoto/fuzzy-repo-finder/pkg/term"
	. "github.com/logrusorgru/aurora"
	"os"
)

func main() {
	config.ParseArgsAndFlags()
	projects := readProjectsFromDisk()
	if config.Debug {
		debugLog(projects)
		os.Exit(0)
	}
	os.Exit(loop(projects))
}

func readProjectsFromDisk() proj.Projects {
	filesystem := io.NewFilesystem(io.Disk{})
	allProjects := proj.NewProjects()
	for _, root := range config.ProjectsRoots {
		projects := filesystem.FindGitProjects(root)
		allProjects.AddAll(projects.List())
	}
	return allProjects
}

func debugLog(projects proj.Projects) {
	fmt.Println()
	fmt.Println("Projects:")
	fmt.Printf("  projects=%s\n", Cyan(projects))
}

func loop(projects proj.Projects) int {
	terminal := term.NewTerminal(projects, config.ProjectNameFilter)
	terminal.Init()
	defer terminal.Close()

	for {
		rc := terminal.Cycle()
		switch rc {
		case term.ContinueRunning:
			continue
		case term.NormalExit:
			fmt.Print(config.SelectedProjectPath)
			return 0
		case term.AbnormalExit:
			return 1
		}
	}
}
