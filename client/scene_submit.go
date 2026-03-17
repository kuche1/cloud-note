package client

import (
	"github.com/quic-go/quic-go"
)

func (self *App) SceneSubmit(conn *quic.Conn, stream *quic.Stream, newText string) {
	// TODO: Add GUI indication

	data := []byte(newText)
	bytesWritten, err := stream.Write(data)
	if err != nil {
		// TODO: Show in GUI
		panic(err)
	}
	if bytesWritten != len(data) {
		// TODO: Loop instead
		// TODO: Show in GUI
		panic("Not all data sent")
	}

	self.SceneCancel(conn, stream)
}
