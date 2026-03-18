package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/kuche1/cloud-note/client/action"
	"github.com/kuche1/cloud-note/client/config"
	"github.com/kuche1/cloud-note/client/settings"
)

func (self *App) SceneReceiveNote(settings *settings.Settings) {
	self.window.ShowDialogOutput(
		"Receive Note",

		func(output *widget.TextGrid) {
			data, err := action.ActionGetNoteContent(self.window.FyneWindow, output, settings, config.NoteName)
			if err != nil {
				self.ScenePanic(err.Error())
				return
			}

			fyne.Do(func() {
				self.SceneEditNote(string(data), settings)
			})
		},
	)
}
