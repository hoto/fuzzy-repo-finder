package config

import (
	"github.com/hoto/fuzzy-repo-finder/pkg/proj"
	"os"
	"strings"
)

var (
	Query string

	home = os.Getenv("HOME")

	projectsRoot   = home + "/projects"
	goProjectsRoot = home + "/go/src"
	Roots          = []string{projectsRoot, goProjectsRoot}

	configDir           = home + "/.fuzzy-repo-finder"
	selectedProjectFile = configDir + "/selected_project.txt"
)

func ParseArguments() {
	args := os.Args[1:]
	Query = strings.Join(args, "")
}

func PersistSelectedProject(project proj.Project) {
	createConfigDir()
	file, err := os.Create(selectedProjectFile)
	check(err)
	defer file.Close()
	_, err = file.WriteString(project.FullPath)
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func createConfigDir() {
	err := os.MkdirAll(configDir, 0755)
	if err != nil {
		panic(err)
	}
}
