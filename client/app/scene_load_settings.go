package app

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/kuche1/cloud-note/client/output"
)

func (self *App) SceneLoadSettings() {
	output, outputWidget := output.NewOutputFyneAny()

	dialog := dialog.NewCustomWithoutButtons(
		"Load Settings",
		outputWidget,
		*self.window.FyneWindow,
	)

	dialog.Show()

	go func() {
		output.Println("Loading settings...")

		dialog.Dismiss()

		err :=
			self.settings.LoadFromPersistentStorage()
		if err != nil {
			self.ScenePanic(fmt.Sprintf("Could not load settings:\n%v", err))
			return
		}

		output.Println("Done!")

		fyne.Do(func() {
			self.settings.SceneInputEssential(
				self.window,
				func() { self.ScenePing() },
			)
		})
	}()
}
