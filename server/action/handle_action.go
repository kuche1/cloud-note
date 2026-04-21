package action

import (
	"context"
	"fmt"

	"github.com/kuche1/cloud-note/lib"
	"github.com/kuche1/cloud-note/server/filesystem"
	"github.com/kuche1/cloud-note/server/srvnet"
	"github.com/quic-go/quic-go"
)

func HandleAction(conn *quic.Conn, fs *filesystem.Filesystem) error {

	stream, err := conn.AcceptStream(context.Background())
	if err != nil {
		return fmt.Errorf("Could not acceept stream: %v", err)
	}

	password, err := srvnet.StreamRecvPassword(stream)
	if err != nil {
		return fmt.Errorf("Could not receive password: %v", err)
	}

	err = fs.CheckPassword(password)
	if err != nil {
		return err
	}

	action, err := lib.StreamRecvAction(stream)
	if err != nil {
		return err
	}

	err = lib.StreamRecvEOF(stream)
	if err != nil {
		return err
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
	case lib.ActionPing:
		actionFunc = actionPing
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
