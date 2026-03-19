package action

import (
	"fmt"

	"github.com/kuche1/cloud-note/client/config"
	"github.com/kuche1/cloud-note/client/output"
	"github.com/kuche1/cloud-note/client/settings"
	"github.com/kuche1/cloud-note/client/window"
	"github.com/kuche1/cloud-note/lib"
)

func ActionListNotes(window *window.Window, output output.Output, settings *settings.Settings) ([]string, error) {
	conn, err := connectToServer(window, output, settings)
	if err != nil {
		return nil, err
	}
	defer func() {
		output.Println("Closing connection...")
		lib.ConnSendEOF(conn)
		output.Println("Done")
	}()

	output.Println("Sending action list notes...")

	err = lib.ChanSendActionEOF(conn, lib.ActionListNotes)
	if err != nil {
		return nil, fmt.Errorf("Could not send action get note: %v", err)
	}

	output.Println("Receiving list of notes...")

	notes, err := lib.ChanRecvSliceStringEOF(conn, config.NumberOfNotesMaxLength)
	if err != nil {
		return nil, err
	}

	output.Println("Received list of notes!")

	return notes, nil
}
