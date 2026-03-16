package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func SceneConnectToServer(window *fyne.Window) {
	info := widget.NewLabel("Connecting to server...")
	(*window).SetContent(info)
}
