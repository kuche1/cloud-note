package app

import (
	"fyne.io/fyne/v2"
	"github.com/kuche1/cloud-note/client/action"
	"github.com/kuche1/cloud-note/client/output"
	"github.com/kuche1/cloud-note/client/settings"
)

func (self *App) SceneReceiveNote(settings *settings.Settings, noteName string) {
	self.window.ShowDialogOutput(
		"Receive Note",

		func(output output.Output) {
			data, err := action.ActionGetNoteContent(self.window, output, settings, noteName)
			if err != nil {
				self.ScenePanic(err.Error())
				return
			}

			fyne.Do(func() {
				self.SceneEditNote(string(data), settings, noteName)
			})
		},
	)
}
