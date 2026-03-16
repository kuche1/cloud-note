package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Simple Label")

	label := widget.NewLabel("Hello, Fyne!")

	w.SetContent(label)
	w.ShowAndRun()
}
