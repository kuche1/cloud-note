package app

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (self *App) IntermissionEditLine(
	contentCurrent string,
	contentOriginal string,
	existsOriginal bool,
	callbackWhenDone func(newLineContent string, deleteLine bool),
) {
	previousFyneContent := self.window.Content()

	editor := widget.NewEntry()
	editor.Text = contentCurrent
	editor.TextStyle.Monospace = true

	editor.OnSubmitted = func(newContent string) {
		self.window.SetContent(previousFyneContent)
		callbackWhenDone(newContent, false)
	}

	btnCancel := widget.NewButton(
		"Cancel",
		func() {
			// I don't think we need to ask the user is he wants to discard the
			// changes considering the fact that he is editing a single line
			self.window.SetContent(previousFyneContent)
			callbackWhenDone(contentCurrent, false)
		},
	)

	btnOk := widget.NewButton(
		"Ok",
		func() {
			self.window.SetContent(previousFyneContent)
			callbackWhenDone(editor.Text, false)
		},
	)

	btnRestoreOriginal := widget.NewButton(
		"Restore Original",
		func() {
			editor.Text = contentOriginal
			editor.Refresh()
			self.window.Focus(editor)
		},
	)
	if !existsOriginal {
		btnRestoreOriginal.Text += " >"
		btnRestoreOriginal.Disable()
	}

	// IMPROVE000: Ideally we would be also able to delete a line from the note edit scene
	btnDelete := widget.NewButton(
		"Delete",
		func() {
			self.window.SetContent(previousFyneContent)
			callbackWhenDone("", true)
		},
	)

	btnUndo := widget.NewButton(
		"Undo",
		func() {
			editor.Undo()
			self.window.Focus(editor)
		},
	)

	btnRedo := widget.NewButton(
		"Redo",
		func() {
			editor.Redo()
			self.window.Focus(editor)
		},
	)

	btnCursorLeft := widget.NewButton(
		"< Cursor",
		func() {
			editor.CursorColumn-- // not adjusting the row is OK since we're editing a single line
			editor.Refresh()
			self.window.Focus(editor)
		},
	)

	btnCursorRight := widget.NewButton(
		"Cursor >",
		func() {
			editor.CursorColumn++
			editor.Refresh()
			self.window.Focus(editor)
		},
	)

	containerTop := container.NewVBox(
		container.NewGridWithColumns(
			2,
			btnCancel, btnOk,
			btnRestoreOriginal, btnDelete,
			btnUndo, btnRedo,
			btnCursorLeft, btnCursorRight,
		),
	)

	self.window.SetContent(
		container.NewBorder(
			containerTop,
			nil,
			nil,
			nil,
			editor,
			// container.NewGridWrap(fyne.NewSize(2000, editorMinSizeY*2), editor),
			// container.New(layout.NewRowWrapLayout(), editor),
		),
	)
	self.window.Focus(editor)
}
