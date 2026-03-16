package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.New()
	window := app.NewWindow("Simple Label")
	window.Resize(fyne.NewSize(400, 600))
	FirstScene(&window)
	window.ShowAndRun()
}
