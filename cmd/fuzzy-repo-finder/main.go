package main

import (
	"fmt"
	"github.com/hoto/fuzzy-repo-finder/internal/io"
	"os"
)

var (
	projectsRoot = os.Getenv("HOME") + "/projects"
)

func main() {
	fs := io.NewFilesystem(io.Disk{})
	projects := fs.FindGitProjects(projectsRoot)
	for _, project := range projects {
		fmt.Println(project)
	}
}
