package term

type Query struct {
	line []rune
}

func (q *Query) Append(char rune) {
	q.line = append(q.line, char)
}

func (q *Query) Runes() []rune {
	return q.line
}

func (q *Query) String() string {
	return string(q.line)
}

func (q *Query) DeleteLastChar() {
	if len(q.line) > 0 {
		q.line = q.line[:len(q.line)-1]
	}
}

func (q *Query) DeleteLastWord() {
	if len(q.line) > 0 {
		q.line = []rune{}
	}
}

func (q *Query) Size() int {
	return len(q.line)
}
