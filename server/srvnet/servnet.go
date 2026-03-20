// Package for network-related shorthands for the server

// TODO:
// Make equivalent functions but for the client ?

package srvnet

import (
	"github.com/kuche1/cloud-note/lib"
	"github.com/kuche1/cloud-note/server/config"
	"github.com/quic-go/quic-go"
)

func ChanRecvNotenameEOF(conn *quic.Conn) (string, error) {
	// IMPROVE000: Send a readable message for the end user
	// if the disconnect reason is the length
	return lib.ChanRecvStringEOF(conn, config.NoteNameMaxLength)
}

func ChanRecvNotecontentEOF(conn *quic.Conn) (string, error) {
	// IMPROVE000: Send a readable message for the end user
	// if the disconnect reason is the length
	return lib.ChanRecvStringEOF(conn, config.NoteContentsMaxLength)
}
