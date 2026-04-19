package app

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/kuche1/cloud-note/client/app/notecontent"
)

func (self *App) IntermissionEditLine(noteLine *notecontent.NoteLine, callbackWhenDone func()) {
	previousFyneContent := self.window.Content()

	editor := widget.NewEntry()
	editor.Text = noteLine.Content()
	// editor.MultiLine = true // wrapping does not work if it is a single line entry
	editor.TextStyle.Monospace = true
	// editor.Wrapping = fyne.TextWrapBreak // fyne.TextWrapWord
	// editorMinSizeY := editor.MinSize().Height

	btnCancel := widget.NewButton(
		"Cancel",
		func() {
			// I don't think we need to ask the user is he wants to discard the
			// changes considering the fact that he is editing a single line
			// // TODO: actually it might be a good idea to add it back
			// TODO?: add undo and redo
			self.window.SetContent(previousFyneContent)
			callbackWhenDone()
		},
	)

	btnOk := widget.NewButton(
		"Ok",
		func() {
			noteLine.SetContent(editor.Text)
			self.window.SetContent(previousFyneContent)
			callbackWhenDone()
		},
	)

	self.window.SetContent(
		container.NewBorder(
			container.NewVBox(
				container.NewGridWithColumns(
					2,
					btnCancel,
					btnOk,
				),
			),
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
