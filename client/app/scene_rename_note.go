package app

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/kuche1/cloud-note/client/output"
)

func (self *App) SceneRenameNote(oldName string, callbackSuccess func(newName string)) {
	var btnOk *widget.Button

	lblInfo := widget.NewLabel("Enter new name:")

	entInput := widget.NewEntry()
	entInput.Text = oldName

	entInput.OnSubmitted = func(string) {
		btnOk.OnTapped()
	}

	btnOk = widget.NewButton(
		"Ok",
		func() {
			newName := entInput.Text
			renameNotePart2(oldName, newName, func() { callbackSuccess(newName) }, self)
		},
	)

	btnCacel := widget.NewButton(
		"Cancel",
		func() {
			self.SceneSelectNote()
		},
	)

	self.window.SetContent(
		container.NewVBox(
			lblInfo,
			entInput,
			container.NewGridWithColumns(
				2,
				btnCacel,
				btnOk,
			),
		),
	)

	self.window.Focus(entInput)
}

func renameNotePart2(oldName string, newName string, callbackSuccess func(), app *App) {
	output, outputWidget := output.NewOutputFyneAny()
	app.window.SetContent(outputWidget)

	go renameNotePart3(oldName, newName, callbackSuccess, app, output)
}

func renameNotePart3(oldName string, newName string, callbackSuccess func(), app *App, output output.Output) {
	refusal, err := app.net.ActionRenameNote(oldName, newName, app.window, output, app.settings)

	if err != nil {
		fyne.Do(func() {
			app.ScenePanic(err.Error())
		})
		return
	}

	if len(refusal) > 0 {
		fyne.Do(func() {
			app.IntermissionInfo(
				fmt.Sprintf(
					"Could not rename note\n%v\nto\n%v\n\nReason:\n%v",
					oldName,
					newName,
					refusal,
				),
				func() { app.SceneSelectNote() },
			)
		})
		return
	}

	fyne.Do(func() {
		callbackSuccess()
		app.SceneSelectNote()
	})
}
