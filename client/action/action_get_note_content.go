package action

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/kuche1/cloud-note/client/settings"
	"github.com/kuche1/cloud-note/lib"
)

func ActionGetNoteContent(window *fyne.Window, output *widget.TextGrid, settings *settings.Settings, noteName string) ([]byte, error) {
	conn, err := connectToServer(window, output, settings)
	if err != nil {
		return nil, err
	}

	fyne.Do(func() {
		output.Append("Sending action get note...")
	})

	err = lib.SendChannelActionEOF(conn, lib.ActionGetNoteContent)
	if err != nil {
		return nil, fmt.Errorf("Could not send action get note: %v", err)
	}

	fyne.Do(func() {
		output.Append("Receiving ACK...")
	})

	err = lib.RecvChannelEOF(conn)
	if err != nil {
		return nil, err
	}

	fyne.Do(func() {
		output.Append("Sending note name...")
	})

	err = lib.SendChannelDatalenSliceByteEOF(conn, []byte(noteName))
	if err != nil {
		return nil, err
	}

	fyne.Do(func() {
		output.Append("Receiving note content...")
	})

	// IMPROVE000: ? Add a loading bar, maybe when sending too
	data, err := lib.RecvChannelDatalenSliceByteEOF(conn)
	if err != nil {
		return nil, fmt.Errorf("Could not receive note content:\n%v", err)
	}

	fyne.Do(func() {
		output.Append("Closing connection...")
	})

	lib.SendConnEOF(conn)

	fyne.Do(func() {
		output.Append("Done!")
	})

	return data, nil
}
