package settings

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/kuche1/cloud-note/client/window"
)

func (self *Settings) sceneInputMissingServerAddr(window *window.Window, callbackWhenAllDone func()) {
	if self.ServerAddr != "" {
		callbackWhenAllDone()
		return
	}

	label := widget.NewLabel("Enter Server Address")

	entry := widget.NewEntry()
	entry.PlaceHolder = "Example: localhost:4242"

	button := widget.NewButton(
		"Ok",
		func() {
			addr := entry.Text
			if addr == "" {
				return
			}
			self.ServerAddr = addr

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
