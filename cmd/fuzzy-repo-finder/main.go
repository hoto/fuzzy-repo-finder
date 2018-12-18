package main

import (
	"fmt"
	"github.com/hoto/fuzzy-repo-finder/internal/finder"
	"os"
)

var (
	projectsRoot = os.Getenv("HOME") + "/projects"
)

func main() {
	projects := finder.FindGitDirectories(projectsRoot)
	for _, project := range projects {
		fmt.Println(project)
	}
}
