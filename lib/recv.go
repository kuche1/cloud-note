package lib

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/quic-go/quic-go"
)

func RecvDatalenSliceByte(stream *quic.Stream) ([]byte, error) {
	length, err := RecvUint64(stream)
	if err != nil {
		return nil, err
	}

	data, err := RecvSliceByte(stream, length)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func RecvUint64(stream *quic.Stream) (uint64, error) {
	buf, err := RecvSliceByte(stream, 8)
	if err != nil {
		return 0, err
	}
	bits := binary.BigEndian.Uint64(buf)
	return bits, nil
}

func RecvSliceByte(stream *quic.Stream, length uint64) ([]byte, error) {
	data := make([]byte, length)
	// Allocating this on every receive should not be a big deal since
	// the network should be much slower than the allocation of the memory

	buf := data
	for len(buf) > 0 {
		read, err := stream.Read(buf)
		if err != nil {
			return nil, err
		}
		buf = buf[read:]
	}

	return data, nil
}

// This is a bit dangeround - an EOF has to be sent only once,
// but it can be read many times
func RecvEOF(stream *quic.Stream) error {
	buf := []byte{0}

	_, err := stream.Read(buf)
	if err != nil {
		if err == io.EOF {
			return nil
		} else {
			return err
		}
	}

	return fmt.Errorf("Expected to receive EOF, instead got valid data")
}
