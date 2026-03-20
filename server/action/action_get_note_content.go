package action

import (
	"fmt"

	"github.com/kuche1/cloud-note/lib"
	"github.com/kuche1/cloud-note/server/filesystem"
	"github.com/kuche1/cloud-note/server/srvnet"
	"github.com/quic-go/quic-go"
)

func actionGetNoteContent(conn *quic.Conn, fs *filesystem.Filesystem) error {
	noteName, err := srvnet.ChanRecvNotenameEOF(conn)
	if err != nil {
		return err
	}

	noteContent, err := fs.FileRead(noteName)
	if err != nil {
		return fmt.Errorf("Could not read note content: %v", err)
	}

	err = lib.ChanSendDatalenSliceByteEOF(conn, noteContent)
	if err != nil {
		return fmt.Errorf("Could not send note content: %v", err)
	}

	// if _, err := os.Stat(config.NoteFile); err != nil {
	// 	if errors.Is(err, importFs.ErrNotExist) {
	// 		err = os.WriteFile(config.NoteFile, []byte{}, 0600)
	// 		if err != nil {
	// 			return fmt.Errorf("Could not create initial note file: %v", err)
	// 		}
	// 	} else {
	// 		return fmt.Errorf("Could not check for note file's existance: %v", err)
	// 	}
	// }

	return nil
}
