package app

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/kuche1/cloud-note/client/app/notecontent"
)

func (self *App) IntermissionEditLine(
	noteLine *notecontent.NoteLine,
	callbackWhenDone func(deleteLine bool),
) {
	previousFyneContent := self.window.Content()

	editor := widget.NewEntry()
	editor.Text = noteLine.Content()
	// editor.MultiLine = true // wrapping does not work if it is a single line entry
	editor.TextStyle.Monospace = true
	// editor.Wrapping = fyne.TextWrapBreak // fyne.TextWrapWord
	// editorMinSizeY := editor.MinSize().Height

	editor.OnSubmitted = func(newContent string) {
		noteLine.SetContent(newContent)
		self.window.SetContent(previousFyneContent)
		callbackWhenDone(false)
	}

	btnCancel := widget.NewButton(
		"Cancel",
		func() {
			// I don't think we need to ask the user is he wants to discard the
			// changes considering the fact that he is editing a single line
			self.window.SetContent(previousFyneContent)
			callbackWhenDone(false)
		},
	)

	btnOk := widget.NewButton(
		"Ok",
		func() {
			noteLine.SetContent(editor.Text)
			self.window.SetContent(previousFyneContent)
			callbackWhenDone(false)
		},
	)

	// TODO: I hate this, there needs to be a way to delete entries from the note edit scene
	btnDelete := widget.NewButton(
		"Delete",
		func() {
			self.window.SetContent(previousFyneContent)
			callbackWhenDone(true)
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
