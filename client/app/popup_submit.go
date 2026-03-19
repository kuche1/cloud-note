package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/kuche1/cloud-note/client/action"
	"github.com/kuche1/cloud-note/client/output"
	"github.com/kuche1/cloud-note/client/settings"
)

// IMPROVE000: Ideally we would only send the new note if the content has actually changed
func (self *App) PopupSubmitNewNoteContent(newText string, settings *settings.Settings, noteName string) {
	output, textGrid := output.NewOutputFyneTextGrid()

	popup := dialog.NewCustomWithoutButtons(
		"Submit Note",
		textGrid,
		*self.window.FyneWindow,
	)

	popup.Show()

	go func() {
		defer fyne.Do(popup.Dismiss)

		err := action.ActionSetNoteContent(self.window, output, newText, settings, noteName)
		if err != nil {
			self.ScenePanic(err.Error())
			return
		}
	}()
}
