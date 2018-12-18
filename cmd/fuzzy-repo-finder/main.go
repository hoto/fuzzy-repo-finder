package main

import (
	"fmt"
	"github.com/hoto/fuzzy-repo-finder/internal/finder"
)

const (
	projectsRoot = "/home/andrzej.rehmann/projects"
)

func main() {
	projects := finder.FindProjects(projectsRoot)
	for _, project := range projects {
		fmt.Println(project)
	}
}
