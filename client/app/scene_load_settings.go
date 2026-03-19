package app

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/kuche1/cloud-note/client/output"
	"github.com/kuche1/cloud-note/client/settings"
)

func (self *App) SceneLoadSettings() {
	output, textGrid := output.NewOutputFyneTextGrid()

	dialog := dialog.NewCustomWithoutButtons(
		"Load Settings",
		textGrid,
		*self.window.FyneWindow,
	)

	dialog.Show()

	go func() {
		defer fyne.Do(func() { dialog.Dismiss() })
		// Won't be perfect but at least I'll know that it will be dismissed 100%

		output.Println("Loading settings...")

		settings, err :=
			settings.Settings{}.NewFromPersistentStorage(self.app.Storage().RootURI().Path())
		if err != nil {
			self.ScenePanic(fmt.Sprintf("Could not load settings:\n%v", err))
			return
		}

		output.Println("Done!")

		fyne.Do(func() {
			self.SceneSelectNote(settings)
		})
	}()
}
