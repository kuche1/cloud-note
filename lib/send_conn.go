package lib

import "github.com/quic-go/quic-go"

func SendConnEOF(conn *quic.Conn) {
	conn.CloseWithError(0, "")
}
