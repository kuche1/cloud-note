package net

import (
	"github.com/kuche1/cloud-note/client/output"
	"github.com/kuche1/cloud-note/client/settings"
	"github.com/kuche1/cloud-note/client/window"
	"github.com/kuche1/cloud-note/lib"
)

func (self *Net) ActionRenameNote(
	oldName string,
	newName string,
	window *window.Window,
	output output.Output,
	settings *settings.Settings,
) error {
	stream, err := self.getStream(window, output, settings, lib.ActionRenameNote)
	if err != nil {
		return err
	}

	output.Println("Sending old name...")

	err = lib.StreamSendDatalenString(stream, oldName)
	if err != nil {
		return err
	}

	output.Println("Sending new name...")

	err = lib.StreamSendDatalenString(stream, newName)
	if err != nil {
		return err
	}

	output.Println("Waiting for ACK...")

	err = lib.StreamRecvACK(stream)
	if err != nil {
		return err
	}

	output.Println("Done")

	return nil
}
