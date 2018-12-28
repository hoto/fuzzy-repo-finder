package config

import (
	"github.com/hoto/fuzzy-repo-finder/pkg/proj"
	"os"
)

var (
	home                = os.Getenv("HOME")
	configDir           = home + "/.fuzzy-repo-finder"
	selectedProjectFile = configDir + "/selected_project.txt"
)

func PersistSelectedProject(project proj.Project) {
	createConfigDir()
	file, err := os.Create(selectedProjectFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString(project.FullPath)
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
