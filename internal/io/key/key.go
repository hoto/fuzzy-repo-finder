package key

type Key int

const (
	Unknown Key = iota
	NonFunctional
	CtrlC
	Enter
	Backspace
	ArrowUp
	ArrowDown
	ArrowLeft
	ArrowRight
)
