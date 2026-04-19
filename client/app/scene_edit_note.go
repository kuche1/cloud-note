package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/kuche1/cloud-note/client/app/notecontent"
)

// TODO: I really want to simplify the logic, I dont want to risk any data corruption

// TODO: there is a bug where if you delete ALL notes, submit, then press cancel
// you get a message "are you sure you want to discard"

// TODO: I would like to put this scene (and maybe all others) in their own folders
// if they need to change to another scene, they can take the given scene as argument OR the `*App` methods can
// be spread across different folder (if this works, but I dont think it will)
func (self *App) SceneEditNote(
	noteName string,
	noteContentStarting string, // TODO: I dont like that we're keeping this in RAM (or is it being optimised ????)
) {
	noteContent := notecontent.NewNoteContent(noteContentStarting)

	editor := widget.NewList(
		func() int {
			return noteContent.Len()
		},

		func() fyne.CanvasObject {
			label := widget.NewLabel("")
			label.TextStyle.Monospace = true
			return label
		},

		func(index widget.ListItemID, canvas fyne.CanvasObject) {
			label := canvas.(*widget.Label)

			if noteContent.IsEmpty() {
				label.SetText("<Click to Edit>")
				return
			}

			line := noteContent.Line(index)
			label.SetText(line.ContentWithStatus())

			//// this cannot be relied upon since some of the labels will be re-used
			// if line.contentHasBeenChanged {
			// 	label.Importance = widget.HighImportance
			// } else {
			// 	label.Importance = widget.MediumImportance
			// }
		},
	)

	editor.OnSelected = func(index widget.ListItemID) {
		editor.UnselectAll()

		self.IntermissionEditLine(
			noteContent.Line(index).Content(),

			func(newLineContent string, deleteLine bool) {
				defer editor.Refresh()

				if deleteLine {
					// TODO: maybe instead mark it as deleted and add a `-` symbol
					// OR maybe just always show the old line with that `-` symbol
					noteContent.Delete(index)
					return
				}

				noteContent.Line(index).SetContent(newLineContent)
			},
		)
	}

	cancel := widget.NewButton(
		"Cancel",

		func() {
			if noteContent.HasBeenChanged() {
				self.IntermissionYesNo(
					"Note content has changed.\nAre you sure you want to discard the new changes?",
					func() { self.SceneSelectNote() },
					func() {},
				)
				return
			}

			self.SceneSelectNote()
		},
	)

	submit := widget.NewButton(
		"Submit",

		func() {
			newContent, err := noteContent.AsString()
			if err != nil {
				self.ScenePanic(err.Error())
				return
			}

			self.IntermissionSubmitNewNoteContent(
				newContent,
				noteName,
				func() {
					noteContent.SetHasNotBeenChanged()
					editor.Refresh() // update any items that previously have been marked as outdated
				},
			)
		},
	)

	btnAddLineTop := widget.NewButton(
		"v Add Line v",
		func() {
			noteContent.AddLineTop()
			editor.Refresh()
		},
	)

	btnAddLineBot := widget.NewButton(
		"^ Add Line ^",
		func() {
			noteContent.AddLineBot()
			editor.Refresh()
		},
	)

	containerButtons := container.NewGridWithColumns(
		2,
		cancel,
		submit,
	)

	containerTop := container.NewVBox(
		containerButtons,
		btnAddLineTop,
	)

	containerBot := btnAddLineBot

	self.window.SetContent(
		container.NewBorder(
			containerTop,
			containerBot,
			nil,
			nil,
			editor,
		),
	)
	// self.window.Focus(editor)
}
