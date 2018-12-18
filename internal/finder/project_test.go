package finder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrue(t *testing.T) {
	var a = "Hello"
	var b = "Hello"

	assert.Equal(t, a, b, "The two words should be the same.")
}
