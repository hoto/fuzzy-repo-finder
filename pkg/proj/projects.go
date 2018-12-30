package proj

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
	groups := []string{}
	for _, project := range p.projects {
		groupPresent := false
		if len(groups) == 0 {
			groups = append(groups, project.Group)
			continue
		}
		for _, group := range groups {
			if group == project.Group {
				groupPresent = true
				break
			}
		}
		if !groupPresent {
			groups = append(groups, project.Group)
		}
	}
	return groups
}

func (p *Projects) Size() int {
	return len(p.projects)
}

// implements fuzzy.Source
func (p Projects) Len() int {
	return p.Size()
}

// implements fuzzy.Source
func (p Projects) String(i int) string {
	return p.projects[i].Name
}

func (p Projects) Copy() Projects {
	return p
}

func (p *Projects) Get(i int) Project {
	if p.Size() == 0 {
		return Project{}
	}
	return p.projects[i]
}

func (p *Projects) GetFirst() Project {
	return p.Get(0)
}
