package app

import (
	"fmt"
	"slices"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/kuche1/cloud-note/client/action"
	"github.com/kuche1/cloud-note/client/output"
)

func (self *App) SceneSelectNote() {
	output, outputWidget := output.NewOutputFyneAny()
	self.window.SetContent(outputWidget)

	go fetchNotes(self, output)
}

func fetchNotes(app *App, output output.Output) {
	notes, err := action.ActionListNotes(app.window, output, app.settings)
	if err != nil {
		app.ScenePanic(fmt.Sprintf("Could not fetch notes: %v", err))
		return
	}

	fyne.Do(func() { sceneSelectNote(app, notes) })
}

func sceneSelectNote(app *App, notes []string) {
	/////

	list := widget.NewSelect(
		notes,
		func(selection string) {},
	)

	if len(notes) == 0 {
		list.PlaceHolder = "[No Pre-Existing Notes]"
		app.settings.SetLastEditedNote("")
	} else {
		list.PlaceHolder = "[Select Note]"
	}

	if app.settings.LastEditedNote == "" {
		if len(notes) == 1 {
			onlyNote := notes[0]
			list.SetSelected(onlyNote)
			app.settings.SetLastEditedNote(onlyNote)
		}
	} else {
		if slices.Contains(notes, app.settings.LastEditedNote) {
			list.SetSelected(app.settings.LastEditedNote)
		} else {
			list.PlaceHolder = fmt.Sprintf("[No Longer Available] %v", app.settings.LastEditedNote)
		}
	}

	/////

	edit := widget.NewButton(
		"View / Edit", // TODO: just make 2 separate buttons
		func() {
			selection := list.Selected
			if selection == "" {
				return
			}
			app.settings.SetLastEditedNote(selection)

			app.SceneReceiveNote(selection)
		},
	)

	newNote := widget.NewButton(
		"New Note",
		func() {
			app.SceneCreateNewNote()
		},
	)

	deleteNote := widget.NewButton(
		"Delete Note",
		func() {
			selection := list.Selected
			if selection == "" {
				return
			}
			app.settings.SetLastEditedNote(selection) // Actually it turns out better if this is here

			app.SceneDeleteNote(selection)
		},
	)

	btnSettings := widget.NewButton(
		"Settings",
		func() {
			app.settings.SceneChangeSettings(
				app.window,
				func(previousSceneErr error) {
					if previousSceneErr != nil {
						app.ScenePanic(previousSceneErr.Error())
						return
					}
					app.ScenePing()
				},
			)
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
			widget.NewSeparator(),
			edit,
			widget.NewSeparator(),
			widget.NewLabel(""),
			widget.NewSeparator(),
			newNote,
			widget.NewSeparator(),
			deleteNote,
			widget.NewSeparator(),
			widget.NewLabel(""),
			widget.NewSeparator(),
			btnSettings,
			widget.NewSeparator(),
			quit,
		),
		nil,
		nil,
	)

	app.window.SetContent(container)
}
