package io

import (
	"github.com/hoto/fuzzy-repo-finder/internal/proj"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

const (
	projectsRoot = "/home/user/projects"
	git          = ".git"
)

var (
	emptyProjects = make([]proj.Project, 0)
)

func Test_should_return_empty_list_when_no_directory_matches(t *testing.T) {
	disk := new(MockDisk)
	disk.On("FindDirs", projectsRoot, git).Return([]string{})
	filesystem := Filesystem{disk}

	projects := filesystem.FindGitProjects(projectsRoot)

	assert.Equal(t, emptyProjects, projects.List())
}

func Test_should_return_matching_projects(t *testing.T) {
	disk := new(MockDisk)
	disk.On("FindDirs", projectsRoot, git).Return([]string{
		"/home/user/projects/project1/.git",
		"/home/user/projects/project2/.git",
	})
	filesystem := Filesystem{disk}

	projects := filesystem.FindGitProjects(projectsRoot)

	project1 := proj.Project{
		Name:     "project1",
		Group:    "",
		FullPath: "/home/user/projects/project1",
	}
	project2 := proj.Project{
		Name:     "project2",
		Group:    "",
		FullPath: "/home/user/projects/project2",
	}
	expectedProjects := []proj.Project{project1, project2}
	assert.Equal(t, expectedProjects, projects.List())
}

func Test_should_return_matching_projects_inside_a_group(t *testing.T) {
	disk := new(MockDisk)
	disk.On("FindDirs", projectsRoot, git).Return([]string{
		"/home/user/projects/dirA/project1/.git",
		"/home/user/projects/dirB/project2/.git",
	})
	filesystem := Filesystem{disk}

	projects := filesystem.FindGitProjects(projectsRoot)

	project1 := proj.Project{
		Name:     "project1",
		Group:    "dirA",
		FullPath: "/home/user/projects/dirA/project1",
	}
	project2 := proj.Project{
		Name:     "project2",
		Group:    "dirB",
		FullPath: "/home/user/projects/dirB/project2",
	}
	expectedProjects := []proj.Project{project1, project2}
	assert.Equal(t, expectedProjects, projects.List())
}

func Test_should_return_matching_projects_inside_a_multiple_level_group(t *testing.T) {
	disk := new(MockDisk)
	disk.On("FindDirs", projectsRoot, git).Return([]string{
		"/home/user/projects/dirA1/dirA2/dirA3/project1/.git",
		"/home/user/projects/dirB1/dirB2/dirB3/project2/.git",
	})
	filesystem := Filesystem{disk}

	projects := filesystem.FindGitProjects(projectsRoot)

	project1 := proj.Project{
		Name:     "project1",
		Group:    "dirA1/dirA2/dirA3",
		FullPath: "/home/user/projects/dirA1/dirA2/dirA3/project1",
	}
	project2 := proj.Project{
		Name:     "project2",
		Group:    "dirB1/dirB2/dirB3",
		FullPath: "/home/user/projects/dirB1/dirB2/dirB3/project2",
	}
	expectedProjects := []proj.Project{project1, project2}
	assert.Equal(t, expectedProjects, projects.List())
}

type MockDisk struct {
	mock.Mock
}

func (m *MockDisk) FindDirs(root string, matchDir string) []string {
	args := m.Called(root, matchDir)
	return args.Get(0).([]string)
}
