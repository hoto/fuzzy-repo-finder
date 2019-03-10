package config

import (
	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

//func TestMain(m *testing.M) {
//	retCode := m.Run()
//	os.Exit(retCode)
//}

func Test_show_usage_when_project_roots_not_passed(t *testing.T) {
	mockOsExit()

	ParseArgsAndFlags()
}

func mockOsExit() {
	fakeExit := func(int) {
		panic("os.Exit called")
	}
	patch := monkey.Patch(os.Exit, fakeExit)
	defer patch.Unpatch()
}

func Test_set_default_values_of_flags(t *testing.T) {
	os.Args = []string{"cmd", "--projectRoots", "PROJECT_ROOT"}

	ParseArgsAndFlags()

	assert.Equal(t, "", Version)
	assert.Equal(t, "", ShortCommit)
	assert.Equal(t, "", BuildDate)
	assert.Equal(t, false, Debug)
	assert.Equal(t, []string{""}, ProjectsRoots)
	assert.Equal(t, "", ProjectNameFilter)
	assert.Equal(t, "", SelectedProjectPath)
}

func Test_prints_version(t *testing.T) {
	//ParseArgsAndFlags()

	assert.Equal(t, true, Debug)
}

func TestDoomed(t *testing.T) {
	fakeExit := func(int) {
		panic("os.Exit called")
	}
	patch := monkey.Patch(os.Exit, fakeExit)
	defer patch.Unpatch()
	assert.PanicsWithValue(t, "os.Exit called", ParseArgsAndFlags,
		"os.Exit was not called")
}
