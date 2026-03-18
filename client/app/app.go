package app

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type App struct {
	app    fyne.App
	window fyne.Window
}

func RunApp() {
	app := app.NewWithID("could-note")
	window := app.NewWindow("Cloud Note")
	window.Resize(fyne.NewSize(400, 600))

	self := App{
		app:    app,
		window: window,
	}

	self.FirstScene()

	window.ShowAndRun()
}

// Must not rely on `self.ScenePanic`
func (self *App) Quit() {
	self.app.Quit()
	// NOTE: This causes the GUI to freeze on mobile,
	// but it does not exit the app, so we have to
	// call `os.Exit` manually

	os.Exit(0)
}
