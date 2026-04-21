// Package for network-related shorthands for the server

// IMPROVE001:
// Make equivalent functions but for the client ?

package srvnet

import (
	"github.com/kuche1/cloud-note/lib"
	"github.com/kuche1/cloud-note/server/config"
	"github.com/quic-go/quic-go"
)

func StreamRecvNotename(stream *quic.Stream) (string, error) {
	// IMPROVE000: Send a readable message for the end user
	// if the disconnect reason is the length
	return lib.StreamRecvDatalenString(stream, config.NoteNameMaxLength)
}

func StreamRecvNotecontent(stream *quic.Stream) (string, error) {
	// IMPROVE000: Send a readable message for the end user
	// if the disconnect reason is the length
	return lib.StreamRecvDatalenString(stream, config.NoteContentsMaxLength)
}

func StreamRecvPassword(stream *quic.Stream) (string, error) {
	// IMPROVE000: Send a readable message for the end user
	// if the disconnect reason is the length
	return lib.StreamRecvDatalenString(stream, config.PasswordMaxLength)
}
