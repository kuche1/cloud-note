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

	err = fs.FileRename(oldName, newName)
	if err != nil {
		return err
	}

	// TODO: instead use `SendStreamOkOrNot(strean, true/false, "Reason"/"")`
	// we need this since otherwise if we try to rename to an invalid note
	// name, we do get an error message explaining that the operation has
	// failed, but we also get a panic after that since we don't reset the
	// stream after the error message and it is dumb to reset the whole
	// stream, better just send an appropriate response we have to send
	// ACK anyways
	err = lib.StreamSendACK(stream)
	if err != nil {
		return err
	}

	return nil
}
