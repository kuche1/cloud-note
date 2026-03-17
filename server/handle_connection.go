package server

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"

	"github.com/kuche1/cloud-note/lib"
	"github.com/quic-go/quic-go"
)

func handleConnection(conn *quic.Conn) {
	err := handleConnectionInner(conn)
	if err != nil {
		fmt.Printf("%v\n", err)

		err := conn.CloseWithError(0, err.Error())
		if err != nil {
			fmt.Printf("Could not close connection with error: %v", err)
			return
		}
	}
}

func handleConnectionInner(conn *quic.Conn) error {
	if _, err := os.Stat(NoteFile); err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			err = os.WriteFile(NoteFile, []byte{}, 0644)
			if err != nil {
				return fmt.Errorf("Could not create initial note file: %v", err)
			}
		} else {
			return fmt.Errorf("Could not check for note file's existance: %v", err)
		}
	}

	// IMPROVE: Read the file by chunks
	data, err := os.ReadFile(NoteFile)
	if err != nil {
		return fmt.Errorf("Could not read note: %v", err)
	}

	// TODO: Use the nonblocking version
	// TODO: Fix the same issue for the client
	stream, err := conn.OpenStreamSync(context.Background())
	if err != nil {
		return fmt.Errorf("Could not open stream: %v", err)
	}

	err = lib.SendDatalenSliceByte(stream, data)
	if err != nil {
		return fmt.Errorf("Clould not send note content: %v", err)
	}

	log.Printf("Getting new note content...")

	data, err = lib.RecvDatalenSliceByte(stream)
	if err != nil {
		return fmt.Errorf("Could not receive note content: %v", err)
	}

	log.Printf("Got new note content")

	log.Printf("Sending EOF")

	err = lib.SendEOF(stream)
	if err != nil {
		return fmt.Errorf("Could not send EOF: %v", err)
	}

	log.Printf("Sent EOF")

	err = os.WriteFile(NoteFile, data, 0644)
	if err != nil {
		return fmt.Errorf("Could not write note: %v", err)
	}

	log.Printf("Wrote new note content")

	return nil
}
