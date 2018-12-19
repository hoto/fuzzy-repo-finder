package io

import (
	"github.com/hoto/fuzzy-repo-finder/internal/project"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var (
	directory     = "directory"
	emptyProjects = []project.Project{}
)

func Test_should_return_empty_list_when_no_directory_matches(t *testing.T) {
	disk := new(MockDisk)
	disk.On("FindDirs", directory, ".git").Return([]string{})
	filesystem := Filesystem{disk}

	projects := filesystem.FindGitProjects(directory)

	var expectedProjects []project.Project
	assert.Equal(t, projects, expectedProjects)
}

type MockDisk struct {
	mock.Mock
}

func (m *MockDisk) FindDirs(root string, matchDir string) []string {
	args := m.Called(root, matchDir)
	return args.Get(0).([]string)
}
