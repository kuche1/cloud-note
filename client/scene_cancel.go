package client

import (
	"fmt"

	"github.com/quic-go/quic-go"
)

func (self *App) SceneCancel(conn *quic.Conn, stream *quic.Stream) {
	fmt.Printf("DBG: Cancelling stream for writing...\n")

	err := stream.Close()
	if err != nil {
		// TODO: Show in GUI
		panic(err)
	}

	fmt.Printf("DBG: Cancelled\n")

	// conn.CloseWithError(0, "")

	self.Quit()
}
