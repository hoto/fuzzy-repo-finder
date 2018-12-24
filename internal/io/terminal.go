package io

import (
	"github.com/hoto/fuzzy-repo-finder/internal/project"
	"github.com/nsf/termbox-go"
)

type Terminal struct {
	queryPrompt    string
	query          Query
	projects       []project.Project
	cursorPosition position
}

func NewTerminal(projects []project.Project) *Terminal {
	return &Terminal{
		queryPrompt:    "Search: ",
		projects:       projects,
		cursorPosition: position{0, 0},
	}
}

// Initializes terminal. This function must be called before any other functions.
// Terminal must be closed using Close() function before exiting the application.
//
//      terminal := terminal.Init()
//      defer terminal.Close()
func (t Terminal) Init() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
}

func (t Terminal) Close() {
	termbox.Close()
}

func (t *Terminal) Cycle() int {
	t.displayCursor()
	t.displayQuery()
	t.displayProjects()
	t.refresh()
	event := termbox.PollEvent()
	if event.Type == termbox.EventKey {
		switch event.Key {
		case 0, termbox.KeySpace:
			t.query.Append(event.Ch)
			t.moveCursor()
		case termbox.KeyBackspace, termbox.KeyBackspace2:
			t.query.DeleteLastChar()
			t.moveCursor()
		case termbox.KeyCtrlW:
			t.query.DeleteLastWord()
			t.moveCursor()
		case termbox.KeyCtrlC:
			return 1
		}
	}
	return 0
}

func (t *Terminal) displayCursor() {
	termbox.SetCursor(t.cursorPosition.x, t.cursorPosition.y)
}

func (t *Terminal) displayQuery() {
	for i, char := range t.query.Read() {
		termbox.SetCell(i, 0, char, termbox.ColorGreen, termbox.ColorDefault)
	}
}

func (t *Terminal) displayProjects() {
	offset := 1
	for projectIndex, _project := range t.projects {
		for runeIndex, char := range []rune(_project.Name) {
			termbox.SetCell(runeIndex, projectIndex+offset, char, termbox.ColorGreen, termbox.ColorDefault)
		}
	}
}

func (t *Terminal) moveCursor() {
	t.cursorPosition.x = t.query.Size()
	termbox.SetCursor(t.cursorPosition.x, t.cursorPosition.y)
}

func (t *Terminal) refresh() {
	t.flush()
	t.clear()
}

func (t *Terminal) flush() {
	err := termbox.Flush()
	if err != nil {
		panic(err)
	}
}

func (t *Terminal) clear() {
	err := termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	if err != nil {
		panic(err)
	}
}
