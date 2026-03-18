package lib

import (
	"encoding/binary"
	"io"
)

// IMPROVE: Make a variant that sends the data chunk by chunk rather than all of it at once
func SendDatalenSliceByte[T io.Writer](stream T, data []byte) error {
	err := SendUint64(stream, uint64(len(data)))
	if err != nil {
		return err
	}

	SendSliceByte(stream, data)

	return nil
}

func SendUint8[T io.Writer](stream T, data uint8) error {
	buf := []byte{data}
	return SendSliceByte(stream, buf)
}

func SendUint64[T io.Writer](stream T, data uint64) error {
	buf := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	binary.BigEndian.PutUint64(buf, data)
	return SendSliceByte(stream, buf)
}

func SendSliceByte[T io.Writer](stream T, data []byte) error {
	for len(data) > 0 {
		sent, err := stream.Write(data)
		if err != nil {
			return err
		}
		data = data[sent:]
	}
	return nil
}

// It seems that if you send some data and then EOF it is not
// guaranteed that the data will be received before the EOF.
// So the receiver might first get EOF and then never get
// the actual data.
func SendEOF[T io.Closer](stream T) error {
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
