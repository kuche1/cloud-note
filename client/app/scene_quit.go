package app

import (
	"fyne.io/fyne/v2"
	"github.com/kuche1/cloud-note/client/output"
)

func (self *App) SceneQuit() {
	output, outputWidget := output.NewOutputFyneAny()
	self.window.SetContent(outputWidget)

	go cancel(self, output)
}

func cancel(app *App, output output.Output) {
	output.Println("Quitting GUI...")

	fyne.Do(func() {
		app.Quit()
	})
}
