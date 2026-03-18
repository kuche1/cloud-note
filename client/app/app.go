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

	windo := app.NewWindow("Cloud Note")
	windo.Resize(fyne.NewSize(400, 600))

	// TODO: Make something like that
	// settings, err := settings.Settings{}.NewFromPersistentStorage(app.Storage().RootURI().Path())
	// if err != nil {
	// 	windo.ScenePanic(fmt.Sprintf("Could not initalise settings:\n%v", err))
	// 	self.window.ShowAndRun()
	//  quit(1)
	// }

	self := App{
		app:    app,
		window: window.Window{}.NewFromFyneWindow(&windo),
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
