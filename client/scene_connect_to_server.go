package client

import (
	"context"
	"crypto/tls"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/kuche1/cloud-note/lib"
	"github.com/quic-go/quic-go"
)

func (self *App) SceneConnectToServer() {
	output := widget.NewTextGrid()
	self.window.SetContent(output)

	go connectToServer(self, output)
}

func connectToServer(app *App, output *widget.TextGrid) {
	fyne.Do(func() {
		output.Append("Connecting to server...")
	})

	conn, err := quic.DialAddr(
		context.Background(),
		ServerAddr,
		&tls.Config{
			InsecureSkipVerify: true,
			NextProtos:         []string{lib.QuicProto},
		},
		nil,
	)
	if err != nil {
		app.ScenePanic(fmt.Sprintf("Could not connect to server:\n%v", err))
		return
	}

	fyne.Do(func() {
		output.Append("Accepting stream...")
	})

	stream, err := conn.AcceptStream(context.Background())
	if err != nil {
		// TODO: Show in GUI
		panic(err)
	}

	fyne.Do(func() {
		output.Append("Downloading data...")
	})

	data, err := lib.RecvDatalenSliceByte(stream)
	if err != nil {
		// TODO: Show in GUI
		panic(err)
	}

	dataAsStr := string(data)

	fyne.Do(func() {
		output.Append("Done!")
		app.SceneEditNote(conn, stream, dataAsStr)
	})
}
