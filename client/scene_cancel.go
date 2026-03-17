package client

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/quic-go/quic-go"
)

func (self *App) SceneCancel(conn *quic.Conn, stream *quic.Stream) {
	output := widget.NewTextGrid()
	self.window.SetContent(output)

	go cancel(self, output, conn, stream)
}

func cancel(app *App, output *widget.TextGrid, conn *quic.Conn, stream *quic.Stream) {
	fyne.Do(func() {
		// IMPROVE: Only do so if the content has actually changed
		// (or maybe not, we'll see what architecture I'll go for)
		output.Append("Closing stream for writing...")
	})

	err := stream.Close()
	if err != nil {
		app.ScenePanic(fmt.Sprintf("Could not close stream for writing:\n%v", err))
		return
	}

	// conn.CloseWithError(0, "")

	fyne.Do(func() {
		// IMPROVE: Only do so if the content has actually changed
		// (or maybe not, we'll see what architecture I'll go for)
		output.Append("Quitting GUI...")
	})

	fyne.Do(func() {
		app.Quit()
	})
}
