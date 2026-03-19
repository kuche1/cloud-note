package action

import (
	"github.com/kuche1/cloud-note/lib"
	"github.com/kuche1/cloud-note/server/filesystem"
	"github.com/quic-go/quic-go"
)

func actionCreateNewNote(conn *quic.Conn, fs *filesystem.Filesystem) error {
	newNoteName, err := lib.ChanRecvStringEOF(conn)
	if err != nil {
		return err
	}

	err = fs.FileCreateNew(newNoteName)
	if err != nil {
		return err
	}

	err = lib.ChanSendEOF(conn)
	if err != nil {
		return err
	}

	return nil
}
