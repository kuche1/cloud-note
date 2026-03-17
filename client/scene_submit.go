package client

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/kuche1/cloud-note/lib"
	"github.com/quic-go/quic-go"
)

func (self *App) SceneSubmit(conn *quic.Conn, stream *quic.Stream, newText string) {
	output := widget.NewTextGrid()
	self.window.SetContent(output)

	go submit(self, output, conn, stream, newText)
}

func submit(self *App, output *widget.TextGrid, conn *quic.Conn, stream *quic.Stream, newText string) {
	fyne.Do(func() {
		// IMPROVE: Only do so if the content has actually changed
		// (or maybe not, we'll see what architecture I'll go for)
		output.Append("Sending new note content...")
	})

	err := lib.SendDatalenSliceByte(stream, []byte(newText))
	if err != nil {
		// TODO: Show in GUI
		panic(err)
	}

	fyne.Do(func() {
		output.Append("Receiving acknowledgement...")
	})

	err = lib.RecvEOF(stream)
	if err != nil {
		// TODO: Show in GUI
		panic(err)
	}

	fyne.Do(func() {
		output.Append("Done!")
	})

	fyne.Do(func() {
		self.SceneCancel(conn, stream)
	})
}
