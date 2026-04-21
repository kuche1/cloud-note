package lib

import (
	"encoding/binary"
	"fmt"
	"io"
)

func StreamRecvAction[T io.Reader](stream T) (Action, error) {
	data, err := StreamRecvUint8(stream)
	if err != nil {
		return 0, fmt.Errorf("Could not receive action: %v", err)
	}

	action, err := Action(0).FromUint8(data)
	if err != nil {
		return 0, fmt.Errorf("Could not convert uint8 to action: %v", err)
	}

	return action, nil
}

func StreamRecvDatalenString[T io.Reader](stream T, maxLength uint64) (string, error) {
	data, err := StreamRecvDatalenSliceByte(stream, maxLength)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func StreamRecvDatalenSliceByte[T io.Reader](stream T, maxLength uint64) ([]byte, error) {
	length, err := StreamRecvUint64(stream)
	if err != nil {
		return nil, err
	}

	if length > maxLength {
		return nil, fmt.Errorf("Item size (%v) exceeds maximum allowed size (%v)", length, maxLength)
	}

	data, err := StreamRecvSliceByte(stream, length)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func StreamRecvUint8[T io.Reader](stream T) (uint8, error) {
	buf, err := StreamRecvSliceByte(stream, 1)
	if err != nil {
		return 0, err
	}
	return buf[0], nil
}

func StreamRecvUint64[T io.Reader](stream T) (uint64, error) {
	buf, err := StreamRecvSliceByte(stream, 8)
	if err != nil {
		return 0, err
	}
	bits := binary.BigEndian.Uint64(buf)
	return bits, nil
}

func StreamRecvSliceByte[T io.Reader](stream T, length uint64) ([]byte, error) {
	data := make([]byte, length)
	// Allocating this on every receive should not be a big deal since
	// the network should be much slower than the allocation of the memory
	// (That is to say, IF an actual allocation is performed here)

	buf := data
	for len(buf) > 0 {
		read, err := stream.Read(buf)

		// OH MY FUCKING GOD
		// QUICK RETURNS BOTH `1` AND `EOF` HERE WHAT THE FUCK
		if read > 0 {
			buf = buf[read:]
			continue
		}

		if err != nil {
			return nil, err
		}
	}

	return data, nil
}

// This is a bit dangeround - an EOF has to be sent only once,
// but it can be read many times
func StreamRecvEOF[T io.Reader](stream T) error {
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
