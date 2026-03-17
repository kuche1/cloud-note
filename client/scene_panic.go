package client

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (self *App) ScenePanic(info string) {
	output := widget.NewTextGrid()
	output.Append("Panic:")
	output.Append(info)

	button := widget.NewButton("Quit", func() { self.Quit() })

	container := container.NewBorder(
		nil,
		button,
		nil,
		nil,
		output,
	)

	self.window.SetContent(container)
}
