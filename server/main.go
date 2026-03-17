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
		conn, err := listener.Accept(context.Background())
		if err != nil {
			log.Printf("Could not accept connection: %v", err)
			continue
		}

		// IMPROVE: Handling only 1 client at a time
		// Reason: So that we don't have to lock the note
		handleConnection(conn)

		conn.CloseWithError(0, "")
	}
}

func handleConnection(conn *quic.Conn) {
	// TODO: Add a timeout
	stream, err := conn.AcceptStream(context.Background()) // TODO: The code should get stuck here since the client won't actually send any data on the stream
	if err != nil {
		log.Printf("Could not accept stream: %v", err)
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
