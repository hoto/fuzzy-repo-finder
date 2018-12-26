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

func (f *field) QueryRunes() []rune {
	return f.query
}

func (f *field) QueryString() string {
	return string(f.query)
}

func (f *field) QuerySize() int {
	return len(f.query)
}

func (f *field) QueryIsEmpty() bool {
	return f.QuerySize() == 0
}

func (f *field) AppendToQuery(char rune) {
	f.query = append(f.query, char)
}

func (f *field) DeleteLastQueryChar() {
	if len(f.query) > 0 {
		f.query = f.query[:len(f.query)-1]
	}
}

func (f *field) EraseQuery() {
	if len(f.query) > 0 {
		f.query = []rune{}
	}
}

func (f *field) TitleRunes() []rune {
	return f.title
}

func (f *field) TitleSize() int {
	return len(f.title)
}

func (f *field) FieldSize() int {
	return len(f.title) + len(f.query)
}
