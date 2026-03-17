package client

import "github.com/quic-go/quic-go"

func (self *App) SceneCancel(conn *quic.Conn, stream *quic.Stream) {
	stream.Close()
	// conn.CloseWithError(0, "")
	self.Quit()
}
