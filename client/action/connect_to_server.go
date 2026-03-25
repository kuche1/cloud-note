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
	window *window.Window, // TODO: No longer needed
	output output.Output,
	settings *settings.Settings,
) (_conn *quic.Conn, _err error) {
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

		//// TODO: This no longer makes much sense
		// ok := settings.PromptNewServerAddr(
		// 	window,
		// 	retErr.Error(),
		// )
		// if ok {
		// 	return connectToServer(window, output, settings)
		// }

		return nil, retErr
	}

	output.Println("Sending password...")

	err = lib.ChanSendStringEOF(conn, settings.ServerPassword)
	if err != nil {
		// TODO: Ask for new password if this fails OR ask for both new password and
		// server address OR redirect to the settings setup
		return nil, fmt.Errorf("Could not send password to server:\n%v", err)
	}

	output.Println("Connected!")

	return conn, nil
}
