package action

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/kuche1/cloud-note/client/settings"
	"github.com/kuche1/cloud-note/lib"
)

func ActionListNotes(window *fyne.Window, output *widget.TextGrid, settings *settings.Settings) ([]string, error) {
	conn, err := connectToServer(window, output, settings)
	if err != nil {
		return nil, err
	}

	fyne.Do(func() {
		output.Append("Sending action list notes...")
	})

	err = lib.ChanSendActionEOF(conn, lib.ActionListNotes)
	if err != nil {
		return nil, fmt.Errorf("Could not send action get note: %v", err)
	}

	fyne.Do(func() {
		output.Append("Receiving list of notes...")
	})

	notes, err := lib.ChanRecvSliceStringEOF(conn)
	if err != nil {
		return nil, err
	}

	fyne.Do(func() {
		output.Append("Received list of notes!")
	})

	return notes, nil
}
