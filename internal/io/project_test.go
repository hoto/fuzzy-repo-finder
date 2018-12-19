package io

//import (
//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/mock"
//	"testing"
//)
//
//const (
//	directory = "directory"
//)
//
//func TestTrue(t *testing.T) {
//	filesystem := new(MockedFilesystem)
//	filesystem.On("FindDirs", directory).Return([]string{"one"})
//	io := IO{
//		filesystem: filesystem,
//	}
//
//	response := io.FindProjects(directory)
//
//	expected := []string{"one"}
//	assert.Equal(t, response, expected)
//}
//
//type MockedFilesystem struct {
//	mock.Mock
//}
//
//func (mock *MockedFilesystem) FindDirs(dir string) []string {
//	//args := mock.Called(dir)
//	return []string{"one"}
//}
