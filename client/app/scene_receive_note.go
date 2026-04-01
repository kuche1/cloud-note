package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/kuche1/cloud-note/client/action"
	"github.com/kuche1/cloud-note/client/output"
)

func (self *App) SceneReceiveNote(noteName string) {
	output, outputWidget := output.NewOutputFyneAny()

	dialog := dialog.NewCustomWithoutButtons(
		"Receive Note",
		outputWidget,
		*self.window.FyneWindow,
	)

	dialog.Show()

	go func() {
		defer fyne.Do(func() { dialog.Dismiss() })

		data, err := action.ActionGetNoteContent(self.window, output, self.settings, noteName)
		if err != nil {
			self.ScenePanic(err.Error())
			return
		}

		fyne.Do(func() {
			self.SceneEditNote(
				string(data),
				noteName,
				false,
				0,
				0,
			)
		})
	}()
}
