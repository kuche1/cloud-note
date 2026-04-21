package net

import (
	"fmt"

	"github.com/kuche1/cloud-note/client/config"
	"github.com/kuche1/cloud-note/client/output"
	"github.com/kuche1/cloud-note/client/settings"
	"github.com/kuche1/cloud-note/client/window"
	"github.com/kuche1/cloud-note/lib"
)

func (self *Net) ActionGetNoteContent(
	window *window.Window,
	output output.Output,
	settings *settings.Settings,
	noteName string,
) ([]byte, error) {
	stream, err := self.getStream(window, output, settings)
	if err != nil {
		return nil, err
	}

	output.Println("Sending action...")

	err = lib.StreamSendAction(stream, lib.ActionGetNoteContent)
	if err != nil {
		return nil, fmt.Errorf("Send action get note content: %v", err)
	}

	output.Println("Sent!")

	output.Println("Sending note name...")

	err = lib.StreamSendDatalenString(stream, noteName)
	if err != nil {
		return nil, err
	}

	output.Println("Receiving note content...")

	// IMPROVE000: ? Add a loading bar, maybe when sending too
	data, err := lib.StreamRecvDatalenSliceByte(stream, config.NoteContentsMaxLength)
	if err != nil {
		return nil, fmt.Errorf("Could not receive note content:\n%v", err)
	}

	output.Println("Done!")

	return data, nil
}
