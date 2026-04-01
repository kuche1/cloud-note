package app

import (
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// IMPROVE001: Marking some text and then pressing scroll bottom deletes some of the text
// IMPROVE001: Make this more pleasent to work with
func (self *App) SceneEditNote(
	previousText string,
	noteName string,
	viewingCachedCopy bool,
	cursorColumn int,
	cursorRow int,
) {
	editor := widget.NewMultiLineEntry()

	editor.Text = previousText
	editor.TextStyle.Monospace = true
	editor.Wrapping = fyne.TextWrapWord // TextWrapBreak
	editor.CursorColumn = cursorColumn
	editor.CursorRow = cursorRow
	// editor.Append("asd gfd hgf\nfdsfdsafdsaf")
	// editor.PlaceHolder = "Enter some text"

	if viewingCachedCopy {
		// IMPROVE001: Would be much better if instead we overwrite the TypedRune and TypedKey or
		// whatever they're called methods
		editor.OnChanged = func(idk string) {
			editor.Text = previousText
		}
	}

	cancel := widget.NewButton(
		"Cancel",

		func() {
			if editor.Text == previousText {
				self.SceneSelectNote()
				return
			}

			self.IntermissionYesNo(
				"Note content has changed.\nAre you sure you want to discard the new changes?",
				func() { self.SceneSelectNote() },
				func() {},
			)
		},
	)

	submit := widget.NewButton(
		"Submit",
		func() {
			self.IntermissionSubmitNewNoteContent(
				editor.Text,
				noteName,
				func() {
					previousText = editor.Text
					self.window.Focus(editor)
				},
			)
		},
	)

	if viewingCachedCopy {
		submit.Disable()
	}

	scrollToTop := widget.NewButton(
		"Jump Top",

		func() {
			editor.CursorColumn = 0
			editor.CursorRow = 0
			editor.TypedRune('\n')
			editor.CursorColumn = 0
			editor.CursorRow = 0
			editor.Refresh()

			self.window.Focus(editor)
		},
	)

	if viewingCachedCopy {
		scrollToTop.Disable()
	}

	scrollToBottom := widget.NewButton(
		"Jump bottom",

		func() {
			editor.CursorRow = math.MaxInt
			editor.Refresh()
			editor.TypedRune('\n')

			self.window.Focus(editor)

			//// This also works but is also hacky
			// editor.CursorRow = len(editor.Text)
			// editor.Refresh()

			//// This works but is hacky
			// editor.CursorColumn = 99999
			// editor.CursorRow = 99999
			// editor.Refresh()
		},
	)

	if viewingCachedCopy {
		scrollToBottom.Disable()
	}

	undo := widget.NewButton(
		"Undo",
		func() {
			editor.Undo()
			self.window.Focus(editor)
		},
	)

	if viewingCachedCopy {
		undo.Disable()
	}

	redo := widget.NewButton(
		"Redo",
		func() {
			editor.Redo()
			self.window.Focus(editor)
		},
	)

	if viewingCachedCopy {
		redo.Disable()
	}

	buttons := container.NewGridWithColumns(
		2,
		cancel,
		submit,
		undo,
		redo,
		scrollToTop,
		scrollToBottom,
	)

	containerTop := container.NewVBox(
		buttons,
	)

	if viewingCachedCopy {
		containerTop.Add(container.NewCenter(widget.NewLabel("Viewing Read-Only Cached Copy")))
	}

	container := container.NewBorder(
		containerTop,
		nil,
		nil,
		nil,
		editor,
	)

	self.window.SetContent(container)
	self.window.Focus(editor)
}
