package server

import (
	"context"
	"log"

	"github.com/kuche1/cloud-note/server/filesystem"
	"github.com/quic-go/quic-go"
)

func Main(address string, filesystemStorage string) {
	fs := filesystem.NewFilesystem(filesystemStorage)

	listener, err := quic.ListenAddr(address, generateTLSConfig(), nil)
	if err != nil {
		log.Fatalf("Could not listen: %v", err)
	}

	log.Printf("Server listening on address %v", address)

	handleNewConnections(listener, fs)
}

func handleNewConnections(listener *quic.Listener, fs *filesystem.Filesystem) {
	for {
		log.Printf("Waiting for new connection...")

		conn, err := listener.Accept(context.Background())
		if err != nil {
			log.Printf("Could not accept connection: %v", err)
			continue
		}

		// IMPROVE000: Currently we're handling only 1 client at a time
		// Reason: So that we don't have to lock the note
		log.Printf("Handling connection...")
		handleNewConnection(conn, fs)

		log.Printf("Connection handled!")
	}
}
