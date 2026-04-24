package net

import (
	"github.com/kuche1/cloud-note/client/config"
	"github.com/kuche1/cloud-note/client/output"
	"github.com/kuche1/cloud-note/client/settings"
	"github.com/kuche1/cloud-note/client/window"
	"github.com/kuche1/cloud-note/lib"
)

func (self *Net) ActionListNotes(
	window *window.Window,
	output output.Output,
	settings *settings.Settings,
) ([]string, error) {
	stream, err := self.getStream(window, output, settings, lib.ActionListNotes)
	if err != nil {
		return nil, err
	}

	output.Println("Receiving list of notes...")

	notes, err := lib.StreamRecvSliceString(
		stream,
		config.NumberOfNotesMaxLength,
		config.NoteNameMaxLength,
	)
	if err != nil {
		return nil, err
	}

	output.Println("Done!")

	return notes, nil
}
