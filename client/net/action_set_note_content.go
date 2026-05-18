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
	stream, deferStream, err := self.Connect(window, output, settings)
	if err != nil {
		return err
	}
	defer deferStream()

	output.Println("Sending action...")

	lib.StreamSendAction(stream, lib.ActionSetNoteContent)

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
