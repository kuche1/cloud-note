// IMPROVE: Call `conn.CloseWithError(0, "Info")` in case of failure

package server

import (
	"context"
	"errors"
	"io/fs"
	"log"
	"os"

	"github.com/kuche1/cloud-note/lib"
	"github.com/quic-go/quic-go"
)

func Main() {
	listener, err := quic.ListenAddr(Addr, generateTLSConfig(), nil)
	if err != nil {
		log.Fatalf("Could not listen: %v", err)
	}

	log.Printf("Server listening on address %v", Addr)

	handleNewConnections(listener)
}

func handleNewConnections(listener *quic.Listener) {
	for {
		log.Printf("Waiting for new connection...")

		conn, err := listener.Accept(context.Background())
		if err != nil {
			log.Printf("Could not accept connection: %v", err)
			continue
		}

		// IMPROVE: Handling only 1 client at a time
		// Reason: So that we don't have to lock the note
		handleConnection(conn)

		// conn.CloseWithError(0, "")

		log.Printf("Connection handled!")
	}
}

func handleConnection(conn *quic.Conn) {
	// TODO: Make some mechanism for automatically sending errors relevant to the client

	if _, err := os.Stat(NoteFile); err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			err = os.WriteFile(NoteFile, []byte{}, 0644)
			if err != nil {
				log.Printf("Could not create initial note file: %v", err)
				return
			}
		} else {
			log.Printf("Could not check for note file's existance: %v", err)
			return
		}
	}

	// IMPROVE: Read the file piece by piece
	data, err := os.ReadFile(NoteFile)
	if err != nil {
		log.Printf("Could not read note: %v", err)
		return
	}

	// TODO: Use the nonblocking version
	stream, err := conn.OpenStreamSync(context.Background())
	if err != nil {
		log.Printf("Could not open stream: %v", err)
		return
	}

	// TODO: Add a timeout
	bytesWritten, err := stream.Write(data)
	if err != nil {
		log.Printf("Could not send data: %v", err)
		return
	}
	if bytesWritten != len(data) {
		// TODO: Loop instead
		panic("Not all data sent")
	}

	// Close only for writing
	err = stream.Close()
	if err != nil {
		log.Printf("Clould not close stream for writing: %v", err)
		return
	}

	log.Printf("Getting new note content...")

	data, err = lib.ReadUntilEOF(stream)
	if err != nil {
		log.Printf("Could not receive new note content: %v", err)
		return
	}
	// // TODO: Add a timeout
	// data, err = io.ReadAll(stream)
	// if err != nil {
	// 	log.Printf("Could not read data: %v", err)
	// 	return
	// }

	log.Printf("Got new note content")

	err = os.WriteFile(NoteFile, data, 0644)
	if err != nil {
		log.Printf("Could not write note: %v", err)
		// TODO: This definetely needs to be sent to the client
		return
	}

	log.Printf("Wrote new note content")
}
