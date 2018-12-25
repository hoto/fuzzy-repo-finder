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
	project := Project{Name: "PROJECT_1"}
	projects := NewProjects()
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

func Test_should_return_empty_groups(t *testing.T) {
	projects := NewProjects()

	groups := projects.ListGroups()

	assert.Equal(t, []string{}, groups)
}

func Test_should_return_groups_of_projects(t *testing.T) {
	projects := NewProjects()
	project1 := Project{Name: "PROJECT_1", Group: "GROUP_1"}
	project2 := Project{Name: "PROJECT_2", Group: "GROUP_1"}
	project3 := Project{Name: "PROJECT_3", Group: "GROUP_2"}
	projects.AddAll([]Project{project1, project2, project3})

	groups := projects.ListGroups()

	assert.EqualValues(t, []string{"GROUP_1", "GROUP_2"}, groups)
}
