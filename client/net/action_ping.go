package net

import (
	"github.com/kuche1/cloud-note/client/output"
	"github.com/kuche1/cloud-note/client/settings"
	"github.com/kuche1/cloud-note/client/window"
	"github.com/kuche1/cloud-note/lib"
)

func (self *Net) ActionPing(
	window *window.Window,
	output output.Output,
	settings *settings.Settings,
) error {
	stream, deferStream, err := self.Connect(window, output, settings)
	if err != nil {
		return err
	}
	defer deferStream()

	output.Println("Sending action...")

	lib.StreamSendAction(stream, lib.ActionPing)

	output.Println("Waiting for ACK...")

	err = lib.StreamRecvACK(stream)
	if err != nil {
		return err
	}

	output.Println("Got ACK!")

	return nil
}
