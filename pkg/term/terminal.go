package term

import (
	"github.com/hoto/fuzzy-repo-finder/pkg/config"
	"github.com/hoto/fuzzy-repo-finder/pkg/proj"
	"github.com/nsf/termbox-go"
	"github.com/sahilm/fuzzy"
)

const (
	LettersNumbersSpecialCharacters = 0
)

type Terminal struct {
	display          *display
	projectNameField *field
	allProjects      proj.Projects
	filteredProjects proj.Projects
}

func NewTerminal(projects proj.Projects, query string) *Terminal {
	return &Terminal{
		display:          NewDisplay(),
		projectNameField: NewField("Name: ", query),
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
	selectedProject := t.filteredProjects.Get(0)
	t.filterProjects()
	t.display.positionCursor(t.projectNameField)
	t.display.displayField(t.projectNameField)
	t.display.displayProjects(&t.filteredProjects, &selectedProject)
	t.display.refresh()
	event := termbox.PollEvent()
	if event.Type == termbox.EventKey {
		switch event.Key {
		case LettersNumbersSpecialCharacters:
			t.projectNameField.appendToQuery(event.Ch)
			t.display.positionCursor(t.projectNameField)
		case termbox.KeyBackspace, termbox.KeyBackspace2:
			t.projectNameField.deleteLastQueryChar()
			t.display.positionCursor(t.projectNameField)
		case termbox.KeyCtrlW:
			t.projectNameField.eraseQuery()
			t.display.positionCursor(t.projectNameField)
		case termbox.KeyEnter:
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
