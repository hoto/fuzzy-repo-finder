package term

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_should_have_zero_query_size(t *testing.T) {
	query := field{}

	assert.Equal(t, 0, query.QuerySize())
}

func Test_should_be_empty(t *testing.T) {
	query := NewField("", "")

	assert.Equal(t, true, query.QueryIsEmpty())
}

func Test_should_return_query_size(t *testing.T) {
	query := NewField("", "QUERY")

	size := query.QuerySize()

	assert.Equal(t, 5, size)
}

func Test_should_return_title_size(t *testing.T) {
	query := NewField("TITLE", "")

	size := query.TitleSize()

	assert.Equal(t, 5, size)
}

func Test_should_return_whole_field_size(t *testing.T) {
	query := NewField("TITLE", "QUERY")

	size := query.FieldSize()

	assert.Equal(t, 10, size)
}

func Test_should_erase_query(t *testing.T) {
	query := NewField("", "QUERY")

	query.EraseQuery()

	assert.Equal(t, true, query.QueryIsEmpty())
}

func Test_should_leave_query_empty_when_attempting_to_delete_last_char(t *testing.T) {
	query := NewField("", "")

	query.DeleteLastQueryChar()

	assert.Equal(t, "", query.QueryString())
}

func Test_should_append_char_to_query(t *testing.T) {
	query := NewField("", "QUERY")
	xRune := rune(120)

	query.AppendToQuery(xRune)

	assert.Equal(t, "QUERYx", query.QueryString())
}

func Test_should_delete_last_query_char(t *testing.T) {
	query := NewField("", "QUERY")

	query.DeleteLastQueryChar()

	assert.Equal(t, "QUER", query.QueryString())
}

func Test_should_return_query_runes(t *testing.T) {
	query := NewField("", "QUERY")

	runes := query.QueryRunes()

	assert.Equal(t, []int32{81, 85, 69, 82, 89}, runes)
}

func Test_should_return_title_runes(t *testing.T) {
	query := NewField("TITLE", "")

	runes := query.TitleRunes()

	assert.Equal(t, []int32{84, 73, 84, 76, 69}, runes)
}
