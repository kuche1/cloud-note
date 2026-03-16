package client

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func Main() {
	app := app.New()
	window := app.NewWindow("Simple Label")
	window.Resize(fyne.NewSize(400, 600))
	FirstScene(&window)
	window.ShowAndRun()
}
