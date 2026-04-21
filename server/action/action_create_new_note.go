package action

import (
	"github.com/kuche1/cloud-note/lib"
	"github.com/kuche1/cloud-note/server/filesystem"
	"github.com/kuche1/cloud-note/server/srvnet"
	"github.com/quic-go/quic-go"
)

func actionCreateNewNote(conn *quic.Conn, stream *quic.Stream, fs *filesystem.Filesystem) error {
	newNoteName, err := srvnet.StreamRecvNotename(stream)
	if err != nil {
		return err
	}

	err = fs.FileCreateNew(newNoteName)
	if err != nil {
		return err
	}

	err = lib.StreamSendACK(stream)
	if err != nil {
		return err
	}

	return nil
}
