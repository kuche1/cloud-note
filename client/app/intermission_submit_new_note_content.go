package app

import (
	"fmt"

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
			fyne.Do(func() {
				self.IntermissionInfo(
					fmt.Sprintf("Could not set note content:\n%v", err),
					func() { self.window.SetContent(previousFyneContent) },
				)
			})
			return
		}

		fyne.Do(func() { self.window.SetContent(previousFyneContent) })
	}()
}
