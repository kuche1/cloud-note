package action

import (
	"fmt"

	"github.com/kuche1/cloud-note/client/output"
	"github.com/kuche1/cloud-note/client/settings"
	"github.com/kuche1/cloud-note/client/window"
	"github.com/kuche1/cloud-note/lib"
)

func ActionSetNoteContent(
	window *window.Window,
	output output.Output,
	newText string,
	settings *settings.Settings,
	noteName string,
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

	output.Println("Sending action set note...")

	err = lib.ChanSendActionEOF(conn, lib.ActionSetNoteContent)
	if err != nil {
		return fmt.Errorf("Could not send action set note: %v", err)
	}

	output.Println("Sending note name...")

	err = lib.ChanSendDatalenSliceByteEOF(conn, []byte(noteName))
	if err != nil {
		return err
	}

	output.Println("Sending note content...")

	err = lib.ChanSendDatalenSliceByteEOF(conn, []byte(newText))
	if err != nil {
		return fmt.Errorf("Could not send new note content:\n%v", err)
	}

	output.Println("Receiving save confirmation...")

	err = lib.ChanRecvEOF(conn)
	if err != nil {
		return fmt.Errorf("Did not receive save confirmation:\n%v", err)
	}

	return nil
}
