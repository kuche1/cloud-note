package action

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/kuche1/cloud-note/client/settings"
	"github.com/kuche1/cloud-note/client/window"
	"github.com/kuche1/cloud-note/lib"
)

func ActionCreateNewNote(newNoteName string, window *window.Window, output *widget.TextGrid, settings *settings.Settings) error {
	conn, err := connectToServer(window, output, settings)
	if err != nil {
		return err
	}

	fyne.Do(func() {
		output.Append("Sending action...")
	})

	err = lib.ChanSendActionEOF(conn, lib.ActionCreateNewNote)
	if err != nil {
		return err
	}

	fyne.Do(func() {
		output.Append("Sending new note name...")
	})

	err = lib.ChanSendStringEOF(conn, newNoteName)
	if err != nil {
		return err
	}

	fyne.Do(func() {
		output.Append("Waiting for ACK...")
	})

	err = lib.ChanRecvEOF(conn)
	if err != nil {
		return err
	}

	fyne.Do(func() {
		output.Append("Done!")
	})

	return nil
}
