package client

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/quic-go/quic-go"
)

func (self *App) SceneEditNote(conn *quic.Conn, stream *quic.Stream, previousText string) {
	editor := widget.NewMultiLineEntry()
	editor.Text = previousText
	// editor.Append("asd gfd hgf\nfdsfdsafdsaf")
	// editor.PlaceHolder = "Enter some text"
	editor.TextStyle.Monospace = true

	cancel := widget.NewButton("Cancel", func() { self.SceneCancel(conn, stream) })
	submit := widget.NewButton("Submit", func() { self.SceneSubmit(conn, stream, editor.Text) })
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
