package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func SceneMainMenu(window *fyne.Window) {
	button := widget.NewButton("Start", func() { SceneEditNote(window) })
	(*window).SetContent(button)
}
