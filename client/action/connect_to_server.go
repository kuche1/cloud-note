package action

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

func connectToServer(
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
			return connectToServer(window, output, settings)
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
