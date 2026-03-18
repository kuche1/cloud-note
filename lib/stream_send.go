package lib

import (
	"encoding/binary"
	"io"
)

func StreamSendDatalenString[T io.Writer](stream T, data string) error {
	return StreamSendDatalenSliceByte(stream, []byte(data))
}

// IMPROVE000: Make a variant that sends the data chunk by chunk rather than all of it at once
func StreamSendDatalenSliceByte[T io.Writer](stream T, data []byte) error {
	err := StreamSendUint64(stream, uint64(len(data)))
	if err != nil {
		return err
	}

	StreamSendSliceByte(stream, data)

	return nil
}

func StreamSendUint8[T io.Writer](stream T, data uint8) error {
	buf := []byte{data}
	return StreamSendSliceByte(stream, buf)
}

func StreamSendUint64[T io.Writer](stream T, data uint64) error {
	buf := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	binary.BigEndian.PutUint64(buf, data)
	return StreamSendSliceByte(stream, buf)
}

func StreamSendSliceByte[T io.Writer](stream T, data []byte) error {
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
func StreamSendEOF[T io.Closer](stream T) error {
	return stream.Close()
}
