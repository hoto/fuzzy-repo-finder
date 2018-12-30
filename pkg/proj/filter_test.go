package proj

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	emptyQuery = ""
)

var (
	emptyProjects = NewProjects()
)

func Test_should_return_empty_projects_when_inputs_are_empty(t *testing.T) {
	filteredProjects := FuzzyMatch(emptyQuery, emptyProjects)

	assert.Equal(t, emptyProjects, filteredProjects)
}

func Test_should_return_empty_projects_when_query_is_not_empty(t *testing.T) {
	filteredProjects := FuzzyMatch("PROJECT_1", emptyProjects)

	assert.Equal(t, emptyProjectsList, filteredProjects.List())
}

func Test_should_return_single_match(t *testing.T) {
	projects := NewProjects()
	project1 := Project{Name: "PROJECT_1"}
	project2 := Project{Name: "PROJECT_2"}
	project3 := Project{Name: "PROJECT_3"}
	project4 := Project{Name: "PROJECT_4"}
	projects.AddAll([]Project{project1, project2, project3, project4})

	filteredProjects := FuzzyMatch("PROJECT_2", projects)

	assert.Equal(t, []Project{project2}, filteredProjects.List())
}

func Test_should_return_all_projects_when_query_is_empty(t *testing.T) {
	projects := NewProjects()
	project1 := Project{Name: "PROJECT_1"}
	project2 := Project{Name: "PROJECT_2"}
	project3 := Project{Name: "PROJECT_3"}
	project4 := Project{Name: "PROJECT_4"}
	projects.AddAll([]Project{project1, project2, project3, project4})

	filteredProjects := FuzzyMatch(emptyQuery, projects)

	assert.Equal(t, []Project{project1, project2, project3, project4}, filteredProjects.List())
}

func Test_should_return_all_projects_when_all_matches(t *testing.T) {
	projects := NewProjects()
	project1 := Project{Name: "A_PROJECT_1", FullPath: "FULL_PATH_1"}
	project2 := Project{Name: "A_PROJECT_2", FullPath: "FULL_PATH_2"}
	project3 := Project{Name: "B_PROJECT_3", FullPath: "FULL_PATH_3"}
	project4 := Project{Name: "B_PROJECT_4", FullPath: "FULL_PATH_4"}
	projects.AddAll([]Project{project1, project2, project3, project4})

	filteredProjects := FuzzyMatch("PROJECT_", projects)

	assert.Equal(t, []Project{project1, project2, project3, project4}, filteredProjects.List())
}

func Test_should_return_multiple_matches(t *testing.T) {
	projects := NewProjects()
	project1 := Project{Name: "A_PROJECT_1", FullPath: "FULL_PATH_1"}
	project2 := Project{Name: "A_PROJECT_2", FullPath: "FULL_PATH_2"}
	project3 := Project{Name: "B_PROJECT_3", FullPath: "FULL_PATH_3"}
	project4 := Project{Name: "B_PROJECT_4", FullPath: "FULL_PATH_4"}
	projects.AddAll([]Project{project1, project2, project3, project4})

	filteredProjects := FuzzyMatch("B_PROJECT", projects)

	assert.Equal(t, []Project{project3, project4}, filteredProjects.List())
}

func Test_should_not_change_the_order_of_projects(t *testing.T) {
	projects := NewProjects()
	project1 := Project{Name: "A_PROJECT_1", FullPath: "FULL_PATH"}
	project2 := Project{Name: "A_PROJECT_2", FullPath: "FULL_PATH"}
	project3 := Project{Name: "B_PROJECT_3", FullPath: "FULL_PATH"}
	project4 := Project{Name: "B_PROJECT_4", FullPath: "FULL_PATH"}
	projects.AddAll([]Project{project1, project2, project3, project4})

	filteredProjects := FuzzyMatch("PROJECT", projects)

	assert.Equal(t, []Project{project1, project2, project3, project4}, filteredProjects.List())
}
