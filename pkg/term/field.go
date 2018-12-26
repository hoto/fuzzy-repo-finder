package term

type field struct {
	title []rune
	query []rune
}

func NewField(title string, query string) *field {
	return &field{
		title: []rune(title),
		query: []rune(query),
	}
}

func (f *field) queryRunes() []rune {
	return f.query
}

func (f *field) queryString() string {
	return string(f.query)
}

func (f *field) querySize() int {
	return len(f.query)
}

func (f *field) queryIsEmpty() bool {
	return f.querySize() == 0
}

func (f *field) appendToQuery(char rune) {
	f.query = append(f.query, char)
}

func (f *field) deleteLastQueryChar() {
	if len(f.query) > 0 {
		f.query = f.query[:len(f.query)-1]
	}
}

func (f *field) eraseQuery() {
	if len(f.query) > 0 {
		f.query = []rune{}
	}
}

func (f *field) titleRunes() []rune {
	return f.title
}

func (f *field) titleSize() int {
	return len(f.title)
}

func (f *field) fieldSize() int {
	return len(f.title) + len(f.query)
}
