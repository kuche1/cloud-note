package net

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"

	"github.com/kuche1/cloud-note/client/output"
	"github.com/kuche1/cloud-note/client/settings"
	"github.com/kuche1/cloud-note/client/window"
	"github.com/kuche1/cloud-note/lib"
	"github.com/quic-go/quic-go"
)

func (self *Net) getStream(
	window *window.Window,
	output output.Output,
	settings *settings.Settings,
) (*quic.Stream, error) {
	// TODO: And also make it auto-refresh if it has expired

	if self.stream == nil {

		conn, stream, err := connServer(window, output, settings)
		if err != nil {
			return nil, err
		}

		self.conn = conn
		self.stream = stream

	} else {

		// TODO:
		// 1) I dont like how we have to wait for ACK before our request, ideally
		//     we would check for timeout elsewhere
		// 2) this makes the `ActionPing` actually send 2 pings
		//     (except if we assume that it's only going to get called once at
		//     startup since the if above does not call `actionPingCompact`)
		err := self.actionPingCompact(output, self.stream)
		if err != nil {

			_, ok := errors.AsType[*quic.IdleTimeoutError](err)
			if ok {
				// remote peer's idle timeout expired

				self.Disconnect()
				// will set `self.conn` and `self.stream` to nil

				return self.getStream(window, output, settings)
			}

			return nil, err
		}

	}

	return self.stream, nil
}

func connServer(
	window *window.Window,
	output output.Output,
	settings *settings.Settings,
) (_conn *quic.Conn, _stream *quic.Stream, _err error) {
	output.Println("Connecting to server...")

	conn, err := quic.DialAddr(
		context.Background(),
		settings.ServerAddr,
		&tls.Config{
			InsecureSkipVerify: true,
			NextProtos:         []string{lib.QuicProto},
		},
		nil,
	)
	if err != nil {
		retErr := fmt.Errorf("Could not connect to server:\n%v", err)

		ok := settings.PromptNewServerAddr(
			window,
			retErr.Error(),
		)
		if ok {
			return connServer(window, output, settings)
		}

		return nil, nil, retErr
	}

	output.Println("Connected!")

	output.Println("Sending password...")

	stream, err := conn.OpenStream()
	if err != nil {
		return nil, nil, fmt.Errorf("Could not open stream:\n%v", err)
	}

	err = lib.StreamSendDatalenString(stream, settings.ServerPassword)
	if err != nil {
		lib.StreamSendEOFUnchecked(stream)
		return nil, nil, fmt.Errorf("Could not send password:\n%v", err)
	}

	output.Println("Sent!")

	return conn, stream, nil
}
