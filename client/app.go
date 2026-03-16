package client

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type App struct {
	app    fyne.App
	window fyne.Window
}

func RunApp() {
	app := app.New()
	window := app.NewWindow("Cloud Note")
	window.Resize(fyne.NewSize(400, 600))

	self := App{
		app:    app,
		window: window,
	}

	self.FirstScene()

	window.ShowAndRun()
}
