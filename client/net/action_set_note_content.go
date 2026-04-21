package net

import (
	"fmt"

	"github.com/kuche1/cloud-note/client/output"
	"github.com/kuche1/cloud-note/client/settings"
	"github.com/kuche1/cloud-note/client/window"
	"github.com/kuche1/cloud-note/lib"
)

func (self *Net) ActionSetNoteContent(
	window *window.Window,
	output output.Output,
	newText string,
	settings *settings.Settings,
	noteName string,
) error {
	stream, err := self.getStream(window, output, settings)
	if err != nil {
		return err
	}

	output.Println("Sending action set note...")

	err = lib.StreamSendAction(stream, lib.ActionSetNoteContent)
	if err != nil {
		return fmt.Errorf("Could not send action set note: %v", err)
	}

	output.Println("Sent!")

	output.Println("Sending note name...")

	err = lib.StreamSendDatalenString(stream, noteName)
	if err != nil {
		return err
	}

	output.Println("Sending note content...")

	err = lib.StreamSendDatalenString(stream, newText)
	if err != nil {
		return fmt.Errorf("Could not send new note content:\n%v", err)
	}

	output.Println("Receiving save confirmation...")

	err = lib.StreamRecvACK(stream)
	if err != nil {
		return fmt.Errorf("Did not receive save confirmation:\n%v", err)
	}

	return nil
}
