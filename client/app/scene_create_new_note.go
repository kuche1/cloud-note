package app

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/kuche1/cloud-note/client/action"
	"github.com/kuche1/cloud-note/client/output"
	"github.com/kuche1/cloud-note/client/settings"
)

func (self *App) SceneCreateNewNote(settings *settings.Settings) {
	entry := widget.NewEntry()
	entry.PlaceHolder = "New Note Name"

	var ok *widget.Button
	ok = widget.NewButton(
		"Ok",
		func() {
			// ok.Disable()
			output, outputWidget := output.NewOutputFyneAny()
			self.window.SetContent(outputWidget)

			go func() {
				newNoteName := entry.Text

				err := action.ActionCreateNewNote(newNoteName, self.window, output, settings)
				if err != nil {
					self.ScenePanic(fmt.Sprintf("Could not create new note:\n%v", err))
					return
				}

				fyne.Do(func() {
					self.SceneEditNote("", settings, newNoteName, false)
				})
			}()
		},
	)

	cancel := widget.NewButton(
		"Cancel",
		func() {
			self.SceneSelectNote(settings)
		},
	)

	container := container.NewVBox(
		entry,
		ok,
		cancel,
	)

	self.window.SetContent(container)
	self.window.Focus(entry)
}
