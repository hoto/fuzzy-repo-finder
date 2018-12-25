package project

type Projects struct {
	projects []Project
}

func NewProjects() Projects {
	return Projects{projects: make([]Project, 0)}
}

func (p *Projects) List() []Project {
	return p.projects
}
