package action

import (
	"github.com/kuche1/cloud-note/client/output"
	"github.com/kuche1/cloud-note/client/settings"
	"github.com/kuche1/cloud-note/client/window"
	"github.com/kuche1/cloud-note/lib"
)

func ActionPing(
	window *window.Window,
	output output.Output,
	settings *settings.Settings,
) error {
	conn, err := connectToServer(window, output, settings)
	if err != nil {
		return err
	}
	defer func() {
		output.Println("Closing connection...")
		lib.ConnSendEOF(conn)
		output.Println("Done")
	}()

	output.Println("Sending ping...")

	err = lib.ChanSendActionEOF(conn, lib.ActionPing)
	if err != nil {
		return err
	}

	output.Println("Waiting for ACK...")

	err = lib.ChanRecvEOF(conn)
	if err != nil {
		return err
	}

	return nil
}
