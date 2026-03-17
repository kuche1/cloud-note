package lib

import (
	"encoding/binary"

	"github.com/quic-go/quic-go"
)

// IMPROVE: Make a variant that sends the data chunk by chunk rather than all of it at once
func SendDatalenSliceByte(stream *quic.Stream, data []byte) error {
	err := SendUint64(stream, uint64(len(data)))
	if err != nil {
		return err
	}

	SendSliceByte(stream, data)

	return nil
}

func SendUint64(stream *quic.Stream, data uint64) error {
	buf := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	binary.BigEndian.PutUint64(buf, data)
	return SendSliceByte(stream, buf)
}

func SendSliceByte(stream *quic.Stream, data []byte) error {
	for len(data) > 0 {
		sent, err := stream.Write(data)
		if err != nil {
			return err
		}
		data = data[sent:]
	}
	return nil
}

func SendEOF(stream *quic.Stream) error {
	return stream.Close()
}

// // UNTESTED
// func SendDatalenSliceByte_RecvEOF(stream *quic.Stream, data []byte) error {
// 	err := SendDatalenSliceByte(stream, data)
// 	if err != nil {
// 		return err
// 	}

// 	err = RecvEOF(stream)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
