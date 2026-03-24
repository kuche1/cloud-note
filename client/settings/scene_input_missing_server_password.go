package settings

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/kuche1/cloud-note/client/window"
)

func (self *Settings) sceneInputMissingServerPassword(window *window.Window, callbackWhenAllDone func()) {
	if self.ServerPassword != "" {
		callbackWhenAllDone()
		return
	}

	label := widget.NewLabel("Enter Server Password")

	entry := widget.NewEntry()
	entry.PlaceHolder = "Example: 123"

	button := widget.NewButton(
		"Ok",
		func() {
			addr := entry.Text
			if addr == "" {
				return
			}
			self.SetServerPassword(addr)

			callbackWhenAllDone()
		},
	)

	container := container.NewVBox(
		label,
		entry,
		button,
	)

	window.SetContent(container)
}
