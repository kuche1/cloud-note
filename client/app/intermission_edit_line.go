package app

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (self *App) IntermissionEditLine(
	lineContent string,
	callbackWhenDone func(newLineContent string, deleteLine bool),
) {
	previousFyneContent := self.window.Content()

	editor := widget.NewEntry()
	editor.Text = lineContent
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
			callbackWhenDone(lineContent, false)
		},
	)

	btnOk := widget.NewButton(
		"Ok",
		func() {
			self.window.SetContent(previousFyneContent)
			callbackWhenDone(editor.Text, false)
		},
	)

	// TODO: I hate this, there needs to be a way to delete entries from the note edit scene
	btnDelete := widget.NewButton(
		"Delete",
		func() {
			self.window.SetContent(previousFyneContent)
			callbackWhenDone("", true)
		},
	)

	btnUndo := widget.NewButton(
		"Undo",
		func() { editor.Undo() },
	)

	btnRedo := widget.NewButton(
		"Redo",
		func() { editor.Redo() },
	)

	containerTop := container.NewVBox(
		container.NewGridWithColumns(
			2,
			btnCancel,
			btnOk,
		),
		btnDelete,
		container.NewGridWithColumns(
			2,
			btnUndo,
			btnRedo,
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
