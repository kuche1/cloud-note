package app

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/kuche1/cloud-note/client/action"
	"github.com/kuche1/cloud-note/client/output"
)

func (self *App) ScenePing() {
	output, outputWidget := output.NewOutputFyneAny()
	self.window.SetContent(outputWidget)

	go func() {
		err := action.ActionPing(self.window, output, self.settings)
		if err != nil {
			pingFailed(self, err)
			return
		}

		fyne.Do(func() { self.SceneSelectNote() })
	}()

}

func pingFailed(app *App, pingErr error) {
	btnNo := widget.NewButton(
		"No",
		func() {
			app.Quit()
		},
	)

	btnYes := widget.NewButton(
		"Yes",
		func() {
			app.settings.SceneChangeSettings(
				app.window,
				func(previousSceneErr error) {
					if previousSceneErr != nil {
						app.ScenePanic(previousSceneErr.Error())
					}
					app.ScenePing()
				},
			)
		},
	)

	container := container.NewVBox(
		widget.NewLabel(
			fmt.Sprintf("Ping Failed:\n%v\n\nDo you want to change your settings?", pingErr),
		),
		container.NewGridWithColumns(2, btnNo, btnYes),
	)

	fyne.Do(func() { app.window.SetContent(container) })
}
