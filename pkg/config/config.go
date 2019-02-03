package config

import (
	"github.com/hoto/fuzzy-repo-finder/pkg/proj"
	"os"
)

var (
	ProjectNameFilter string
	ProjectsRoots     []string

	configDir           = os.Getenv("HOME") + "/.fuzzy-repo-finder"
	configFile          = configDir + "/config.yml"
	selectedProjectFile = configDir + "/selected_project.txt"
)

func InitConfig() {
	argsConfig := newArgsConfig()
	ymlConfig := newYmlConfig()

	ProjectNameFilter = argsConfig.projectNameFilter
	ProjectsRoots = ymlConfig.ProjectRoots
}

func PersistSelectedProject(project proj.Project) {
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
