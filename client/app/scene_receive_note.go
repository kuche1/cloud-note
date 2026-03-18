package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/kuche1/cloud-note/client/action"
	"github.com/kuche1/cloud-note/client/settings"
)

func (self *App) SceneReceiveNote(settings *settings.Settings) {
	output := widget.NewTextGrid()

	dialog := dialog.NewCustomWithoutButtons(
		"Receving Note",
		output,
		*self.window.FyneWindow,
	)

	dialog.Show()

	go receiveNote(self, settings, output, dialog)
}

func receiveNote(app *App, settings *settings.Settings, output *widget.TextGrid, dialog *dialog.CustomDialog) {
	data, err := action.ActionGetNoteContent(app.window.FyneWindow, output, settings)
	if err != nil {
		dialog.Dismiss()
		app.ScenePanic(err.Error())
		return
	}

	fyne.Do(func() {
		dialog.Dismiss()
		app.SceneEditNote(string(data), settings)
	})
}
