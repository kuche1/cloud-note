package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.New()
	window := app.NewWindow("Simple Label")
	window.Resize(fyne.NewSize(400, 600))

	cancel := widget.NewButton("Cancel", func() { fmt.Printf("Cancel pressed\n") })
	submit := widget.NewButton("Submit", func() { fmt.Printf("Submit pressed\n") })
	buttons := container.NewGridWithColumns(2,
		cancel,
		submit,
	)

	editor := widget.NewMultiLineEntry()
	editor.PlaceHolder = "Enter some text"
	// editor.Append("asd gfd hgf\nfdsfdsafdsaf")

	container := container.NewBorder(
		buttons,
		nil,
		nil,
		nil,
		editor,
	)

	window.SetContent(container)
	window.ShowAndRun()
}
