package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func makeUI() fyne.CanvasObject {
	list := widget.NewList(
		func() int {
			return 5
		},
		func() fyne.CanvasObject {
			return widget.NewCheck("An item title", nil)
		},
		func(di widget.ListItemID, co fyne.CanvasObject) {
		})

	input := widget.NewEntry()
	add := widget.NewButtonWithIcon("",
		theme.ContentAddIcon(), func() {})

	top := container.NewBorder(nil, nil, nil, add, input)
	return container.NewBorder(top, nil, nil, nil, list)
}
