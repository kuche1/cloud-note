package action

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/kuche1/cloud-note/client/settings"
	"github.com/kuche1/cloud-note/client/window"
	"github.com/kuche1/cloud-note/lib"
)

func ActionGetNoteContent(window *window.Window, output *widget.TextGrid, settings *settings.Settings, noteName string) ([]byte, error) {
	conn, err := connectToServer(window, output, settings)
	if err != nil {
		return nil, err
	}

	fyne.Do(func() {
		output.Append("Sending action get note...")
	})

	err = lib.ChanSendActionEOF(conn, lib.ActionGetNoteContent)
	if err != nil {
		return nil, fmt.Errorf("Could not send action get note: %v", err)
	}

	fyne.Do(func() {
		output.Append("Sending note name...")
	})

	err = lib.ChanSendDatalenSliceByteEOF(conn, []byte(noteName))
	if err != nil {
		return nil, err
	}

	fyne.Do(func() {
		output.Append("Receiving note content...")
	})

	// IMPROVE000: ? Add a loading bar, maybe when sending too
	data, err := lib.ChanRecvDatalenSliceByteEOF(conn)
	if err != nil {
		return nil, fmt.Errorf("Could not receive note content:\n%v", err)
	}

	fyne.Do(func() {
		output.Append("Closing connection...")
	})

	lib.ConnSendEOF(conn)

	fyne.Do(func() {
		output.Append("Done!")
	})

	return data, nil
}
