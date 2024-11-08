package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.NewWithID("io.fyne.workshop.custom")
	w := a.NewWindow("IconSwapper")

	w.SetContent(newImageSwapper(theme.HomeIcon(), theme.ComputerIcon()))

	w.ShowAndRun()
}

type imageSwapper struct {
	widget.BaseWidget

	Reversed bool

	img1, img2 fyne.Resource
}

func newImageSwapper(img1, img2 fyne.Resource) fyne.Widget {
	i := &imageSwapper{img1: img1, img2: img2}
	i.ExtendBaseWidget(i)

	return i
}

func (i *imageSwapper) CreateRenderer() fyne.WidgetRenderer {
	return &imageSwapRenderer{img: i}
}

func (i *imageSwapper) Swap() {
	i.Reversed = !i.Reversed
	i.Refresh()
}

func (i *imageSwapper) Tapped(_ *fyne.PointEvent) {
	i.Swap()
}

type imageSwapRenderer struct {
	img *imageSwapper

	left, right *canvas.Image
	content     *fyne.Container
}

func (r *imageSwapRenderer) Destroy() {}

func (r *imageSwapRenderer) Layout(s fyne.Size) {
	if r.content == nil {
		return
	}

	r.content.Resize(s)
}

func (r *imageSwapRenderer) MinSize() fyne.Size {
	iconMin := theme.SizeForWidget(theme.SizeNameInlineIcon, r.img)
	pad := theme.SizeForWidget(theme.SizeNamePadding, r.img)

	return fyne.NewSize(iconMin*2+pad, iconMin)
}

func (r *imageSwapRenderer) Objects() []fyne.CanvasObject {
	if r.content == nil {
		r.left = canvas.NewImageFromResource(r.img.img1)
		r.right = canvas.NewImageFromResource(r.img.img2)

		r.content = container.NewGridWithColumns(2, r.left, r.right)
	}

	return []fyne.CanvasObject{r.content}
}

func (r *imageSwapRenderer) Refresh() {
	if r.img.Reversed {
		r.left.Resource = r.img.img2
		r.right.Resource = r.img.img1
	} else {
		r.left.Resource = r.img.img1
		r.right.Resource = r.img.img2
	}

	r.left.Refresh()
	r.right.Refresh()
}
