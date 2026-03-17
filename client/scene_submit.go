package client

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/kuche1/cloud-note/lib"
	"github.com/quic-go/quic-go"
)

func (self *App) SceneSubmit(conn *quic.Conn, stream *quic.Stream, newText string, serverAddr string) {
	output := widget.NewTextGrid()
	self.window.SetContent(output)

	go submit(self, output, conn, stream, newText, serverAddr)
}

func submit(app *App, output *widget.TextGrid, conn *quic.Conn, stream *quic.Stream, newText string, serverAddr string) {
	fyne.Do(func() {
		// IMPROVE: Only do so if the content has actually changed
		// (or maybe not, we'll see what architecture I'll go for)
		output.Append("Sending new note content...")
	})

	err := lib.SendDatalenSliceByte(stream, []byte(newText))
	if err != nil {
		app.ScenePanic(fmt.Sprintf("Could not send new note content:\n%v", err))
		return
	}

	fyne.Do(func() {
		output.Append("Receiving acknowledgement...")
	})

	err = lib.RecvEOF(stream)
	if err != nil {
		app.ScenePanic(fmt.Sprintf("Did not receive acknowledgement that the new note content has been received:\n%v", err))
		return
	}

	fyne.Do(func() {
		output.Append("Done!")
	})

	fyne.Do(func() {
		app.SceneCancel(conn, stream, serverAddr)
	})
}
