package lib

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
)

func StreamSendAction[T io.Writer](stream T, action Action) error {
	data := action.ToUint8()

	err := StreamSendUint8(stream, data)
	if err != nil {
		return fmt.Errorf("Could not send action: %v", err)
	}

	return nil
}

func StreamSendDatalenString[T io.Writer](stream T, data string) error {
	return StreamSendDatalenSliceByte(stream, []byte(data))
}

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

		// I don't think I've ever seen quic's `Write` return
		// both `1` and `EOF` but since I've seen it do that
		// on `Recv`, I'm putting a similar check here
		if sent > 0 {
			data = data[sent:]
			continue
		}

		if err != nil {
			return err
		}
	}

	return nil
}

// TODO: This is actually not quite right - a regular "EOF" frame
// cannot cancel out the data sent before it, but what can cancel
// it is `conn.CloseWithError`
// UPDATE: YES I HAVE JUST FUCKING TESTED THIS AND IT TURNS OUT
// QUIC'S RECEIVE METHOD CAN RETURN BOTH THE NUMBER OF BYTES
// READ AND AN ERROR SIMULTANEOUSLY
//
// It seems that if you send some data and then EOF it is not
// guaranteed that the data will be received before the EOF.
// So the receiver might first get EOF and then never get
// the actual data.
func StreamSendEOF[T io.Closer](stream T) error {
	return stream.Close()
}

func StreamSendEOFUnchecked[T io.Closer](stream T) {
	err := StreamSendEOF(stream)
	if err != nil {
		log.Print("Error: Could not send EOF on stream")
	}
}
