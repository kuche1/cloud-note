package action

import (
	"fmt"

	"github.com/kuche1/cloud-note/lib"
	"github.com/kuche1/cloud-note/server/filesystem"
	"github.com/quic-go/quic-go"
)

func actionGetNoteContent(conn *quic.Conn, fs *filesystem.Filesystem) error {
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

	// TODO: This currently crashes if you have not manually created the `note.txt` file

	// TODO: Maybe make it a standard to send an ACK after receiving the action
	err := lib.SendChannelEOF(conn)
	if err != nil {
		return err
	}

	noteName, err := lib.RecvChannelDatalenSliceByteEOF(conn)
	if err != nil {
		return err
	}

	// IMPROVE000: Read the file by chunks
	noteContent, err := fs.FileRead(string(noteName))
	if err != nil {
		return fmt.Errorf("Could not read note content: %v", err)
	}

	err = lib.SendChannelDatalenSliceByteEOF(conn, noteContent)
	if err != nil {
		return fmt.Errorf("Could not send note content: %v", err)
	}

	return nil
}
