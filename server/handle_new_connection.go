package server

import (
	"fmt"
	"log"

	"github.com/kuche1/cloud-note/server/action"
	"github.com/quic-go/quic-go"
)

func handleNewConnection(conn *quic.Conn) {
	err := action.HandleAction(conn)
	if err != nil {
		log.Printf("Error: %v\n", err)

		err := conn.CloseWithError(0, err.Error())
		if err != nil {
			fmt.Printf("Could not close connection with error: %v", err)
			return
		}
	}
}
