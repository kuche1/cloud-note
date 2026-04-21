package action

import (
	"fmt"

	"github.com/kuche1/cloud-note/lib"
	"github.com/kuche1/cloud-note/server/filesystem"
	"github.com/kuche1/cloud-note/server/srvnet"
	"github.com/quic-go/quic-go"
)

func actionGetNoteContent(conn *quic.Conn, stream *quic.Stream, fs *filesystem.Filesystem) error {
	noteName, err := srvnet.StreamRecvNotename(stream)
	if err != nil {
		return err
	}

	noteContent, err := fs.FileRead(noteName)
	if err != nil {
		return fmt.Errorf("Could not read note content: %v", err)
	}

	err = lib.StreamSendDatalenSliceByte(stream, noteContent)
	if err != nil {
		return fmt.Errorf("Could not send note content: %v", err)
	}

	return nil
}
