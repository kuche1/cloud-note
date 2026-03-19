package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/kuche1/cloud-note/client/action"
	"github.com/kuche1/cloud-note/client/output"
	"github.com/kuche1/cloud-note/client/settings"
)

func (self *App) SceneReceiveNote(settings *settings.Settings, noteName string) {
	output, textGrid := output.NewOutputFyneTextGrid()

	dialog := dialog.NewCustomWithoutButtons(
		"Receive Note",
		textGrid,
		*self.window.FyneWindow,
	)

	dialog.Show()

	go func() {
		defer fyne.Do(func() { dialog.Dismiss() })

		data, err := action.ActionGetNoteContent(self.window, output, settings, noteName)
		if err != nil {
			self.ScenePanic(err.Error())
			return
		}

		fyne.Do(func() {
			self.SceneEditNote(string(data), settings, noteName)
		})
	}()
}
