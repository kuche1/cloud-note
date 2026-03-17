// IMPROVE: Call `conn.CloseWithError(0, "Info")` in case of failure

package server

import (
	"context"
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

		// IMPROVE: Currently we're handling only 1 client at a time
		// Reason: So that we don't have to lock the note
		log.Printf("Handling connection...")
		handleConnection(conn)

		log.Printf("Connection handled!")
	}
}
