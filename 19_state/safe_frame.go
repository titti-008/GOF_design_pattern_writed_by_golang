package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type safeFrame struct {
	window      fyne.Window
	textClock   *widget.Entry
	textArea    *widget.Entry
	buttonUse   *widget.Button
	buttonAlarm *widget.Button
	buttonPhone *widget.Button
	buttonExit  *widget.Button
	state       state
}

var _ context = new(safeFrame)

func newSafeFrame(title string, app fyne.App) *safeFrame {
	f := new(safeFrame)
	f.window = app.NewWindow(title)
	f.state = getDayState()
	f.textClock = widget.NewEntry()
	f.textArea = widget.NewMultiLineEntry()

	f.buttonUse = widget.NewButton("Use emma safe.", func() {
		f.state.doUse(f)
	})
	f.buttonAlarm = widget.NewButton("Emergency Bell!!!", func() {
		f.state.doAlarm(f)
	})
	f.buttonPhone = widget.NewButton("Normal phone.", func() {
		f.state.doPhone(f)
	})

	clockContainer := container.New(layout.NewMaxLayout(), f.textArea)
	textContainer := container.New(layout.NewMaxLayout(), f.textClock)
	btnContainer := container.New(layout.NewHBoxLayout(), f.buttonUse, f.buttonAlarm, f.buttonPhone)
	mainContainer := container.New(layout.NewVBoxLayout(), clockContainer, textContainer, btnContainer)

	f.window.SetContent(mainContainer)
	return f
}

func (f *safeFrame) setClock(hour int) {
	clockStr := fmt.Sprintf("Now is %02d:00 o'clock", hour)
	fmt.Println(clockStr)
	f.textClock.SetText(clockStr)
	f.state.doClock(f, hour)
}

func (f *safeFrame) changeState(s state) {
	fmt.Printf("Change state from %v to %v.\n", f.state, s)
	f.state = s
}

func (f *safeFrame) callSecurityCenter(msg string) {
	f.textArea.SetText(f.textArea.Text + "call!" + msg + "\n")
}

func (f *safeFrame) recordLog(msg string) {
	f.textArea.SetText(f.textArea.Text + "record ..." + msg + "\n")
}
