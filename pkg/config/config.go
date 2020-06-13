package config

import (
	"flag"
	"fmt"
	"github.com/logrusorgru/aurora"
	"os"
	"strings"
)

var (
	Version             string
	ShortCommit         string
	BuildDate           string
	Debug               bool
	ProjectsRoots       []string
	ProjectNameFilter   string
	SelectedProjectPath = ""
)

func ParseArgsAndFlags() {
	flag.Usage = overrideUsage()

	flag.BoolVar(&Debug, "debug", false, "Show verbose debug information")
	showVersion := flag.Bool("version", false, "Show version")
	projectRoots := flag.String("projectRoots", "",
		"Comma separated list of project roots directories")

	flag.Parse()

	ProjectNameFilter = strings.Join(flag.Args(), "")
	ProjectsRoots = strings.Split(*projectRoots, ",")

	if *showVersion {
		fmt.Printf("fuzzy-repo-finder version %s, commit %s, build %s\n",
			Version, ShortCommit, BuildDate)
		os.Exit(0)
	}

	if *projectRoots == "" {
		flag.Usage()
		os.Exit(1)
	}

	if Debug {
		debugLog(projectRoots)
	}
}

func overrideUsage() func() {
	return func() {
		_, _ = fmt.Fprintf(
			os.Stdout,
			"Usage:"+
				"\n\t"+
				"cd $(fuzzy-repo-finder --projectRoots=\"${HOME}/projects\" [flags] [QUERY])"+
				"\n\n"+
				"Flags:"+
				"\n")
		flag.PrintDefaults()
	}
}

func debugLog(projectRoots *string) {
	fmt.Println("Flags:")
	fmt.Printf("  projectRoots=%s\n", aurora.Cyan(*projectRoots))
	fmt.Println()
	fmt.Println("Args:")
	fmt.Printf("  args=%s\n", aurora.Cyan(flag.Args()))
	fmt.Println()
	fmt.Println("Config:")
	fmt.Printf("  ProjectRoots=%s\n", aurora.Cyan(ProjectsRoots))
	fmt.Printf("  ProjectNameFilter=%s\n", aurora.Cyan(ProjectNameFilter))
}
