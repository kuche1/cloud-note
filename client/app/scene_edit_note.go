package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/kuche1/cloud-note/client/app/libnote"
)

// Eyeball it (not great)
const _ItemHeightOneLine float32 = 32.296875 // TODO: wow
const _ItemHeightTwoLines float32 = 47.5     // 45-50

// TODO: would be better if we instead got the label regular height and doubled it

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
	// noteContent := notecontent.NewNoteContent(noteContentStarting)
	note := libnote.NewNote(noteContentStarting)

	var editor *widget.List

	editor = widget.NewList(
		func() int {
			return note.Len()
		},

		func() fyne.CanvasObject {
			label := widget.NewLabel("")
			label.TextStyle.Monospace = true
			// label.Alignment = fyne.TextAlignCenter // horizontal alignment
			return label
		},

		func(index widget.ListItemID, canvas fyne.CanvasObject) {
			// TODO: I dont like this solution, only use it if we have to
			// if noteContent.IsEmpty() {
			// 	label.SetText("<Click to Edit>")
			// 	return
			// }

			repr, twoLines := note.LineStatusAndContent(index)

			label := canvas.(*widget.Label)
			label.SetText(repr)
			// fmt.Printf("DBG: height is %v\n", label.Size().Height)

			height := _ItemHeightOneLine
			if twoLines {
				height = _ItemHeightTwoLines
			}

			editor.SetItemHeight(index, height)

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

		contentCurrent, _ := note.LineContent(index)
		// we can forbid editing of notes that do not exist
		// if we want to

		contentOriginal, existsOriginal := note.LineContentOriginal(index)

		self.IntermissionEditLine(
			contentCurrent,
			contentOriginal,
			existsOriginal,

			func(newLineContent string, deleteLine bool) {
				defer editor.Refresh()

				if deleteLine {
					note.LineDelete(index)
					return
				}

				note.SetLineContent(index, newLineContent)
			},
		)
	}

	cancel := widget.NewButton(
		"Cancel",

		func() {
			if note.HasBeenChanged() {
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
			newContent := note.Content()

			self.IntermissionSubmitNewNoteContent(
				newContent,
				noteName,
				func() {
					// noteContent.SetHasNotBeenChanged() // TODO: this is now missing, better reset the whole scene
					// TODO: actually, there is ScrollTo (I think)
					editor.Refresh() // update any items that previously have been marked as outdated
				},
			)
		},
	)

	btnAddLineTop := widget.NewButton(
		"v Add Line v",
		func() {
			note.AddLineTop()
			editor.Refresh()
		},
	)

	btnAddLineBot := widget.NewButton(
		"^ Add Line ^",
		func() {
			note.AddLineBot()
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
