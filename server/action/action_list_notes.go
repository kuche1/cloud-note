package action

import (
	"github.com/kuche1/cloud-note/lib"
	"github.com/kuche1/cloud-note/server/filesystem"
	"github.com/quic-go/quic-go"
)

func actionListNotes(conn *quic.Conn, stream *quic.Stream, fs *filesystem.Filesystem) error {
	files, err := fs.ListFiles()
	if err != nil {
		return err
	}

	err = lib.StreamSendSliceString(stream, files)
	if err != nil {
		return err
	}

	return nil
}
