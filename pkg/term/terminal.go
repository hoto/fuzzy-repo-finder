package term

import (
	"github.com/hoto/fuzzy-repo-finder/pkg/config"
	"github.com/hoto/fuzzy-repo-finder/pkg/proj"
	"github.com/nsf/termbox-go"
	"github.com/sahilm/fuzzy"
)

type Terminal struct {
	display          *display
	projectNameField *field
	allProjects      proj.Projects
	filteredProjects proj.Projects
}

func NewTerminal(projects proj.Projects) *Terminal {
	return &Terminal{
		display:          NewDisplay(),
		projectNameField: NewField("Name: ", ""),
		allProjects:      projects,
		filteredProjects: projects,
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
	t.display.adjustQueryCursorPosition(t.projectNameField)
	t.display.displayQuery(t.projectNameField)
	t.display.displayProjects(&t.filteredProjects)
	t.display.refresh()
	event := termbox.PollEvent()
	if event.Type == termbox.EventKey {
		switch event.Key {
		case 0, termbox.KeySpace:
			t.projectNameField.appendToQuery(event.Ch)
			t.display.adjustQueryCursorPosition(t.projectNameField)
		case termbox.KeyBackspace, termbox.KeyBackspace2:
			t.projectNameField.deleteLastQueryChar()
			t.display.adjustQueryCursorPosition(t.projectNameField)
		case termbox.KeyCtrlW:
			t.projectNameField.eraseQuery()
			t.display.adjustQueryCursorPosition(t.projectNameField)
		case termbox.KeyEnter:
			selectedProject := t.filteredProjects.Get(0)
			config.PersistSelectedProject(selectedProject)
			return NORMAL_EXIT
		case termbox.KeyCtrlC:
			return ABNORMAL_EXIT
		}
	}
	t.filterProjects()
	return CONTINUE
}

func (t *Terminal) filterProjects() {
	if t.projectNameField.queryIsEmpty() {
		t.filteredProjects = t.allProjects
		return
	}
	matchedProjects := proj.NewProjects()
	matches := fuzzy.FindFrom(t.projectNameField.queryString(), t.allProjects)
	for _, match := range matches {
		matchedProjects.Add(t.allProjects.List()[match.Index])
	}
	t.filteredProjects = matchedProjects
}
