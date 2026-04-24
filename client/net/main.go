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

func (self *Net) Disconnect() {
	// TODO: maybe take better care of the 2 return errors ?
	self.stream.Close()
	self.conn.CloseWithError(0, "")

	self.stream = nil
	self.conn = nil
}

func (self *Net) Quit() {
	// TODO: implement something (?)
	// careful: if we are to call `Disconnect` it is theoretically
	// possible that `self.stream` is `nil` and that would cause
	// a crash
}
