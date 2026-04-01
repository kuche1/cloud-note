package app

import (
	"fmt"

	"fyne.io/fyne/v2"
	"github.com/kuche1/cloud-note/client/action"
	"github.com/kuche1/cloud-note/client/output"
)

// IMPROVE000: Ideally we would only send the new note if the content has actually changed
// BUT if we are to do that we need to make that a setting as to let paranoid users send
// the same note a billion time (perhaps add a setting for this, and when the button is clicked,
// perform a check weather the content has changed or not)
func (self *App) IntermissionSubmitNewNoteContent(
	newText string,
	noteName string,
	callbackSuccess func(),
) {
	previousContent := self.window.Content()

	output, outputWidget := output.NewOutputFyneAny()
	self.window.SetContent(outputWidget)

	go func() {
		message := "Upload Successful"

		err := action.ActionSetNoteContent(self.window, output, newText, self.settings, noteName)
		if err != nil {
			message = fmt.Sprintf("Could not set note content:\n%v", err)
		}

		fyne.Do(func() {
			self.IntermissionInfo(
				message,
				func() {
					self.window.SetContent(previousContent)
					callbackSuccess()
				},
			)
		})
	}()
}
