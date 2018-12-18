package main

import (
	"fmt"
	"github.com/hoto/fuzzy-repo-finder/internal/finder"
)

const (
	projectsDir = "/home/andrzej.rehmann/projects"
)

func main() {
	projects := finder.FindProjects(projectsDir)
	for _, project := range projects {
		fmt.Println(project)
	}
}

