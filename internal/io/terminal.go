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
func (t Terminal) Init() *Keyboard {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	return newKeyboard()
}

func (t Terminal) Close() {
	termbox.Close()
}

func (t Terminal) Flush() {
	err := termbox.Flush()
	if err != nil {
		panic(err)
	}
}

func (t Terminal) Clear() {
	err := termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	if err != nil {
		panic(err)
	}
}
