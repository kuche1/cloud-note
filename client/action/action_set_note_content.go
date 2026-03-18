package action

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/kuche1/cloud-note/client/settings"
	"github.com/kuche1/cloud-note/lib"
)

func ActionSetNoteContent(window *fyne.Window, output *widget.TextGrid, newText string, settings *settings.Settings, noteName string) error {
	conn, err := connectToServer(window, output, settings)
	if err != nil {
		return err
	}

	fyne.Do(func() {
		output.Append("Sending action set note...")
	})

	err = lib.SendChannelActionEOF(conn, lib.ActionSetNoteContent)
	if err != nil {
		return fmt.Errorf("Could not send action set note: %v", err)
	}

	fyne.Do(func() {
		output.Append("Sending note name...")
	})

	err = lib.SendChannelDatalenSliceByteEOF(conn, []byte(noteName))
	if err != nil {
		return err
	}

	fyne.Do(func() {
		output.Append("Sending note content...")
	})

	err = lib.SendChannelDatalenSliceByteEOF(conn, []byte(newText))
	if err != nil {
		return fmt.Errorf("Could not send new note content:\n%v", err)
	}

	fyne.Do(func() {
		output.Append("Receiving save confirmation...")
	})

	err = lib.RecvChannelEOF(conn)
	if err != nil {
		return fmt.Errorf("Did not receive save confirmation:\n%v", err)
	}

	fyne.Do(func() {
		output.Append("Closing connection...")
	})

	lib.SendConnEOF(conn)

	fyne.Do(func() {
		output.Append("Done!")
	})

	return nil
}
