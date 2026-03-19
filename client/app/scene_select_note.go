package app

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/kuche1/cloud-note/client/action"
	"github.com/kuche1/cloud-note/client/output"
	"github.com/kuche1/cloud-note/client/settings"
)

func (self *App) SceneSelectNote(settings *settings.Settings) {
	output, textGrid := output.NewOutputFyneTextGrid()
	self.window.SetContent(textGrid)

	go fetchNotes(self, output, settings)
}

func fetchNotes(app *App, output output.Output, settings *settings.Settings) {
	notes, err := action.ActionListNotes(app.window, output, settings)
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

	edit := widget.NewButton(
		"Edit",
		func() {
			selection := list.Selected
			if selection == "" {
				return
			}
			app.SceneReceiveNote(settings, selection)
		},
	)

	newNote := widget.NewButton(
		"Create New Note",
		func() {
			app.SceneCreateNewNote(settings)
		},
	)

	quit := widget.NewButton(
		"Quit",
		func() {
			app.SceneQuit()
		},
	)

	container := container.NewBorder(
		list,
		container.NewVBox(
			edit,
			widget.NewSeparator(),
			newNote,
			widget.NewSeparator(),
			quit,
		),
		nil,
		nil,
	)

	app.window.SetContent(container)
}
