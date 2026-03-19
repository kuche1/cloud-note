package app

import (
	"fmt"

	"fyne.io/fyne/v2"
	"github.com/kuche1/cloud-note/client/output"
	"github.com/kuche1/cloud-note/client/settings"
)

func (self *App) SceneLoadSettings() {
	self.window.ShowDialogOutput(
		"Load Settings",

		func(output output.Output) {
			output.Println("Loading settings...")

			settings, err :=
				settings.Settings{}.NewFromPersistentStorage(self.app.Storage().RootURI().Path())
			if err != nil {
				self.ScenePanic(fmt.Sprintf("Could not load settings:\n%v", err))
				return
			}

			output.Println("Done!")

			fyne.Do(func() { self.SceneSelectNote(settings) })
		},
	)
}
