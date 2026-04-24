package net

import (
	"context"
	"crypto/tls"
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
	action lib.Action,
) (*quic.Stream, error) {
	if self.stream == nil {

		conn, stream, err := connServer(window, output, settings)
		if err != nil {
			output.Println("Failure!")
			return nil, err
		}

		self.conn = conn
		self.stream = stream

	}

	output.Println("Sending action...")

	err := lib.StreamSendAction(self.stream, action)

	if err != nil {

		if lib.ErrorIsTimeout(err) {
			self.Disconnect()
			// will set `self.conn` and `self.stream` to nil

			// TODO: we can potentially ented a never-ending timeout loop (or maybe not really)
			return self.getStream(window, output, settings, action)
		}

		return nil, err
	}

	output.Println("Sent!")

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
