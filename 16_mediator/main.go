package main

import (
	// "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	// "fmt"
	// "fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	// "log"
	"time"
)

func updateClock(clock *widget.Label) {
	clock.SetText(time.Now().Format("15:04:05"))
}

func main() {
	a := app.New()
	f := NewLoginFrame("Mesiator Sample", a)
	f.window.Show()

	// w := a.NewWindow("hello")
	// // hello := widget.NewLabel("hello Fyne!")
	// clock := widget.NewLabel("")
	// updateClock(clock)

	// w.SetContent(clock)
	// w.SetMaster()
	// w.Show()

	a.Run()
	// // w.SetContent(container.NewVBox(
	// // 	hello,
	// // 	widget.NewButton("Hi!", func() {
	// // 		hello.SetText("Welcome :)")
	// // 	}),
	// // ))

	// go func() {
	// 	for range time.Tick(time.Second) {
	// 		updateClock(clock)
	// 	}
	// }()

	// nameEntry := widget.NewEntry()
	// passEntry := widget.NewEntry()
	// nameEntry.Disable()
	// passEntry.Disable()
	// go func() {
	// 	for range time.Tick(time.Millisecond * 10) {
	// 		if len(nameEntry.Text) > 0 {
	// 			passEntry.Enable()
	// 		}
	// 	}
	// }()

	// w2 := a.NewWindow("Larger")
	// w2.Resize(fyne.NewSize(700, 400))
	// w2.SetContent(widget.NewButton("More content", func() {
	// 	w3 := a.NewWindow("third")
	// 	form := &widget.Form{
	// 		Items: []*widget.FormItem{
	// 			{Text: "Username:", Widget: nameEntry},
	// 			{Text: "Password:", Widget: passEntry},
	// 		},
	// 		OnSubmit: func() {
	// 			log.Printf("Sumbit name: %v, pass: %v\n", nameEntry.Text, passEntry.Text)
	// 			passEntry.SetPlaceHolder("place Holder")

	// 		},
	// 		OnCancel: func() {
	// 			w3.Close()
	// 		},
	// 	}
	// 	radio := widget.NewRadioGroup([]string{"Guest", "Login"}, func(r string) {
	// 		if r == "Guest" {
	// 			nameEntry.Disable()
	// 			passEntry.Disable()
	// 		}
	// 		if r == "Login" {
	// 			nameEntry.Enable()
	// 			log.Println(nameEntry.Text)
	// 			if len(nameEntry.Text) > 0 {
	// 				passEntry.Enable()
	// 			}
	// 		}
	// 	})
	// 	radioBox := container.New(layout.NewHBoxLayout(), radio)
	// 	formBox := container.New(layout.NewMaxLayout(), form)
	// 	w3.SetContent(container.New(layout.NewVBoxLayout(), radioBox, formBox))
	// 	w3.Resize(fyne.NewSize(700, 400)

	// 	w3.Show()
	// }))

	// // w.ShowAndRun()
	// w2.Show()
	// a.Run()
	// fmt.Println("Existed")
}
