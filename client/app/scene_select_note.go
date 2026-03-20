package app

import (
	"fmt"
	"slices"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/kuche1/cloud-note/client/action"
	"github.com/kuche1/cloud-note/client/output"
	"github.com/kuche1/cloud-note/client/settings"
)

func (self *App) SceneSelectNote(settings *settings.Settings) {
	output, outputWidget := output.NewOutputFyneAny()
	self.window.SetContent(outputWidget)

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
	/////

	list := widget.NewSelect(
		notes,
		func(selection string) {},
	)

	if len(notes) == 0 {
		list.PlaceHolder = "[No Pre-Existing Notes]"
		settings.SetLastEditedNote("")
	} else {
		list.PlaceHolder = "[Select Note]"
	}

	if settings.LastEditedNote == "" {
		if len(notes) == 1 {
			onlyNote := notes[0]
			list.SetSelected(onlyNote)
			settings.SetLastEditedNote(onlyNote)
		}
	} else {
		if slices.Contains(notes, settings.LastEditedNote) {
			list.SetSelected(settings.LastEditedNote)
		} else {
			list.PlaceHolder = fmt.Sprintf("[No Longer Available] %v", settings.LastEditedNote)
		}
	}

	/////

	edit := widget.NewButton(
		"Edit",
		func() {
			selection := list.Selected
			if selection == "" {
				return
			}
			settings.SetLastEditedNote(selection)

			app.SceneReceiveNote(settings, selection)
		},
	)

	newNote := widget.NewButton(
		"Create New Note",
		func() {
			app.SceneCreateNewNote(settings)
		},
	)

	deleteNote := widget.NewButton(
		"Delete Note",
		func() {
			selection := list.Selected
			if selection == "" {
				return
			}
			settings.SetLastEditedNote(selection) // Actually it turns out better if this is here

			app.SceneDeleteNote(settings, selection)
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
			deleteNote,
			widget.NewSeparator(),
			quit,
		),
		nil,
		nil,
	)

	app.window.SetContent(container)
}
