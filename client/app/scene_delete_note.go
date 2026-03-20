package app

import (
	"fmt"

	"fyne.io/fyne/v2"
	"github.com/kuche1/cloud-note/client/action"
	"github.com/kuche1/cloud-note/client/output"
	"github.com/kuche1/cloud-note/client/settings"
)

func (self *App) SceneDeleteNote(settings *settings.Settings, noteName string) {
	output, outputWidget := output.NewOutputFyneAny()
	self.window.SetContent(outputWidget)

	self.IntermissionYesNo(
		fmt.Sprintf("Are you sure you want to delete note:\n%v", noteName),
		func() {
			go deleteNote(self, noteName, output, settings)
		},
		func() {
			self.SceneSelectNote(settings)
		},
	)
}

func deleteNote(app *App, noteName string, output output.Output, settings *settings.Settings) {
	err := action.ActionDeleteExistingNote(noteName, app.window, output, settings)
	if err != nil {
		fyne.Do(func() {
			app.IntermissionInfo(
				fmt.Sprintf("Could not delete note:\n%v\n\nReason:\n%v", noteName, err),
				func() { app.SceneSelectNote(settings) },
			)
		})
		return
	}

	fyne.Do(func() { app.SceneSelectNote(settings) })
}
