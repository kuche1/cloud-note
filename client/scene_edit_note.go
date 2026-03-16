package client

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (self *App) SceneEditNote() {
	editor := widget.NewMultiLineEntry()
	editor.PlaceHolder = "Enter some text"
	// editor.Append("asd gfd hgf\nfdsfdsafdsaf")

	cancel := widget.NewButton("Cancel", func() { self.SceneCancel() })
	submit := widget.NewButton("Submit", func() { self.SceneSubmit() })
	buttons := container.NewGridWithColumns(
		2,
		cancel,
		submit,
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
