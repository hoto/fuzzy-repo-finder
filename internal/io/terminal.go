package io

import (
	"github.com/nsf/termbox-go"
)

type Terminal struct{}

func NewTerminal() *Terminal {
	return &Terminal{}
}

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
