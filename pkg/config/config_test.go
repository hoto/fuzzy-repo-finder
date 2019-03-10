package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_set_default_values_of_flags(t *testing.T) {
	ParseArgsAndFlags()

	assert.Equal(t, "", Version)
	assert.Equal(t, "", ShortCommit)
	assert.Equal(t, "", BuildDate)
	assert.Equal(t, false, Debug)
	assert.Equal(t, []string{"/projects"}, ProjectsRoots)
	assert.Equal(t, "", ProjectNameFilter)
	assert.Equal(t, "", SelectedProjectPath)
}
