package action

import (
	"github.com/kuche1/cloud-note/lib"
	"github.com/kuche1/cloud-note/server/filesystem"
	"github.com/kuche1/cloud-note/server/srvnet"
	"github.com/quic-go/quic-go"
)

func actionDeleteExistingNote(conn *quic.Conn, stream *quic.Stream, fs *filesystem.Filesystem) error {
	noteName, err := srvnet.StreamRecvNotename(stream)
	if err != nil {
		return err
	}

	err = fs.FileDeleteExisting(noteName)
	if err != nil {
		return err
	}

	err = lib.StreamSendACK(stream)
	if err != nil {
		return err
	}

	return nil
}
