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

	assert.Equal(t, projects.List(), emptyProjects)
}
