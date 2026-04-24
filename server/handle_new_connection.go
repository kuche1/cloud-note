package server

import (
	"fmt"
	"log"

	"github.com/kuche1/cloud-note/server/action"
	"github.com/kuche1/cloud-note/server/filesystem"
	"github.com/quic-go/quic-go"
)

func handleNewConnection(conn *quic.Conn, fs *filesystem.Filesystem) {
	errString, errCode := action.HandleAction(conn, fs)
	if errString != nil {
		log.Printf("Error: %v\n", errString)

		err := conn.CloseWithError(errCode, errString.Error())
		if err != nil {
			fmt.Printf("Could not close connection with error: %v", err)
			return
		}
	}
}
