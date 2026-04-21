package app

import (
	"fmt"

	"fyne.io/fyne/v2"
	"github.com/kuche1/cloud-note/client/output"
)

func (self *App) SceneDeleteNote(noteName string) {
	output, outputWidget := output.NewOutputFyneAny()
	self.window.SetContent(outputWidget)

	self.IntermissionYesNo(
		fmt.Sprintf("Are you sure you want to delete note:\n%v", noteName),
		func() {
			go deleteNote(self, noteName, output)
		},
		func() {
			self.SceneSelectNote()
		},
	)
}

func deleteNote(app *App, noteName string, output output.Output) {
	err := app.net.ActionDeleteExistingNote(noteName, app.window, output, app.settings)
	if err != nil {
		fyne.Do(func() {
			app.IntermissionInfo(
				fmt.Sprintf("Could not delete note:\n%v\n\nReason:\n%v", noteName, err),
				func() { app.SceneSelectNote() },
			)
		})
		return
	}

	fyne.Do(func() { app.SceneSelectNote() })
}
