package app

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/kuche1/cloud-note/client/action"
	"github.com/kuche1/cloud-note/client/settings"
)

func (self *App) SceneSelectNote(settings *settings.Settings) {
	output := widget.NewTextGrid()
	self.window.SetContent(output)

	go fetchNotes(self, output, settings)
}

func fetchNotes(app *App, output *widget.TextGrid, settings *settings.Settings) {
	notes, err := action.ActionListNotes(app.window.FyneWindow, output, settings)
	if err != nil {
		app.ScenePanic(fmt.Sprintf("Could not fetch notes: %v", err))
		return
	}

	fyne.Do(func() { sceneSelectNote(app, settings, notes) })
}

func sceneSelectNote(app *App, settings *settings.Settings, notes []string) {
	list := widget.NewSelect(
		notes,
		func(selection string) {},
	)
	list.PlaceHolder = "[Select Note]"

	button := widget.NewButton(
		"Edit",
		func() {
			selection := list.Selected
			if selection == "" {
				return
			}
			app.SceneReceiveNote(settings, selection)
		},
	)

	container := container.NewBorder(
		list,
		button,
		nil,
		nil,
	)

	app.window.SetContent(container)
}
