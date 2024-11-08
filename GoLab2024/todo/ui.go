package main

import (
	"slices"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func makeUI(data []string) fyne.CanvasObject {
	var list *widget.List
	list = widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewCheck("An item title", nil)
		},
		func(i widget.ListItemID, co fyne.CanvasObject) {
			ch := co.(*widget.Check)
			ch.SetText(data[i])

			ch.OnChanged = func(done bool) {
				if !done {
					return
				}
				ch.SetChecked(false)

				data = slices.Delete(data, i, i+1)
				list.Refresh()
				fyne.CurrentApp().Preferences().SetStringList("todos", data)
			}
		})

	input := widget.NewEntry()
	add := widget.NewButtonWithIcon("",
		theme.ContentAddIcon(), func() {
			if input.Text == "" {
				return
			}

			data = append(data, input.Text)
			list.Refresh()
			input.SetText("")
			fyne.CurrentApp().Preferences().SetStringList("todos", data)
		})

	top := container.NewBorder(nil, nil, nil, add, input)
	return container.NewBorder(top, nil, nil, nil, list)
}
