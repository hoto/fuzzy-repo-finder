package term

import "github.com/hoto/fuzzy-repo-finder/pkg/proj"

type projectSelector struct {
	projects             proj.Projects
	selectedProjectIndex int
}

func NewProjectSelector(projects proj.Projects) *projectSelector {
	return &projectSelector{
		projects:             projects,
		selectedProjectIndex: 0,
	}
}

func (p *projectSelector) selectPreviousProject() {
	p.selectedProjectIndex -= 1
	p.bindIndexInRange()
}

func (p *projectSelector) selectNextProject() {
	p.selectedProjectIndex += 1
	p.bindIndexInRange()
}

func (p *projectSelector) bindIndexInRange() {
	projectsSize := p.projects.Size()
	if p.selectedProjectIndex < 0 {
		p.selectedProjectIndex = projectsSize - 1
	}
	if projectsSize < p.selectedProjectIndex+1 {
		p.selectedProjectIndex = 0
	}
}

func (p *projectSelector) resetSelectedProject() {
	p.selectedProjectIndex = 0
}

func (p *projectSelector) index() int {
	return p.selectedProjectIndex
}

func (p *projectSelector) setProjects(projects proj.Projects) {
	p.projects = projects
}
