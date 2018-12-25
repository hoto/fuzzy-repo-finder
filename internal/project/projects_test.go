package project

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

func Test_should_return_a_project(t *testing.T) {
	projects := NewProjects()
	project := Project{}
	projects.Add(project)

	assert.Equal(t, projects.List(), []Project{project})
}
