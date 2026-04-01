package app

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/kuche1/cloud-note/client/output"
)

// I just checked, and this is not what causes the app to resize to insane height on desktop after
// the users clicks `submit`
func (self *App) IntermissionInfo(messaage string, callbackAfterTheUserHasPressedOk func()) {
	previousFyneContent := self.window.Content()

	output, outputWidget := output.NewOutputFyneAny()
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
