package action

import (
	"github.com/kuche1/cloud-note/lib"
	"github.com/kuche1/cloud-note/server/filesystem"
	"github.com/quic-go/quic-go"
)

func actionPing(conn *quic.Conn, stream *quic.Stream, fs *filesystem.Filesystem) error {
	err := lib.StreamSendACK(stream)
	if err != nil {
		return err
	}

	return nil
}
