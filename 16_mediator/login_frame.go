package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"log"
)

var guest = "Guest"
var login = "Login"

type loginFrame struct {
	// checkGuest   *colleagueCheckbox
	// checkLogin   *colleagueCheckbox
	radio        *colleagueCheckbox
	textUser     *colleagueTextField
	textPass     *colleagueTextField
	buttonOk     *colleagueButton
	buttonCancel *colleagueButton
	window       fyne.Window
}

func NewLoginFrame(title string, app fyne.App) *loginFrame {
	f := new(loginFrame)
	f.window = app.NewWindow(title)
	f.createColleagues()
	return f

}

var _ mediator = new(loginFrame)

func (f *loginFrame) createColleagues() {
	// CheckBox
	f.radio = newColleagueCheckbox([]string{guest, login})
	f.radio.radio.Selected = guest
	f.buttonOk = newColleagueButton("OK", func() {
		log.Printf("Sumbit name: %v, pass: %v\n", f.textPass.entry.Text, f.textPass.entry.Text)
	})
	f.buttonCancel = newColleagueButton("Cancel", func() {
		log.Printf("Cancel name: %v, pass: %v\n", f.textPass.entry.Text, f.textPass.entry.Text)
	})
	f.textUser = newColleagueTextField("", 10)
	f.textPass = newColleagueTextField("", 10)
	f.textUser.entry.SetPlaceHolder("username")
	f.textPass.entry.SetPlaceHolder("password")
	f.radio.setMediator(f)
	f.textUser.setMediator(f)
	f.textPass.setMediator(f)
	f.buttonOk.setMediator(f)
	f.buttonCancel.setMediator(f)

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Username:", Widget: f.textUser.entry},
			{Text: "Password:", Widget: f.textPass.entry},
		},

		// submitButton: f.buttonOk.button,
		// OnSubmit: f.buttonOk.button.OnTapped,
		// OnCancel: f.buttonCancel.button.OnTapped,
	}
	radioContainer := container.New(layout.NewHBoxLayout(), f.radio.radio)
	formContainer := container.New(layout.NewMaxLayout(), form)
	OkContainer := container.New(layout.NewMaxLayout(), f.buttonOk.button)
	CancelContainer := container.New(layout.NewMaxLayout(), f.buttonCancel.button)
	c := container.New(layout.NewVBoxLayout(), radioContainer, formContainer, OkContainer, CancelContainer)
	f.window.SetContent(c)
	f.window.Resize(fyne.NewSize(1000, 700))
}
func (f *loginFrame) colleagueChanged() {
	if f.radio.radio.Selected == guest {
		f.textUser.setColleagueEnabled(false)
		f.textPass.setColleagueEnabled(false)
		f.buttonOk.setColleagueEnabled(true)
	} else {
		f.textUser.setColleagueEnabled(true)
		f.userpassChanged()
	}
}

func (f *loginFrame) userpassChanged() {
	if len(f.textUser.entry.Text) > 0 {
		f.textPass.entry.Enable()
		if len(f.textPass.entry.Text) > 0 {
			f.buttonOk.setColleagueEnabled(true)
		} else {
			f.buttonOk.setColleagueEnabled(false)
		}
	} else {
		f.textPass.setColleagueEnabled(false)
		f.buttonOk.setColleagueEnabled(false)
	}
}
