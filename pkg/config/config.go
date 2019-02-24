package config

import (
	"flag"
	"fmt"
	. "github.com/logrusorgru/aurora"
	"strings"
)

var (
	Debug               bool
	ProjectsRoots       []string
	ProjectNameFilter   string
	SelectedProjectPath = ""
)

func ParseArgsAndFlags() {

	flag.BoolVar(&Debug, "debug", false, "Show verbose debug information")
	//version := flag.Bool("version", false, "Show version")
	projectRoots := flag.String("projectRoots", "/default-projects-path", "Comma separated list of project roots directories")
	flag.Parse()

	ProjectNameFilter = strings.Join(flag.Args(), "")
	ProjectsRoots = strings.Split(*projectRoots, ",")

	if Debug {
		fmt.Println("Flags:")
		fmt.Printf("  projectRoots=%s\n", Cyan(*projectRoots))

		fmt.Println()
		fmt.Println("Args:")
		fmt.Printf("  args=%s\n", Cyan(flag.Args()))

		fmt.Println()
		fmt.Println("Config:")
		fmt.Printf("  ProjectRoots=%s\n", Cyan(ProjectsRoots))
		fmt.Printf("  ProjectNameFilter=%s\n", Cyan(ProjectNameFilter))
	}
}
