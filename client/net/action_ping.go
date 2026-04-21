package net

import (
	"github.com/kuche1/cloud-note/client/output"
	"github.com/kuche1/cloud-note/client/settings"
	"github.com/kuche1/cloud-note/client/window"
	"github.com/kuche1/cloud-note/lib"
	"github.com/quic-go/quic-go"
)

func (self *Net) ActionPing(
	window *window.Window,
	output output.Output,
	settings *settings.Settings,
) error {
	stream, err := self.getStream(window, output, settings)
	if err != nil {
		return err
	}

	return self.actionPingCompact(output, stream)
}

func (self *Net) actionPingCompact(
	output output.Output,
	stream *quic.Stream,
) error {
	output.Println("Sending ping...")

	err := lib.StreamSendAction(stream, lib.ActionPing)
	if err != nil {
		return err
	}

	output.Println("Sent!")

	output.Println("Waiting for ACK...")

	err = lib.StreamRecvACK(stream)
	if err != nil {
		return err
	}

	output.Println("Got ACK!")

	return nil
}
