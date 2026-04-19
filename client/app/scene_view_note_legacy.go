package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// It's not that the code here is bad, it's just that the new editor already covers read-only cases
func (self *App) SceneViewNoteLegacy(
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

	editLegacy := widget.NewButton(
		"Edit",
		func() {
			self.SceneEditNoteLegacy(
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
		editLegacy,
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
