package main

import (
	"fyne.io/fyne/v2/widget"
)

type colleagueTextField struct {
	mediator mediator
	entry    *widget.Entry
}

var _ colleague = new(colleagueTextField)

func newColleagueTextField(text string, columns int) *colleagueTextField {
	f := new(colleagueTextField)
	f.entry = widget.NewEntry()
	f.entry.Text = text
	f.entry.CursorColumn = columns
	f.entry.OnChanged = f.textValueChanged
	return f
}

func (f *colleagueTextField) setMediator(m mediator) {
	f.mediator = m
}

func (f *colleagueTextField) setColleagueEnabled(enabled bool) {
	if enabled {
		f.entry.Enable()
	} else {
		f.entry.Disable()
	}
}

func (f *colleagueTextField) textValueChanged(s string) {
	f.mediator.colleagueChanged()
}
