package action

import (
	"github.com/kuche1/cloud-note/lib"
	"github.com/kuche1/cloud-note/server/filesystem"
	"github.com/quic-go/quic-go"
)

func actionPing(conn *quic.Conn, fs *filesystem.Filesystem) error {
	err := lib.ChanSendEOF(conn)
	if err != nil {
		return err
	}

	return nil
}
