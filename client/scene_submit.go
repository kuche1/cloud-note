package client

import (
	"fmt"
	"time"

	"github.com/quic-go/quic-go"
)

func (self *App) SceneSubmit(conn *quic.Conn, stream *quic.Stream, newText string) {
	// TODO: Add GUI indication

	fmt.Printf("DBG: Sending new text: %v\n", newText)

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

	fmt.Printf("DBG: Sent\n")

	// TODO: Without this THE SHIT DOESNT WORK AAAAAA
	time.Sleep(time.Second * 1)

	self.SceneCancel(conn, stream)
}
