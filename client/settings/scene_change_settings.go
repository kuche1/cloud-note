package settings

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/kuche1/cloud-note/client/window"
)

func (self *Settings) SceneChangeSettings(window *window.Window, callbackWhenDone func(previousSceneErr error)) {
	widgetServerAddr, getServerAddr := generateWidgetServerAddr(self)
	widgetServerPassword, getServerPassword := generateWidgetServerPassword(self)

	save := widget.NewButton(
		"Save and Continue",
		func() {
			if new := getServerAddr(); new != "" {
				self.ServerAddr = new
			}
			if new := getServerPassword(); new != "" {
				self.ServerPassword = new
			}

			err := self.Save()

			// We're passing the error to the callback since we do not
			// really have access to `app` from here
			callbackWhenDone(err)
		},
	)

	container := container.NewVBox(
		widgetServerAddr,
		widget.NewSeparator(),
		widgetServerPassword,
		widget.NewSeparator(),
		save,
	)

	window.SetContent(container)
}

func generateWidgetServerAddr(settings *Settings) (_widget *fyne.Container, _getValue func() string) {
	label := widget.NewLabel("Server Address:")

	entry := widget.NewEntry()
	entry.Text = settings.ServerAddr
	entry.PlaceHolder = "Example: localhost:4242"

	return container.NewVBox(
		label,
		entry,
	), func() string { return entry.Text }
}

func generateWidgetServerPassword(settings *Settings) (_widget *fyne.Container, _getValue func() string) {
	label := widget.NewLabel("Server Password:")

	entry := widget.NewEntry()
	entry.Text = settings.ServerPassword
	entry.PlaceHolder = "Example: 123"

	return container.NewVBox(
		label,
		entry,
	), func() string { return entry.Text }
}
