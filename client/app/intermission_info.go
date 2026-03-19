package app

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/kuche1/cloud-note/client/output"
)

func (self *App) IntermissionInfo(messaage string, callbackAfterTheUserHasPressedOk func()) {
	previousFyneContent := self.window.Content()

	output, outputWidget := output.NewOutputFyneTextGrid()
	output.Println(messaage)

	button := widget.NewButton(
		"Ok",
		func() {
			self.window.SetContent(previousFyneContent)
			callbackAfterTheUserHasPressedOk()
		},
	)

	container := container.NewBorder(
		nil,
		button,
		nil,
		nil,
		outputWidget,
	)

	self.window.SetContent(
		container,
	)
}
