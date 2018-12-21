package main

import (
	"fmt"
	"github.com/hoto/fuzzy-repo-finder/internal/io"
	"github.com/hoto/fuzzy-repo-finder/internal/io/key"
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
		terminal.Clear()
		terminal.Flush()

		switch pressedKey := terminal.WaitForKeyPress(); pressedKey {
		case key.CtrlC:
			return 1
		case key.Backspace:
			fmt.Println("Backspace")
		case key.NonFunctional:
			fmt.Println("NonFunctional")
		}

	}
}
