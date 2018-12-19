package io

import (
	"github.com/hoto/fuzzy-repo-finder/internal/project"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

const (
	directory = "directory"
)

func Test_should_return_empty_list_when_no_directory_matches(t *testing.T) {
	disk := new(MockDisk)
	disk.On("FindDirs", directory, ".git").Return([]string{})
	filesystem := Filesystem{disk}

	projects := filesystem.FindProjects(directory)

	var expectedProjects []project.Project
	assert.Equal(t, projects, expectedProjects)
}

type MockDisk struct {
	mock.Mock
}

func (mock *MockDisk) FindDirs(root string, matchDir string) []string {
	//args := mock.Called(root)
	return []string{}
}
