package app

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/kuche1/cloud-note/client/settings"
)

func (self *App) SceneLoadSettings() {
	self.window.ShowDialogOutput(
		"Load Settings",

		func(output *widget.TextGrid) {
			fyne.Do(func() { output.Append("Loading settings...") })

			settings, err :=
				settings.Settings{}.NewFromPersistentStorage(self.app.Storage().RootURI().Path())
			if err != nil {
				self.ScenePanic(fmt.Sprintf("Could not load settings:\n%v", err))
				return
			}

			fyne.Do(func() { output.Append("Done!") })

			fyne.Do(func() { self.SceneReceiveNote(settings) })
		},
	)
}
