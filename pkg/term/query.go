package term

type query struct {
	line []rune
}

func (q *query) Append(char rune) {
	q.line = append(q.line, char)
}

func (q *query) Runes() []rune {
	return q.line
}

func (q *query) String() string {
	return string(q.line)
}

func (q *query) DeleteLastChar() {
	if len(q.line) > 0 {
		q.line = q.line[:len(q.line)-1]
	}
}

func (q *query) DeleteLastWord() {
	if len(q.line) > 0 {
		q.line = []rune{}
	}
}

//TODO: need a test for size == 0 there is a bug atm, also add IsEmpty() function
func (q *query) Size() int {
	return len(q.line)
}
