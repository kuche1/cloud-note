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
	conn, stream, err := connectToServer(window, output, settings)
	if err != nil {
		return err
	}
	defer func() {
		output.Println("Closing stream...")
		lib.StreamSendEOFUnchecked(stream)
		output.Println("Closing connection...")
		lib.ConnSendEOF(conn)
		output.Println("Done")
	}()

	output.Println("Sending action set note...")

	err = lib.StreamSendAction(stream, lib.ActionSetNoteContent)
	if err != nil {
		return fmt.Errorf("Could not send action set note: %v", err)
	}

	lib.StreamSendEOFUnchecked(stream) // TODO: not great

	output.Println("Sending note name...")

	err = lib.ChanSendStringEOF(conn, noteName)
	if err != nil {
		return err
	}

	output.Println("Sending note content...")

	err = lib.ChanSendStringEOF(conn, newText)
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
