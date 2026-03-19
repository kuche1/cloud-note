package action

import (
	"fmt"

	"github.com/kuche1/cloud-note/client/config"
	"github.com/kuche1/cloud-note/client/output"
	"github.com/kuche1/cloud-note/client/settings"
	"github.com/kuche1/cloud-note/client/window"
	"github.com/kuche1/cloud-note/lib"
)

func ActionGetNoteContent(window *window.Window, output output.Output, settings *settings.Settings, noteName string) ([]byte, error) {
	conn, err := connectToServer(window, output, settings)
	if err != nil {
		return nil, err
	}
	defer func() {
		output.Println("Closing connection...")
		lib.ConnSendEOF(conn)
		output.Println("Done")
	}()

	output.Println("Sending action get note...")

	err = lib.ChanSendActionEOF(conn, lib.ActionGetNoteContent)
	if err != nil {
		return nil, fmt.Errorf("Could not send action get note: %v", err)
	}

	output.Println("Sending note name...")

	err = lib.ChanSendDatalenSliceByteEOF(conn, []byte(noteName))
	if err != nil {
		return nil, err
	}

	output.Println("Receiving note content...")

	// IMPROVE000: ? Add a loading bar, maybe when sending too
	data, err := lib.ChanRecvDatalenSliceByteEOF(conn, config.NoteContentsMaxLength)
	if err != nil {
		return nil, fmt.Errorf("Could not receive note content:\n%v", err)
	}

	return data, nil
}
