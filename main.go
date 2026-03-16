package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.New()
	window := app.NewWindow("Simple Label")
	window.Resize(fyne.NewSize(400, 600))

	button := widget.NewButton("Start", func() { EditNote(&window) })

	window.SetContent(button)
	window.ShowAndRun()
}
