package main

import (
	"fyne.io/fyne/v2/widget"
)

type colleagueCheckbox struct {
	mediator mediator
	radio    *widget.RadioGroup
}

var _ colleague = new(colleagueCheckbox)

func newColleagueCheckbox(options []string) *colleagueCheckbox {
	c := new(colleagueCheckbox)
	c.radio = widget.NewRadioGroup(options, c.itemStateChanged)
	return c
}

func (c *colleagueCheckbox) setMediator(m mediator) {
	c.mediator = m
}

func (c *colleagueCheckbox) setColleagueEnabled(enabled bool) {
	if enabled {
		c.radio.Enable()
	} else {
		c.radio.Disable()
	}
}

func (c *colleagueCheckbox) itemStateChanged(option string) {
	c.mediator.colleagueChanged()
}
