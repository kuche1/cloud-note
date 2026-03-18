package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/kuche1/cloud-note/client/action"
	"github.com/kuche1/cloud-note/client/settings"
)

// IMPROVE000: Ideally we would only send the new note if the content has actually changed
func (self *App) SceneSubmit(newText string, settings *settings.Settings, noteName string) {
	output := widget.NewTextGrid()
	self.window.SetContent(output)

	go func() {
		err := action.ActionSetNoteContent(self.window.FyneWindow, output, newText, settings, noteName)
		if err != nil {
			self.ScenePanic(err.Error())
			return
		}

		fyne.Do(func() {
			self.SceneCancel()
		})
	}()
}
