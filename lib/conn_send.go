package lib

import "github.com/quic-go/quic-go"

func ConnSendEOF(conn *quic.Conn) {
	conn.CloseWithError(0, "")
}
