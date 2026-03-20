package app

import (
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/kuche1/cloud-note/client/settings"
)

// TODO: Make this more pleasent to work with
func (self *App) SceneEditNote(previousText string, settings *settings.Settings, noteName string) {
	editor := widget.NewMultiLineEntry()
	editor.Text = previousText
	editor.TextStyle.Monospace = true
	editor.Wrapping = fyne.TextWrapBreak // TextWrapWord
	// editor.Append("asd gfd hgf\nfdsfdsafdsaf")
	// editor.PlaceHolder = "Enter some text"

	cancel := widget.NewButton(
		"Cancel",

		func() {
			if editor.Text == previousText {
				self.SceneSelectNote(settings)
				return
			}

			self.IntermissionYesNo(
				"Note content has changed.\nAre you sure you want to discard the new changes?",
				func() { self.SceneSelectNote(settings) },
				func() {},
			)
		},
	)

	submit := widget.NewButton(
		"Submit",
		func() {
			self.IntermissionSubmitNewNoteContent(
				editor.Text,
				settings,
				noteName,
				func() { previousText = editor.Text },
			)
		},
	)

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

	undo := widget.NewButton(
		"Undo",
		func() {
			editor.Undo()
			self.window.Focus(editor)
		},
	)

	redo := widget.NewButton(
		"Redo",
		func() {
			editor.Redo()
			self.window.Focus(editor)
		},
	)

	buttons := container.NewGridWithColumns(
		2,
		cancel,
		submit,
		undo,
		redo,
		scrollToTop,
		scrollToBottom,
	)

	container := container.NewBorder(
		buttons,
		nil,
		nil,
		nil,
		editor,
	)

	self.window.SetContent(container)
}
