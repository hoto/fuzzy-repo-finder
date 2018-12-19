package io

import (
	"github.com/hoto/fuzzy-repo-finder/internal/project"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

const (
	projectsRoot = "/projects"
	git          = ".git"
)

var (
	emptyProjects []project.Project
)

func Test_should_return_empty_list_when_no_directory_matches(t *testing.T) {
	disk := new(MockDisk)
	disk.On("FindDirs", projectsRoot, git).Return([]string{})
	filesystem := Filesystem{disk}

	projects := filesystem.FindGitProjects(projectsRoot)

	assert.Equal(t, emptyProjects, projects)
}

func Test_should_return_matching_projects_with_no_group(t *testing.T) {
	disk := new(MockDisk)
	disk.On("FindDirs", projectsRoot, git).Return([]string{
		"/projects/project1/.git",
		"/projects/project2/.git",
	})
	filesystem := Filesystem{disk}

	projects := filesystem.FindGitProjects(projectsRoot)

	project1 := project.Project{Name: "project1", Group: "", FullPath: "/projects/project1"}
	project2 := project.Project{Name: "project2", Group: "", FullPath: "/projects/project2"}
	expectedProjects := []project.Project{project1, project2}
	assert.Equal(t, expectedProjects, projects)
}

type MockDisk struct {
	mock.Mock
}

func (m *MockDisk) FindDirs(root string, matchDir string) []string {
	args := m.Called(root, matchDir)
	return args.Get(0).([]string)
}
