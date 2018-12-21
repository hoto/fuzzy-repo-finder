package io

import (
	"github.com/hoto/fuzzy-repo-finder/internal/io/key"
	"github.com/nsf/termbox-go"
)

type Keyboard struct{}

func newKeyboard() *Keyboard {
	return &Keyboard{}
}

func (k Keyboard) WaitForKeyPress() key.Key {
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
