package net

import (
	"github.com/quic-go/quic-go"
)

type Net struct {
	conn   *quic.Conn
	stream *quic.Stream
}

func NewNet(
// window *window.Window,
// output output.Output,
// settings *settings.Settings,
) *Net {
	// conn, stream, err := connectToServer(window, output, settings)
	// if err != nil {
	// 	return nil, err
	// }

	// return &Con{
	// 	conn:   conn,
	// 	stream: stream,
	// }, nil

	return &Net{} // TOOD: I hate this
}
