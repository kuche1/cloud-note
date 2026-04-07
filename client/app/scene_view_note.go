package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (self *App) SceneViewNote(
	noteName string,
	noteContents string,
) {
	viewer := widget.NewLabel(noteContents)

	viewer.TextStyle.Monospace = true
	viewer.Wrapping = fyne.TextWrapWord // TextWrapBreak

	cancel := widget.NewButton(
		"Cancel",
		func() {
			self.SceneSelectNote()
		},
	)

	edit := widget.NewButton(
		"Edit",
		func() {
			self.SceneEditNote(
				noteName,
				noteContents,
				false,
				0,
				0,
			)
		},
	)

	buttons := container.NewGridWithColumns(
		2,
		cancel,
		edit,
	)

	containerTop := container.NewVBox(
		buttons,
	)

	container := container.NewBorder(
		containerTop,
		nil,
		nil,
		nil,
		container.NewVScroll(viewer),
	)

	self.window.SetContent(container)
}
