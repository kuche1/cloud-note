package app

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/kuche1/cloud-note/client/config"
	"github.com/kuche1/cloud-note/client/settings"
)

func (self *App) SceneSelectNote(settings *settings.Settings) {
	list := widget.NewSelect(
		[]string{config.NoteName}, // TODO: Get this list from the server
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
			self.SceneReceiveNote(settings, selection)
		},
	)

	container := container.NewBorder(
		list,
		button,
		nil,
		nil,
	)

	self.window.SetContent(container)
}
