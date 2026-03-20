package action

import (
	"fmt"

	"github.com/kuche1/cloud-note/lib"
	"github.com/kuche1/cloud-note/server/filesystem"
	"github.com/kuche1/cloud-note/server/srvnet"
	"github.com/quic-go/quic-go"
)

func actionSetNoteContent(conn *quic.Conn, fs *filesystem.Filesystem) error {
	noteName, err := srvnet.ChanRecvNotenameEOF(conn)
	if err != nil {
		return err
	}

	noteContent, err := srvnet.ChanRecvNotecontentEOF(conn)
	if err != nil {
		return fmt.Errorf("Could not receive new note content: %v", err)
	}

	err = fs.FileWrite(noteName, []byte(noteContent))
	if err != nil {
		return fmt.Errorf("Could not write new note content: %v", err)
	}

	err = lib.ChanSendEOF(conn)
	if err != nil {
		return fmt.Errorf("Could not send save confirmation: %v", err)
	}

	return nil
}
