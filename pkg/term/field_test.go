package term

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_should_have_zero_query_size(t *testing.T) {
	query := field{}

	assert.Equal(t, 0, query.querySize())
}

func Test_should_have_empty_query(t *testing.T) {
	query := NewField("", "")

	assert.Equal(t, true, query.queryIsEmpty())
}

func Test_should_return_query_size(t *testing.T) {
	query := NewField("", "QUERY")

	size := query.querySize()

	assert.Equal(t, 5, size)
}

func Test_should_return_title_size(t *testing.T) {
	query := NewField("TITLE", "")

	size := query.titleSize()

	assert.Equal(t, 5, size)
}

func Test_should_return_whole_field_size(t *testing.T) {
	query := NewField("TITLE", "QUERY")

	size := query.fieldSize()

	assert.Equal(t, 10, size)
}

func Test_should_erase_query(t *testing.T) {
	query := NewField("", "QUERY")

	query.eraseQuery()

	assert.Equal(t, true, query.queryIsEmpty())
}

func Test_should_leave_query_empty_when_attempting_to_delete_last_char(t *testing.T) {
	query := NewField("", "")

	query.deleteLastQueryChar()

	assert.Equal(t, "", query.queryString())
}

func Test_should_append_char_to_query(t *testing.T) {
	query := NewField("", "QUERY")
	xRune := rune(120)

	query.appendToQuery(xRune)

	assert.Equal(t, "QUERYx", query.queryString())
}

func Test_should_delete_last_query_char(t *testing.T) {
	query := NewField("", "QUERY")

	query.deleteLastQueryChar()

	assert.Equal(t, "QUER", query.queryString())
}

func Test_should_return_query_runes(t *testing.T) {
	query := NewField("", "QUERY")

	runes := query.queryRunes()

	assert.Equal(t, []int32{81, 85, 69, 82, 89}, runes)
}

func Test_should_return_title_runes(t *testing.T) {
	query := NewField("TITLE", "")

	runes := query.titleRunes()

	assert.Equal(t, []int32{84, 73, 84, 76, 69}, runes)
}
