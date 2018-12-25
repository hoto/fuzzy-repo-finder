package io

import (
	"github.com/hoto/fuzzy-repo-finder/internal/proj"
	"github.com/nsf/termbox-go"
)

const (
	queryVerticalOffset         = 0
	projectsVerticalOffset      = 1
	projectNameHorizontalOffset = 4
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
	for charHorizontalOffset, char := range t.queryPrompt {
		termbox.SetCell(
			charHorizontalOffset,
			queryVerticalOffset,
			char,
			termbox.ColorMagenta,
			termbox.ColorDefault)
	}
	promptHorizontalOffset := len(t.queryPrompt)
	for charHorizontalOffset, char := range t.query.Read() {
		termbox.SetCell(
			promptHorizontalOffset+charHorizontalOffset,
			queryVerticalOffset,
			char,
			termbox.ColorGreen,
			termbox.ColorDefault)
	}
}

func (t *Terminal) displayProjects() {
	currentLineNum := projectsVerticalOffset
	for _, group := range t.projects.ListGroups() {
		for charOffset, char := range []rune(group) {
			termbox.SetCell(
				charOffset,
				currentLineNum,
				char,
				termbox.ColorYellow,
				termbox.ColorDefault)
		}
		currentLineNum += 1
		for _, project := range t.projects.List() {
			if project.Group == group {
				for charOffset, char := range []rune(project.Name) {
					termbox.SetCell(
						projectNameHorizontalOffset+charOffset,
						currentLineNum,
						char,
						termbox.ColorDefault,
						termbox.ColorDefault)
				}
				currentLineNum += 1
			}
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
