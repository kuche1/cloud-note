package app

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/kuche1/cloud-note/client/settings"
)

func (self *App) SceneLoadSettings() {
	output := widget.NewTextGrid()
	self.window.SetContent(output)

	go loadSettings(self, output)
}

func loadSettings(app *App, output *widget.TextGrid) {
	fyne.Do(func() { output.Append("Loading settings...") })

	settings, err :=
		settings.Settings{}.NewFromPersistentStorage(app.app.Storage().RootURI().Path())
	if err != nil {
		app.ScenePanic(fmt.Sprintf("Could not load settings:\n%v", err))
		return
	}

	fyne.Do(func() { output.Append("Done!") })

	fyne.Do(func() { app.SceneReceiveNote(settings) })
}
