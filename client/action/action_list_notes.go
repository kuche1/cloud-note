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
	conn, stream, err := connectToServer(window, output, settings)
	if err != nil {
		return nil, err
	}
	defer func() {
		output.Println("Closing stream...")
		lib.StreamSendEOFUnchecked(stream)
		output.Println("Closing connection...")
		lib.ConnSendEOF(conn)
		output.Println("Done")
	}()

	output.Println("Sending action list notes...")

	err = lib.StreamSendAction(stream, lib.ActionListNotes)
	if err != nil {
		return nil, fmt.Errorf("Could not send action get note: %v", err)
	}

	lib.StreamSendEOFUnchecked(stream) // TODO: not great

	output.Println("Receiving list of notes...")

	notes, err := lib.ChanRecvSliceStringEOF(conn, config.NumberOfNotesMaxLength)
	if err != nil {
		return nil, err
	}

	output.Println("Received list of notes!")

	return notes, nil
}
