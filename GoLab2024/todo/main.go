package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.NewWithID("io.fyne.workshop.todo")
	w := a.NewWindow("TODO")

	data := a.Preferences().StringListWithFallback("todos",
		[]string{"Use this TODO list", "Build more Fyne apps"})
	w.SetContent(makeUI(data))

	w.Resize(fyne.NewSize(180, 240))
	w.ShowAndRun()
}
