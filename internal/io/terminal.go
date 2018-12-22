package io

import (
	"github.com/nsf/termbox-go"
)

type Terminal struct{}

func NewTerminal() *Terminal {
	return &Terminal{}
}

// Initializes terminal. This function must be called before any other functions.
// Terminal must be closed using Close() function before exiting the application.
//
//      keyboard := terminal.Init()
//      defer terminal.Close()
func (t Terminal) Init() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
}

func (t Terminal) Close() {
	termbox.Close()
}

func (t Terminal) Cycle() int {
	event := termbox.PollEvent()
	if event.Type == termbox.EventKey {
		char := event.Ch
		switch event.Key {
		case 0:
			t.writeLine(char)
		case termbox.KeyCtrlC:
			return 1
		}
	}
	return 0
}

func (t Terminal) flush() {
	err := termbox.Flush()
	if err != nil {
		panic(err)
	}
}

func (t Terminal) clear() {
	err := termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	if err != nil {
		panic(err)
	}
}

func (t Terminal) writeLine(char rune) {
	t.clear()
	termbox.SetCell(1, 2, char, termbox.ColorGreen, termbox.ColorDefault)
	t.flush()
}
