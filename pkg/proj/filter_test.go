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
