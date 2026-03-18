package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/kuche1/cloud-note/client/action"
	"github.com/kuche1/cloud-note/client/settings"
)

func (self *App) SceneSubmit(newText string, settings *settings.Settings) {
	output := widget.NewTextGrid()
	self.window.SetContent(output)

	go func() {
		err := action.ActionSetNoteContent(&self.window, output, newText, settings, self.app.Storage())
		if err != nil {
			self.ScenePanic(err.Error())
			return
		}

		fyne.Do(func() {
			self.SceneCancel()
		})
	}()
}
