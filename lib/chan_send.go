package lib

import (
	"fmt"

	"github.com/quic-go/quic-go"
)

func ChanSendSliceStringEOF(conn *quic.Conn, data []string) error {
	stream, err := conn.OpenStream()
	if err != nil {
		return err
	}

	err = StreamSendUint64(stream, uint64(len(data)))
	if err != nil {
		return err
	}

	for _, item := range data {
		err = StreamSendDatalenString(stream, item)
		if err != nil {
			return err
		}
	}

	err = StreamRecvEOF(stream)
	if err != nil {
		return err
	}

	err = StreamSendEOF(stream)
	if err != nil {
		return err
	}

	return nil
}

func ChanSendActionEOF(conn *quic.Conn, action Action) error {
	data := action.ToUint8()

	err := ChanSendUint8EOF(conn, data)
	if err != nil {
		return fmt.Errorf("Could not send action: %v", err)
	}

	return nil
}

func ChanSendUint8EOF(conn *quic.Conn, data uint8) error {
	stream, err := conn.OpenStream()
	if err != nil {
		return fmt.Errorf("Could not open stream: %v", err)
	}

	err = StreamSendUint8(stream, data)
	if err != nil {
		return fmt.Errorf("Clould not send uint8: %v", err)
	}

	err = StreamRecvEOF(stream)
	if err != nil {
		return fmt.Errorf("Could not receive EOF: %v", err)
	}

	err = StreamSendEOF(stream)
	if err != nil {
		return fmt.Errorf("Could not send EOF: %v", err)
	}

	return nil
}

func ChanSendStringEOF(conn *quic.Conn, data string) error {
	return ChanSendDatalenSliceByteEOF(conn, []byte(data))
}

func ChanSendDatalenSliceByteEOF(conn *quic.Conn, data []byte) error {
	// log.Printf("Creating channel")

	stream, err := conn.OpenStream()
	if err != nil {
		return fmt.Errorf("Could not open stream: %v", err)
	}

	// log.Printf("Sending datalen slice byte | %v", data)

	err = StreamSendDatalenSliceByte(stream, data)
	if err != nil {
		return fmt.Errorf("Clould not send data: %v", err)
	}

	// log.Printf("Receiving EOF")

	err = StreamRecvEOF(stream)
	if err != nil {
		return fmt.Errorf("Could not receive EOF: %v", err)
	}

	// log.Printf("Sending EOF")

	err = StreamSendEOF(stream)
	if err != nil {
		return fmt.Errorf("Could not send EOF: %v", err)
	}

	return nil
}

func ChanSendEOF(conn *quic.Conn) error {
	streamSend, err := conn.OpenUniStream()
	if err != nil {
		return fmt.Errorf("Could not open stream send: %v", err)
	}

	err = StreamSendEOF(streamSend)
	if err != nil {
		return fmt.Errorf("Could not send EOF: %v", err)
	}

	return nil
}
