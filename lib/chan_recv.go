package lib

import (
	"context"
	"fmt"

	"github.com/quic-go/quic-go"
)

func ChanRecvSliceStringEOF(conn *quic.Conn) ([]string, error) {
	fmt.Printf("DBG: accept stream\n")

	stream, err := conn.AcceptStream(context.Background())
	if err != nil {
		return nil, err
	}

	fmt.Printf("DBG: recv uint64\n")

	numberOfItems, err := StreamRecvUint64(stream)
	if err != nil {
		return nil, err
	}

	data := make([]string, 0, numberOfItems)

	for range numberOfItems {
		fmt.Printf("DBG: recv item [%v/%v]\n", "?", numberOfItems)

		item, err := StreamRecvDatalenString(stream)
		if err != nil {
			return nil, err
		}
		data = append(data, item)
	}

	fmt.Printf("DBG: send EOF\n")

	err = StreamSendEOF(stream)
	if err != nil {
		return nil, err
	}

	fmt.Printf("DBG: recv EOF\n")

	err = StreamRecvEOF(stream)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func ChanRecvActionEOF(conn *quic.Conn) (Action, error) {
	data, err := ChanRecvUint8EOF(conn)
	if err != nil {
		return 0, fmt.Errorf("Could not receive action: %v", err)
	}

	action, err := Action(0).FromUint8(data)
	if err != nil {
		return 0, fmt.Errorf("Could not convert uint8 to action: %v", err)
	}

	return action, nil
}

func ChanRecvUint8EOF(conn *quic.Conn) (uint8, error) {
	stream, err := conn.AcceptStream(context.Background())
	if err != nil {
		return 0, fmt.Errorf("Could not accept stream: %v", err)
	}

	data, err := StreamRecvUint8(stream)
	if err != nil {
		return 0, fmt.Errorf("Clould not receive uint8: %v", err)
	}

	err = StreamSendEOF(stream)
	if err != nil {
		return 0, fmt.Errorf("Could not send EOF: %v", err)
	}

	err = StreamRecvEOF(stream)
	if err != nil {
		return 0, fmt.Errorf("Could not receive EOF: %v", err)
	}

	return data, nil
}

func ChanRecvDatalenSliceByteEOF(conn *quic.Conn) ([]byte, error) {
	// log.Printf("Accepting stream")

	stream, err := conn.AcceptStream(context.Background())
	if err != nil {
		return nil, fmt.Errorf("Could not accept stream: %v", err)
	}

	// log.Printf("Receiving datalen slice byte")

	data, err := StreamRecvDatalenSliceByte(stream)
	if err != nil {
		return nil, fmt.Errorf("Clould not receive data: %v", err)
	}

	// log.Printf("Received datalen slice byte | %v", data)

	// log.Printf("Sending EOF")

	err = StreamSendEOF(stream)
	if err != nil {
		return nil, fmt.Errorf("Could not send EOF: %v", err)
	}

	// log.Printf("Receiving EOF")

	err = StreamRecvEOF(stream)
	if err != nil {
		return nil, fmt.Errorf("Could not receive EOF: %v", err)
	}

	// log.Printf("Received EOF")

	return data, nil
}

func ChanRecvEOF(conn *quic.Conn) error {
	streamRecv, err := conn.AcceptUniStream(context.Background())
	if err != nil {
		return fmt.Errorf("Could not accept stream recv: %v", err)
	}

	err = StreamRecvEOF(streamRecv)
	if err != nil {
		return fmt.Errorf("Could not receive EOF: %v", err)
	}

	return nil
}
