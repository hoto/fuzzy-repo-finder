package proj

import "fmt"

type Projects struct {
	projects []Project
}

func NewProjects() Projects {
	return Projects{projects: make([]Project, 0)}
}

func (p *Projects) List() []Project {
	return p.projects
}

func (p *Projects) Add(project Project) {
	p.projects = append(p.projects, project)
}

func (p *Projects) String() string {
	return fmt.Sprintf("projects=[%s]", p.projects)
}
