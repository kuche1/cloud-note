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
// [TODO: fix in progress] IMPROVE000: For some reason after we switch back to the previous scene the window expands
// dramatically (vertically) on desktop; and this happens exclusively after having pressed the
// submit button
func (self *App) SceneSubmitNewNoteContent(
	newText string,
	noteName string,
	cursorColumn int,
	cursorRow int,
	callbackSuccess func(),
) {
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

					self.SceneEditNote(
						newText,
						noteName,
						false,
						cursorColumn,
						cursorRow,
					)

					callbackSuccess()

				},
			)
		})
	}()
}
