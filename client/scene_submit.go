package client

import (
	"fmt"

	"github.com/kuche1/cloud-note/lib"
	"github.com/quic-go/quic-go"
)

func (self *App) SceneSubmit(conn *quic.Conn, stream *quic.Stream, newText string) {
	// TODO: Add GUI indication, a simple lable or the other thing will suffice

	fmt.Printf("DBG: Sending new text: %v\n", newText)

	err := lib.SendDatalenSliceByte(stream, []byte(newText))
	if err != nil {
		// TODO: Show in GUI
		panic(err)
	}

	fmt.Printf("DBG: Sent\n")

	fmt.Printf("DBG: Receiving EOF\n")

	err = lib.RecvEOF(stream)
	if err != nil {
		// TODO: Show in GUI
		panic(err)
	}

	fmt.Printf("DBG: Received EOF\n")

	// TODO: Without this THE SHIT DOESNT WORK AAAAAA
	// time.Sleep(time.Second * 1)

	self.SceneCancel(conn, stream)
}
