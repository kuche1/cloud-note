package client

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func (self *App) SceneConnectToServer() {
	output := widget.NewTextGrid()
	self.window.SetContent(output)

	go connectToServer(self, output)
}

func connectToServer(app *App, output *widget.TextGrid) {
	fyne.Do(func() {
		output.Append("Connecting to server...")
	})

	time.Sleep(1 * time.Second)

	fyne.Do(func() {
		output.Append("Downloading data...")
	})

	time.Sleep(1 * time.Second)

	fyne.Do(func() {
		output.Append("Done!")
		app.SceneEditNote()
	})
}
