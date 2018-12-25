package proj

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	emptyProjects = make([]Project, 0)
)

func Test_should_return_empty_projects(t *testing.T) {
	projects := NewProjects()

	assert.Equal(t, emptyProjects, projects.List())
}

func Test_should_retain_a_project(t *testing.T) {
	projects := NewProjects()
	project := Project{Name: "PROJECT_1"}
	projects.Add(project)

	assert.Equal(t, projects.List(), []Project{project})
}

func Test_should_retain_added_projects(t *testing.T) {
	projects := NewProjects()
	newProjects := []Project{
		{Name: "PROJECT_1"},
		{Name: "PROJECT_2"},
	}

	projects.AddAll(newProjects)

	assert.Equal(t, projects.List(), newProjects)
}
