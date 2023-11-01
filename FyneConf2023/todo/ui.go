package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func makeUI(data binding.StringList) fyne.CanvasObject {
	list := widget.NewListWithData(data,
		func() fyne.CanvasObject {
			return widget.NewCheck("An item title", nil)
		},
		func(di binding.DataItem, co fyne.CanvasObject) {
			str, _ := di.(binding.String).Get()
			ch := co.(*widget.Check)
			ch.OnChanged = func(done bool) {
				if !done {
					return
				}
				ch.SetChecked(false)

				// TODO StringList.Remove
				items, _ := data.Get()
				items = removeItem(items, ch.Text)
				data.Set(items)
			}
			ch.SetText(str)
		})

	input := widget.NewEntry()
	add := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		if input.Text == "" {
			return
		}

		data.Prepend(input.Text)
		input.SetText("")
	})

	top := container.NewBorder(nil, nil, nil, add, input)
	return container.NewBorder(top, nil, nil, nil, list)
}

func removeItem(items []string, rm string) []string {
	for i, s := range items {
		if s != rm {
			continue
		}

		if i == len(items)-1 {
			items = items[:i]
		} else {
			items = append(items[:i], items[i+1:]...)
		}
		break
	}

	return items
}
