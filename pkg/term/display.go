package term

import (
	"github.com/hoto/fuzzy-repo-finder/pkg/proj"
	"github.com/nsf/termbox-go"
)

const (
	queryVerticalOffset         = 0
	projectsVerticalOffset      = 1
	projectNameHorizontalOffset = 4
)

type display struct {
	queryCursorPosition position
}

func NewDisplay() *display {
	return &display{queryCursorPosition: position{0, 0}}
}

func (display) displayQuery(queryPrompt string, query query) {
	for charHorizontalOffset, char := range queryPrompt {
		termbox.SetCell(
			charHorizontalOffset,
			queryVerticalOffset,
			char,
			termbox.ColorCyan,
			termbox.ColorDefault)
	}
	promptHorizontalOffset := len(queryPrompt)
	for charHorizontalOffset, char := range query.Runes() {
		termbox.SetCell(
			promptHorizontalOffset+charHorizontalOffset,
			queryVerticalOffset,
			char,
			termbox.ColorGreen,
			termbox.ColorDefault)
	}
}

func (display) displayProjects(projects *proj.Projects) {
	currentLineNum := projectsVerticalOffset
	for _, group := range projects.ListGroups() {
		for charOffset, char := range []rune(group) {
			termbox.SetCell(
				charOffset,
				currentLineNum,
				char,
				termbox.ColorMagenta,
				termbox.ColorDefault)
		}
		currentLineNum += 1
		for _, project := range projects.List() {
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

func (d *display) adjustQueryCursorPosition(queryPrompt string, query query) {
	d.queryCursorPosition.x = len(queryPrompt) + query.Size()
	termbox.SetCursor(d.queryCursorPosition.x, d.queryCursorPosition.y)
}

func (d *display) refresh() {
	d.flush()
	d.clear()
}

func (display) flush() {
	err := termbox.Flush()
	if err != nil {
		panic(err)
	}
}

func (display) clear() {
	err := termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	if err != nil {
		panic(err)
	}
}
