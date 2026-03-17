package client

import (
	"fmt"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (self *App) SceneLoadSettings() {
	addr, ok, err := LoadServerAddr()
	if err != nil {
		self.ScenePanic(fmt.Sprintf("Could not load server address: %v", err))
		return
	}
	if !ok {
		sceneEnterServerAddr(self)
		return
	}

	self.SceneConnectToServer(addr)
}

func sceneEnterServerAddr(app *App) {
	label := widget.NewLabel("Enter Server Address:")

	entry := widget.NewEntry()
	entry.PlaceHolder = "some.address:1234"

	button := widget.NewButton("OK", func() {
		app.SceneConnectToServer(entry.Text)
	})

	container := container.NewVBox(
		label,
		entry,
		button,
	)

	app.window.SetContent(container)
}
