package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func (self *App) SceneCancel() {
	output := widget.NewTextGrid()
	self.window.SetContent(output)

	go cancel(self, output)
}

func cancel(app *App, output *widget.TextGrid) {
	fyne.Do(func() {
		output.Append("Quitting GUI...")
	})

	fyne.Do(func() {
		app.Quit()
	})
}
