package app

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/kuche1/cloud-note/client/window"
)

type App struct {
	app    fyne.App
	window *window.Window
}

func RunApp() {
	app := app.NewWithID("cloud-note")

	fyneWindow := app.NewWindow("Cloud Note")
	fyneWindow.Resize(fyne.NewSize(400, 600))

	windo := window.Window{}.NewFromFyneWindow(&fyneWindow)

	// IMPROVE000: Make something like that
	// settings, err := settings.Settings{}.NewFromPersistentStorage(app.Storage().RootURI().Path())
	// if err != nil {
	// 	windo.ScenePanic(fmt.Sprintf("Could not initalise settings:\n%v", err))
	// 	self.window.ShowAndRun()
	//  quit(1)
	// }

	self := App{
		app:    app,
		window: windo,
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
