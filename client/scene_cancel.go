package client

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/quic-go/quic-go"
)

func (self *App) SceneCancel(conn *quic.Conn, stream *quic.Stream, serverAddr string) {
	output := widget.NewTextGrid()
	self.window.SetContent(output)

	go cancel(self, output, conn, stream, serverAddr)
}

// IMPROVE: Maybe just have a `NetState` that contains both `conn` and `stream`
// And actually attempt to clean everything and just collect and print errorrs instead
// of panicing
func cancel(app *App, output *widget.TextGrid, conn *quic.Conn, stream *quic.Stream, serverAddr string) {
	fyne.Do(func() {
		output.Append("Saving settings...")
	})

	// IMPROVE: This is good because if something fails the user will not be
	// stuck, but ideally we would simply let the user change it
	// OR make it so that if the connection fails the user is prompted to change
	// the server address
	err := SaveServerAddr(serverAddr)
	if err != nil {
		app.ScenePanic(fmt.Sprintf("Could not save settings:\n%v", err))
		return
	}

	fyne.Do(func() {
		// IMPROVE: Only do so if the content has actually changed
		// (or maybe not, we'll see what architecture I'll go for)
		output.Append("Closing stream for writing...")
	})

	err = stream.Close()
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
