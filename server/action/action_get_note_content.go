package action

import (
	"errors"
	"fmt"
	"io/fs"
	"os"

	"github.com/kuche1/cloud-note/lib"
	"github.com/kuche1/cloud-note/server/config"
	"github.com/quic-go/quic-go"
)

func actionGetNoteContent(conn *quic.Conn) error {
	if _, err := os.Stat(config.NoteFile); err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			err = os.WriteFile(config.NoteFile, []byte{}, 0600)
			if err != nil {
				return fmt.Errorf("Could not create initial note file: %v", err)
			}
		} else {
			return fmt.Errorf("Could not check for note file's existance: %v", err)
		}
	}

	// IMPROVE000: Read the file by chunks
	data, err := os.ReadFile(config.NoteFile)
	if err != nil {
		return fmt.Errorf("Could not read note: %v", err)
	}

	err = lib.SendChannelDatalenSliceByteEOF(conn, data)
	if err != nil {
		return fmt.Errorf("Could not send note content: %v", err)
	}

	return nil
}
