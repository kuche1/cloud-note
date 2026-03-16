package client

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func SceneConnectToServer(window *fyne.Window) {
	output := widget.NewTextGrid()
	(*window).SetContent(output)
	go connectToServer(window, output)
}

func connectToServer(window *fyne.Window, output *widget.TextGrid) {
	fyne.Do(func() {
		output.Append("Connecting to server...")
	})

	time.Sleep(1 * time.Second)

	fyne.Do(func() {
		output.Append("Connected!")
		SceneEditNote(window)
	})
}
