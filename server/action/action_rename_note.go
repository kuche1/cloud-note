package action

import (
	"github.com/kuche1/cloud-note/lib"
	"github.com/kuche1/cloud-note/server/filesystem"
	"github.com/kuche1/cloud-note/server/srvnet"
	"github.com/quic-go/quic-go"
)

func actionRenameNote(conn *quic.Conn, stream *quic.Stream, fs *filesystem.Filesystem) error {
	oldName, err := srvnet.StreamRecvNotename(stream)
	if err != nil {
		return err
	}

	newName, err := srvnet.StreamRecvNotename(stream)
	if err != nil {
		return err
	}

	refusal, err := fs.FileRename(oldName, newName)
	if err != nil {
		return err
	}

	if len(refusal) > 0 {
		err = lib.StreamSendOkOrNot(stream, false, refusal)
		if err != nil {
			return err
		}
		return nil
	}

	err = lib.StreamSendOkOrNot(stream, true, "")
	if err != nil {
		return err
	}

	return nil
}
