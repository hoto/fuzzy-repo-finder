package io

import (
	"github.com/hoto/fuzzy-repo-finder/internal/io/key"
	"github.com/nsf/termbox-go"
)

type Terminal struct{}

func NewTerminal() *Terminal {
	return &Terminal{}
}

func (t Terminal) Init() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
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

func (t Terminal) WaitForKeyPress() key.Key {
	switch event := termbox.PollEvent(); event.Type {
	case termbox.EventKey:
		switch event.Key {
		case 0:
			return key.NonFunctional
		case termbox.KeyCtrlC:
			return key.CtrlC
		case termbox.KeyEnter:
			return key.Enter
		case termbox.KeyBackspace, termbox.KeyBackspace2:
			return key.Backspace
		case termbox.KeyArrowDown:
			return key.ArrowDown
		case termbox.KeyArrowUp:
			return key.ArrowUp
		case termbox.KeyArrowLeft:
			return key.ArrowLeft
		case termbox.KeyArrowRight:
			return key.ArrowRight
		}

	}
	return key.Unknown
}
