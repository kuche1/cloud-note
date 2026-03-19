package app

import (
	"fyne.io/fyne/v2"
	"github.com/kuche1/cloud-note/client/action"
	"github.com/kuche1/cloud-note/client/output"
	"github.com/kuche1/cloud-note/client/settings"
)

// IMPROVE000: Ideally we would only send the new note if the content has actually changed
func (self *App) IntermissionSubmitNewNoteContent(newText string, settings *settings.Settings, noteName string) {
	previousFyneContent := self.window.Content()

	output, textGrid := output.NewOutputFyneTextGrid()
	self.window.SetContent(textGrid)

	go func() {
		err := action.ActionSetNoteContent(self.window, output, newText, settings, noteName)
		if err != nil {
			// TODO: Show popup and exit
			self.ScenePanic(err.Error())
			return
		}

		fyne.Do(func() {
			self.window.SetContent(previousFyneContent)
		})
	}()
}
