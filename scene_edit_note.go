package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func SceneEditNote(window *fyne.Window) {
	cancel := widget.NewButton("Cancel", func() { SceneCancel(window) })
	submit := widget.NewButton("Submit", func() { SceneSubmit(window) })
	buttons := container.NewGridWithColumns(2,
		cancel,
		submit,
	)

	editor := widget.NewMultiLineEntry()
	editor.PlaceHolder = "Enter some text"
	// editor.Append("asd gfd hgf\nfdsfdsafdsaf")

	container := container.NewBorder(
		buttons,
		nil,
		nil,
		nil,
		editor,
	)

	(*window).SetContent(container)
}
