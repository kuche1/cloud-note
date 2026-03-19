package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/kuche1/cloud-note/client/output"
)

// Can also be called from within other threads
// Theoretically it is still possible to bug the panic screen, if
// any `self.window.SetContent`s have call queued after the panic
func (self *App) ScenePanic(info string) {
	// output := widget.NewRichTextWithText(fmt.Sprintf("Panic:\n%v", info))
	// output.Wrapping = fyne.TextWrapWord
	// TODO: test this
	output, richText := output.NewOutputFyneRichText()
	output.Println("Panic:")
	output.Println(info)

	button := widget.NewButton("Quit", func() { self.Quit() })

	container := container.NewBorder(
		nil,
		button,
		nil,
		nil,
		richText,
	)

	// Wrapping this in a `fyne.Do` so that it can be called from anywhere, even other threads
	fyne.Do(func() {
		self.window.SetContent(container)
	})
}
