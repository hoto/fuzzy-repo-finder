package proj

import (
	"sort"
)

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
