package proj

import (
	"errors"
	"sort"
)

type Projects struct {
	projects        []Project
	selectedProject *Project
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

func (p *Projects) AddAll(projects []Project) {
	for _, project := range projects {
		p.Add(project)
	}
}

func (p *Projects) ListGroups() []string {
	groupsSet := make(map[string]bool)
	for _, project := range p.projects {
		groupsSet[project.Group] = true
	}
	groups := make([]string, 0)
	for k := range groupsSet {
		groups = append(groups, string(k))
	}
	sort.Strings(groups)
	return groups
}

func (p *Projects) Size() int {
	return len(p.projects)
}

func (p Projects) Len() int {
	return p.Size()
}

func (p Projects) String(i int) string {
	return p.projects[i].Name
}

func (p Projects) Copy() Projects {
	return p
}

func (p *Projects) Get(i int) Project {
	return p.projects[i]
}

func (p *Projects) GetSelected() (*Project, error) {
	if p.selectedProject == nil {
		return nil, errors.New("no project is selected")
	}
	return p.selectedProject, nil
}

func (p *Projects) MarkSelected(selectedProject Project) {
	for i, project := range p.projects {
		if project.FullPath == selectedProject.FullPath {
			p.selectedProject = &p.projects[i]
			break
		}
	}
}
