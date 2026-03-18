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

// IMPROVE: Maybe just have a `NetState` that contains both `conn` and `stream`
// And actually attempt to clean everything and just collect and print errorrs instead
// of panicing
func cancel(app *App, output *widget.TextGrid) {
	fyne.Do(func() {
		// IMPROVE: Only do so if the content has actually changed
		// (or maybe not, we'll see what architecture I'll go for)
		output.Append("Quitting GUI...")
	})

	fyne.Do(func() {
		app.Quit()
	})
}
