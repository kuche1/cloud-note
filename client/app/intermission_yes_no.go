package app

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/kuche1/cloud-note/client/output"
)

func (self *App) IntermissionYesNo(messaage string, callbackYes func(), callbackNo func()) {
	previousFyneContent := self.window.Content()

	output, outputWidget := output.NewOutputFyneAny()
	output.Println(messaage)

	yes := widget.NewButton(
		"Yes",
		func() {
			self.window.SetContent(previousFyneContent)
			callbackYes()
		},
	)

	no := widget.NewButton(
		"No",
		func() {
			self.window.SetContent(previousFyneContent)
			callbackNo()
		},
	)

	container := container.NewBorder(
		nil,
		container.NewAdaptiveGrid(2, no, yes),
		nil,
		nil,
		outputWidget,
	)

	self.window.SetContent(
		container,
	)
}
