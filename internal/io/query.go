package io

type Query struct {
	line []rune
}

func (q *Query) Append(char rune) {
	q.line = append(q.line, char)
}

func (q *Query) Read() []rune {
	return q.line
}

func (q *Query) DeleteLastChar() {
	if len(q.line) > 0 {
		q.line = q.line[:len(q.line)-1]
	}
}
