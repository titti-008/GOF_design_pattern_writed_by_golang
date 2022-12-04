package main

import (
	"fyne.io/fyne/v2/widget"
)

type colleagueButton struct {
	mediator mediator
	button   *widget.Button
}

var _ colleague = new(colleagueButton)

func newColleagueButton(caption string, f func()) *colleagueButton {
	b := new(colleagueButton)
	b.button = &widget.Button{
		Text:     caption,
		OnTapped: f,
	}
	return b
}

func (b *colleagueButton) setMediator(m mediator) {
	b.mediator = m
}

func (b *colleagueButton) setColleagueEnabled(enabled bool) {
	if enabled {
		b.button.Enable()
	} else {
		b.button.Disable()
	}
}
