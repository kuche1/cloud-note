package action

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/kuche1/cloud-note/lib"
	"github.com/kuche1/cloud-note/server/config"
	"github.com/kuche1/cloud-note/server/filesystem"
	"github.com/kuche1/cloud-note/server/srvnet"
	"github.com/quic-go/quic-go"
)

func HandleAction(conn *quic.Conn, fs *filesystem.Filesystem) (_errString error, _errCode quic.ApplicationErrorCode) {
	stream, err := conn.AcceptStream(context.Background())
	if err != nil {
		return fmt.Errorf("Could not acceept stream: %v", err), 0
	}
	defer func() {
		err := stream.Close()
		if err != nil {
			log.Printf("Could not close stream: %v", err)
		}
	}()

	password, err := srvnet.StreamRecvPassword(stream)
	if err != nil {
		return fmt.Errorf("Could not receive password: %v", err), 0
	}

	err = fs.CheckPassword(password)
	if err != nil {
		return err, 0
	}

	for {
		action, err := recvAction(stream)

		if err != nil {

			netErr, ok := errors.AsType[net.Error](err)
			if ok {
				if netErr.Timeout() {
					return err, lib.ErrorCodeTimeoutDuringActionRead
				}
			}

			return err, 0
		}

		actionFunc := func(
			conn *quic.Conn,
			stream *quic.Stream,
			fs *filesystem.Filesystem,
		) error {
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
			return fmt.Errorf("Unhandled action: %v", action), 0
		}

		err = actionFunc(conn, stream, fs)
		if err != nil {
			return fmt.Errorf("Could not execute action with ID %v: %v", action, err), 0
		}

		// err = lib.ConnRecvEOF(conn)
		// if err != nil {
		// 	return fmt.Errorf("Could not receive connection EOF: %v", err)
		// }
	}

	// return nil
}

func recvAction(stream *quic.Stream) (lib.Action, error) {
	err := lib.DeadlineSet(stream, config.RecvActionDeadline)
	if err != nil {
		return 0, err
	}

	action, err := lib.StreamRecvAction(stream)

	if err != nil {
		return 0, err
	}

	err = lib.DeadlineClear(stream)
	if err != nil {
		return 0, err
	}

	return action, nil
}
