// TODO: None of this has been tested

package server

import (
	"context"
	"io"
	"log"

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
	stream, err := conn.OpenStreamSync(context.Background())
	if err != nil {
		log.Printf("Could not open stream: %v", err)
		return
	}

	// TODO: Send the actual file content
	// TODO: Add a timeout
	data := []byte("asd fgh\ngfdsf\ncgsrevcgsre resgcfrsegcser ggsrescgsresc\ntsrcghetrcgstre")
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
	stream.Close()

	// TODO: Add a timeout
	data, err = io.ReadAll(stream)
	if err != nil {
		log.Printf("Could not read data: %v", err)
		return
	}
	// TODO: write new `data` to file
}
