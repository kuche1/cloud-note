package app

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Can also be called from within other threads
// Theoretically it is still possible to bug the panic screen, if
// any `self.window.SetContent`s have call queued after the panic
func (self *App) ScenePanic(info string) {
	output := widget.NewRichTextWithText(fmt.Sprintf("Panic:\n%v", info))
	output.Wrapping = fyne.TextWrapWord

	button := widget.NewButton("Quit", func() { self.Quit() })

	container := container.NewBorder(
		nil,
		button,
		nil,
		nil,
		output,
	)

	// Wrapping this in a `fyne.Do` so that it can be called from anywhere, even other threads
	fyne.Do(func() {
		self.window.SetContent(container)
	})
}
