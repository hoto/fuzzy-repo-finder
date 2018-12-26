package term

import (
	"github.com/hoto/fuzzy-repo-finder/pkg/proj"
	"github.com/nsf/termbox-go"
	"github.com/sahilm/fuzzy"
)

type Terminal struct {
	display          display
	queryPrompt      string
	query            Query
	allProjects      proj.Projects
	filteredProjects proj.Projects
	cursorPosition   position
}

func NewTerminal(projects proj.Projects) *Terminal {
	return &Terminal{
		queryPrompt:      "Name: ",
		allProjects:      projects,
		filteredProjects: projects,
		cursorPosition:   position{0, 0},
	}
}

// Initializes terminal. This function must be called before any other functions.
// Terminal must be closed using Close() function before exiting the application.
//
//      terminal := terminal.Init()
//      defer terminal.Close()
func (Terminal) Init() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
}

func (Terminal) Close() {
	termbox.Close()
}

func (t *Terminal) Cycle() ExitCode {
	t.positionCursor()
	t.display.displayQuery(t.queryPrompt, t.query)
	t.display.displayProjects(&t.filteredProjects)
	t.display.refresh()
	event := termbox.PollEvent()
	if event.Type == termbox.EventKey {
		switch event.Key {
		case 0, termbox.KeySpace:
			t.query.Append(event.Ch)
			t.positionCursor()
			t.filterProjects()
		case termbox.KeyBackspace, termbox.KeyBackspace2:
			t.query.DeleteLastChar()
			t.positionCursor()
			t.filterProjects()
		case termbox.KeyCtrlW:
			t.query.DeleteLastWord()
			t.positionCursor()
			t.filterProjects()
		case termbox.KeyEnter:
			return NORMAL_EXIT
		case termbox.KeyCtrlC:
			return ABNORMAL_EXIT
		}
	}
	return CONTINUE
}

func (t *Terminal) filterProjects() {
	if t.query.Size() == 0 {
		t.filteredProjects = t.allProjects
	}
	matchedProjects := proj.NewProjects()
	matches := fuzzy.FindFrom(t.query.String(), t.allProjects)
	for _, match := range matches {
		matchedProjects.Add(t.allProjects.List()[match.Index])
	}
	t.filteredProjects = matchedProjects
}
