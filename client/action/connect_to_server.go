package action

import (
	"context"
	"crypto/tls"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/kuche1/cloud-note/client/settings"
	"github.com/kuche1/cloud-note/lib"
	"github.com/quic-go/quic-go"
)

// TODO: Make this require `Window` rather than `fyne.Window`
func connectToServer(window *fyne.Window, output *widget.TextGrid, settings *settings.Settings) (*quic.Conn, error) {
	fyne.Do(func() {
		output.Append("Connecting to server...")
	})

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

		return nil, retErr
	}

	fyne.Do(func() {
		output.Append("Connected!")
	})

	return conn, nil
}
