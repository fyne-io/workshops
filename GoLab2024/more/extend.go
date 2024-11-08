package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.NewWithID("io.fyne.workshop.extend")
	w := a.NewWindow("Extend Label")

	w.SetContent(newRightClickLabel("Right Click Me", w.Canvas()))

	w.ShowAndRun()
}

type rightClickLabel struct {
	widget.Label

	canvas fyne.Canvas
}

func newRightClickLabel(text string, c fyne.Canvas) fyne.Widget {
	r := &rightClickLabel{canvas: c}
	r.Text = text
	r.ExtendBaseWidget(r)

	return r
}

func (r *rightClickLabel) TappedSecondary(e *fyne.PointEvent) {
	m := fyne.NewMenu("",
		fyne.NewMenuItem("Item 1", func() {}),
		fyne.NewMenuItem("Item 2", func() {}))
	widget.ShowPopUpMenuAtPosition(m, r.canvas, e.AbsolutePosition)
}
