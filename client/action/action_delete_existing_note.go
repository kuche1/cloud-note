package action

import (
	"github.com/kuche1/cloud-note/client/output"
	"github.com/kuche1/cloud-note/client/settings"
	"github.com/kuche1/cloud-note/client/window"
	"github.com/kuche1/cloud-note/lib"
)

func ActionDeleteExistingNote(
	noteName string,
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

	output.Println("Sending action...")

	err = lib.ChanSendActionEOF(conn, lib.ActionDeleteExistingNote)
	if err != nil {
		return err
	}

	output.Println("Sending new note name...")

	err = lib.ChanSendStringEOF(conn, noteName)
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
