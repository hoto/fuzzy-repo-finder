package io

import (
	"github.com/hoto/fuzzy-repo-finder/internal/proj"
	"github.com/nsf/termbox-go"
)

const (
	queryLineVerticalOffset    = 0
	projectsLineVerticalOffset = 1
)

type Terminal struct {
	queryPrompt    string
	query          Query
	projects       proj.Projects
	cursorPosition position
}

func NewTerminal(projects proj.Projects) *Terminal {
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

func (t *Terminal) Cycle() ExitCode {
	t.positionCursor()
	t.displayQuery()
	t.displayProjects()
	t.refresh()
	event := termbox.PollEvent()
	if event.Type == termbox.EventKey {
		switch event.Key {
		case 0, termbox.KeySpace:
			t.query.Append(event.Ch)
			t.positionCursor()
		case termbox.KeyBackspace, termbox.KeyBackspace2:
			t.query.DeleteLastChar()
			t.positionCursor()
		case termbox.KeyCtrlW:
			t.query.DeleteLastWord()
			t.positionCursor()
		case termbox.KeyEnter:
			return NORMAL_EXIT
		case termbox.KeyCtrlC:
			return ABNORMAL_EXIT
		}
	}
	return CONTINUE
}

func (t *Terminal) displayQuery() {
	for i, char := range t.queryPrompt {
		termbox.SetCell(i, queryLineVerticalOffset, char, termbox.ColorMagenta, termbox.ColorDefault)
	}
	horizontalOffset := len(t.queryPrompt)
	for i, char := range t.query.Read() {
		termbox.SetCell(i+horizontalOffset, queryLineVerticalOffset, char, termbox.ColorGreen, termbox.ColorDefault)
	}
}

func (t *Terminal) displayProjects() {
	for projectIndex, project := range t.projects.List() {
		for charIndex, char := range []rune(project.Name) {
			termbox.SetCell(
				charIndex, projectIndex+projectsLineVerticalOffset, char,
				termbox.ColorDefault, termbox.ColorDefault)
		}
	}
}

func (t *Terminal) positionCursor() {
	t.cursorPosition.x = len(t.queryPrompt) + t.query.Size()
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
