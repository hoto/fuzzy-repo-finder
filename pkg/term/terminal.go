package term

import (
	"github.com/hoto/fuzzy-repo-finder/pkg/config"
	"github.com/hoto/fuzzy-repo-finder/pkg/proj"
	"github.com/nsf/termbox-go"
)

const (
	LettersNumbersAndSpecialCharacters = 0
)

type Terminal struct {
	display          *display
	projectNameField *field
	allProjects      proj.Projects
	filteredProjects proj.Projects
	projectSelector  *projectSelector
}

func NewTerminal(projects proj.Projects, query string) *Terminal {
	return &Terminal{
		display:          NewDisplay(),
		projectNameField: NewField("Search: ", query),
		allProjects:      projects,
		filteredProjects: projects,
		projectSelector:  NewProjectSelector(projects),
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
	t.filterProjects()
	t.projectSelector.setProjects(t.filteredProjects)
	t.projectSelector.bindSelectedProject()
	t.display.positionCursor(t.projectNameField)
	t.display.displayField(t.projectNameField)
	t.display.displayProjects(&t.filteredProjects, t.projectSelector.index())
	t.display.refresh()
	event := termbox.PollEvent()
	if event.Type == termbox.EventKey {
		switch event.Key {
		case LettersNumbersAndSpecialCharacters:
			t.projectNameField.appendToQuery(event.Ch)
			t.display.positionCursor(t.projectNameField)
			t.projectSelector.resetSelectedProject()
		case termbox.KeyBackspace, termbox.KeyBackspace2:
			t.projectNameField.deleteLastQueryChar()
			t.display.positionCursor(t.projectNameField)
			t.projectSelector.resetSelectedProject()
		case termbox.KeyCtrlW:
			t.projectNameField.eraseQuery()
			t.display.positionCursor(t.projectNameField)
			t.projectSelector.resetSelectedProject()
		case termbox.KeyArrowUp:
			t.projectSelector.selectPreviousProject()
		case termbox.KeyArrowDown:
			t.projectSelector.selectNextProject()
		case termbox.KeyEnter:
			selectedProject := t.filteredProjects.Get(t.projectSelector.selectedProjectIndex)
			config.PersistSelectedProject(selectedProject)
			return NormalExit
		case termbox.KeyCtrlC:
			return AbnormalExit
		}
	}
	t.filterProjects()
	return ContinueRunning
}

func (t *Terminal) filterProjects() {
	t.filteredProjects = proj.FuzzyMatch(t.projectNameField.queryString(), t.allProjects)
}
