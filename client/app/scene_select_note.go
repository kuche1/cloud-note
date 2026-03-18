package app

import (
	"fmt"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/kuche1/cloud-note/client/settings"
)

func (self *App) SceneSelectNote(settings *settings.Settings) {
	list := widget.NewSelect(
		[]string{"a.txt", "b.txt", "c.txt"},
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
			fmt.Printf("DBG: Ignoring the actual choice of: %v\n", selection) // TODO
			self.SceneReceiveNote(settings)
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
