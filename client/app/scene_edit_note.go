package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/kuche1/cloud-note/client/app/libnote"
)

const _ItemHeightOneLine float32 = 32.296875 // IMPROVE001: Ideally we would dynamically get this value HOWEVER this actually seems to be good enough for both desktop and phone so I'll keep it as is
const _ItemHeightTwoLines float32 = 47.5

// TODO: allow for inserting new lines in the middle of the note

func (self *App) SceneEditNote(
	noteName string,
	noteContentStarting string,
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
				noteName,
				newContent,
				func() {
					// IMPROVE000: We could use `editor.ScrollTo` if we wanted to be super fancy
					note.ConsiderContentUpdated()
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
		// widget.NewSeparator(),
	)

	containerBot := container.NewVBox(
		// widget.NewSeparator(),
		btnAddLineBot,
	)

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
