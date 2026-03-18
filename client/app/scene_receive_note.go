package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/kuche1/cloud-note/client/action"
	"github.com/kuche1/cloud-note/client/settings"
)

func (self *App) SceneReceiveNote(settings *settings.Settings) {
	output := widget.NewTextGrid()
	self.window.SetContent(output)

	go receiveNote(self, output, settings)
}

func receiveNote(app *App, output *widget.TextGrid, settings *settings.Settings) {
	data, err := action.ActionGetNoteContent(app.window.FyneWindow, output, settings, app.app.Storage())
	if err != nil {
		app.ScenePanic(err.Error())
		return
	}

	fyne.Do(func() { app.SceneEditNote(string(data), settings) })
}
