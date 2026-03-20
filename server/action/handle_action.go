package action

import (
	"fmt"

	"github.com/kuche1/cloud-note/lib"
	"github.com/kuche1/cloud-note/server/filesystem"
	"github.com/quic-go/quic-go"
)

func HandleAction(conn *quic.Conn, fs *filesystem.Filesystem) error {
	action, err := lib.ChanRecvActionEOF(conn)
	if err != nil {
		return fmt.Errorf("Could not receive action: %v", err)
	}

	actionFunc := func(conn *quic.Conn, fs *filesystem.Filesystem) error {
		return fmt.Errorf("Unreachable code reached, there is something wrong with the action dispatch")
	}

	switch action {
	case lib.ActionGetNoteContent:
		actionFunc = actionGetNoteContent
	case lib.ActionSetNoteContent:
		actionFunc = actionSetNoteContent
	case lib.ActionListNotes:
		actionFunc = actionListNotes
	case lib.ActionCreateNewNote:
		actionFunc = actionCreateNewNote
	case lib.ActionDeleteExistingNote:
		actionFunc = actionDeleteExistingNote
	default:
		return fmt.Errorf("Unhandled action: %v", action)
	}

	err = actionFunc(conn, fs)
	if err != nil {
		return fmt.Errorf("Could not execute action with ID %v: %v", action, err)
	}

	err = lib.ConnRecvEOF(conn)
	if err != nil {
		return fmt.Errorf("Could not receive connection EOF: %v", err)
	}

	return nil
}
