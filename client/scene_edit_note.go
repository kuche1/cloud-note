package client

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/quic-go/quic-go"
)

// IMPROVE: Add some space between the buttons and the text editing
// and maybe move the buttons so that they don't get in the say (in case
// they currently do)
// IMPROVE: I want to use something instead of dragging all this state, maybe have different structs
// for the different states and embed the previous structs along the way
func (self *App) SceneEditNote(conn *quic.Conn, stream *quic.Stream, previousText string, serverAddr string) {
	editor := widget.NewMultiLineEntry()
	editor.Text = previousText
	// editor.Append("asd gfd hgf\nfdsfdsafdsaf")
	// editor.PlaceHolder = "Enter some text"
	editor.TextStyle.Monospace = true

	cancel := widget.NewButton("Cancel", func() { self.SceneCancel(conn, stream, serverAddr) })
	submit := widget.NewButton("Submit", func() { self.SceneSubmit(conn, stream, editor.Text, serverAddr) })
	buttons := container.NewGridWithColumns(
		2,
		cancel,
		submit,
	)

	container := container.NewBorder(
		buttons,
		nil,
		nil,
		nil,
		editor,
	)

	self.window.SetContent(container)
}
