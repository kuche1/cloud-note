package app

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/kuche1/cloud-note/client/app/theme"
	"github.com/kuche1/cloud-note/client/settings"
	"github.com/kuche1/cloud-note/client/window"
)

type App struct {
	app      fyne.App
	window   *window.Window
	settings *settings.Settings
}

func RunApp() {
	app := app.NewWithID("cloud-note")
	app.Settings().SetTheme(theme.NewTheme())

	fyneWindow := app.NewWindow("Cloud Note")
	fyneWindow.Resize(fyne.NewSize(400, 600))

	windo := window.Window{}.NewFromFyneWindow(&fyneWindow)

	self := App{
		app:      app,
		window:   windo,
		settings: settings.Settings{}.NewFromDefaults(app.Storage().RootURI().Path()),
	}

	self.FirstScene()
	self.window.ShowAndRun()
}

// Must not rely on `self.ScenePanic`
func (self *App) Quit() {
	self.app.Quit()
	// NOTE: This causes the GUI to freeze on mobile,
	// but it does not exit the app, so we have to
	// call `os.Exit` manually

	os.Exit(0)
}
