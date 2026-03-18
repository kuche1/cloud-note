package app

import (
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/kuche1/cloud-note/client/settings"
)

func (self *App) SceneEditNote(previousText string, settings *settings.Settings, noteName string) {
	editor := widget.NewMultiLineEntry()
	editor.Text = previousText
	editor.TextStyle.Monospace = true
	editor.Wrapping = fyne.TextWrapBreak // TextWrapWord
	// editor.Append("asd gfd hgf\nfdsfdsafdsaf")
	// editor.PlaceHolder = "Enter some text"

	cancel := widget.NewButton(
		"Cancel",
		func() { self.SceneCancel() },
	)
	submit := widget.NewButton(
		"Submit",
		func() { self.SceneSubmit(editor.Text, settings, noteName) },
	)

	scrollToTop := widget.NewButton(
		"Scroll to Top",
		func() {
			editor.CursorColumn = 0
			editor.CursorRow = 0
			editor.Refresh()
		},
	)
	scrollToBottom := widget.NewButton(
		"Scroll to Bottom",
		func() {
			editor.CursorRow = math.MaxInt
			editor.Refresh()

			//// This also works but is also hacky
			// editor.CursorRow = len(editor.Text)
			// editor.Refresh()

			//// This works but is hacky
			// editor.CursorColumn = 99999
			// editor.CursorRow = 99999
			// editor.Refresh()
		},
	)

	buttons := container.NewGridWithColumns(
		2,
		cancel,
		submit,
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
