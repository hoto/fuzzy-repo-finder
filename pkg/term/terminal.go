package term

import (
	"github.com/hoto/fuzzy-repo-finder/pkg/config"
	"github.com/hoto/fuzzy-repo-finder/pkg/proj"
	"github.com/nsf/termbox-go"
	"github.com/sahilm/fuzzy"
)

const (
	LettersNumbersAndSpecialCharacters = 0
)

type Terminal struct {
	display              *display
	projectNameField     *field
	allProjects          proj.Projects
	filteredProjects     proj.Projects
	selectedProjectIndex int
}

func NewTerminal(projects proj.Projects, query string) *Terminal {
	return &Terminal{
		display:              NewDisplay(),
		projectNameField:     NewField("Name: ", query),
		allProjects:          projects,
		filteredProjects:     projects,
		selectedProjectIndex: 0,
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
	t.bindSelectedProject()
	t.display.positionCursor(t.projectNameField)
	t.display.displayField(t.projectNameField)
	t.display.displayProjects(&t.filteredProjects, t.selectedProjectIndex)
	t.display.refresh()
	event := termbox.PollEvent()
	if event.Type == termbox.EventKey {
		switch event.Key {
		case LettersNumbersAndSpecialCharacters:
			t.projectNameField.appendToQuery(event.Ch)
			t.display.positionCursor(t.projectNameField)
			t.resetSelectedProject()
		case termbox.KeyBackspace, termbox.KeyBackspace2:
			t.projectNameField.deleteLastQueryChar()
			t.display.positionCursor(t.projectNameField)
			t.resetSelectedProject()
		case termbox.KeyCtrlW:
			t.projectNameField.eraseQuery()
			t.display.positionCursor(t.projectNameField)
			t.resetSelectedProject()
		case termbox.KeyArrowUp:
			t.selectPreviousProject()
		case termbox.KeyArrowDown:
			t.selectNextProject()
		case termbox.KeyEnter:
			config.PersistSelectedProject(t.filteredProjects.Get(t.selectedProjectIndex))
			return NormalExit
		case termbox.KeyCtrlC:
			return AbnormalExit
		}
	}
	t.filterProjects()
	return ContinueRunning
}

func (t *Terminal) filterProjects() {
	if t.projectNameField.queryIsEmpty() {
		t.filteredProjects = t.allProjects
		return
	}
	matchedProjects := proj.NewProjects()
	matches := fuzzy.FindFrom(t.projectNameField.queryString(), t.allProjects)
	for _, match := range matches {
		matchedProjects.Add(t.allProjects.Get(match.Index))
	}
	t.filteredProjects = matchedProjects
}

func (t *Terminal) selectPreviousProject() {
	t.selectedProjectIndex -= 1
	t.bindSelectedProject()
}

func (t *Terminal) selectNextProject() {
	t.selectedProjectIndex += 1
	t.bindSelectedProject()
}

func (t *Terminal) bindSelectedProject() {
	projectsSize := t.filteredProjects.Size()
	if t.selectedProjectIndex < 0 {
		t.selectedProjectIndex = projectsSize - 1
	}
	if projectsSize < t.selectedProjectIndex+1 {
		t.selectedProjectIndex = 0
	}
}

func (t *Terminal) resetSelectedProject() {
	t.selectedProjectIndex = 0
}
