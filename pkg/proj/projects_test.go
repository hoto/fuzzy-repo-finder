package proj

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	emptyProject      = Project{}
	emptyProjectsList = make([]Project, 0)
)

func Test_be_empty(t *testing.T) {
	projects := NewProjects()

	assert.Equal(t, 0, projects.Size())
	assert.Equal(t, 0, projects.Len())
}

func Test_return_empty_projects(t *testing.T) {
	projects := NewProjects()

	assert.Equal(t, emptyProjectsList, projects.List())
}

func Test_return_empty_project(t *testing.T) {
	projects := NewProjects()

	assert.Equal(t, emptyProject, projects.Get(0))
}

func Test_return_empty_first_project(t *testing.T) {
	projects := NewProjects()

	assert.Equal(t, emptyProject, projects.GetFirst())
}

func Test_return_first_project(t *testing.T) {
	project1 := Project{Name: "PROJECT_1"}
	project2 := Project{Name: "PROJECT_2"}
	projects := NewProjects()
	projects.AddAll([]Project{project1, project2})

	assert.Equal(t, project1, projects.GetFirst())
}

func Test_retain_a_project(t *testing.T) {
	project := Project{Name: "PROJECT_1"}
	projects := NewProjects()
	projects.Add(project)

	assert.Equal(t, projects.List(), []Project{project})
	assert.Equal(t, projects.Get(0), project)
}

func Test_have_one_element(t *testing.T) {
	project := Project{Name: "PROJECT_1"}
	projects := NewProjects()
	projects.Add(project)

	assert.Equal(t, 1, projects.Size())
}

func Test_retain_added_projects(t *testing.T) {
	projects := NewProjects()
	newProjects := []Project{
		{Name: "PROJECT_1"},
		{Name: "PROJECT_2"},
	}

	projects.AddAll(newProjects)

	assert.Equal(t, projects.List(), newProjects)
}

func Test_return_empty_groups(t *testing.T) {
	projects := NewProjects()

	groups := projects.ListGroups()

	assert.Equal(t, []string{}, groups)
}

func Test_list_groups_in_order(t *testing.T) {
	projects := NewProjects()
	project1 := Project{Name: "PROJECT_1", Group: "GROUP_1"}
	project2 := Project{Name: "PROJECT_2", Group: "GROUP_1"}
	project3 := Project{Name: "PROJECT_3", Group: "GROUP_2"}
	project4 := Project{Name: "PROJECT_4", Group: "GROUP_3"}
	project5 := Project{Name: "PROJECT_5", Group: "GROUP_1"}
	projects.AddAll([]Project{project1, project2, project3, project4, project5})

	groups := projects.ListGroups()

	assert.EqualValues(t, []string{"GROUP_1", "GROUP_2", "GROUP_3"}, groups)
}

func Test_make_an_empty_copy(t *testing.T) {
	projects := NewProjects()

	projectsCopy := projects.Copy()

	assert.Equal(t, projectsCopy, projects)
}

func Test_make_a_copy(t *testing.T) {
	projects := NewProjects()
	project1 := Project{Name: "PROJECT_1", Group: "GROUP_1"}
	projects.AddAll([]Project{project1})

	projectsCopy := projects.Copy()

	assert.Equal(t, projectsCopy, projects)
}

func Test_make_a_deep_copy(t *testing.T) {
	projects := NewProjects()
	project1 := Project{Name: "PROJECT_1", Group: "GROUP_1"}
	project2 := Project{Name: "PROJECT_2", Group: "GROUP_2"}
	projects.AddAll([]Project{project1})

	projectsCopy := projects.Copy()
	projects.Add(project2)

	assert.NotEqual(t, projectsCopy, projects)
}
