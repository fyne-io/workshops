package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")

	hello := widget.NewLabel("Hello World!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Welcome", func() {
			hello.SetText("Fyne Conf!")
		}),
	))

	w.ShowAndRun()
}
