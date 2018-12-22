package main

import (
	"github.com/hoto/fuzzy-repo-finder/internal/io"
	"os"
)

func main() {
	os.Exit(run())
}

func run() int {
	terminal := io.NewTerminal()
	terminal.Init()
	defer terminal.Close()

	for {
		rc := terminal.Cycle()
		if rc != 0 {
			return rc
		}
	}
}
