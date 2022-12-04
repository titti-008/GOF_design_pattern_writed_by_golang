package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"time"
)

func main() {
	fmt.Println("Hello State")

	a := app.New()
	frame := newSafeFrame("State Sample", a)

	go func() {
		for range time.Tick(time.Second) {
			hour := time.Now().Second() % 24
			frame.setClock(hour)
		}
	}()

	frame.window.Show()
	a.Run()

}
