package lib

import (
	"fmt"
	"io"
)

func StreamRecvOkOrNot[T io.Reader](stream T, refusalMaxlen uint64) (_ok bool, _refusal string, _err error) {
	data, err := StreamRecvUint8(stream)
	if err != nil {
		return false, "", err
	}

	if data == BytelikeOk {
		return true, "", nil
	}

	if data != BytelikeNotOk {
		return false, "", fmt.Errorf("Expected ok (%v) or not ok (%v), instead got %v", BytelikeOk, BytelikeNotOk, data)
	}

	refusal, err := StreamRecvDatalenString(stream, refusalMaxlen)
	if err != nil {
		return false, "", err
	}

	return false, refusal, nil
}

func StreamRecvACK[T io.Reader](stream T) error {
	data, err := StreamRecvUint8(stream)
	if err != nil {
		return err
	}

	if data != BytelikeACK {
		return fmt.Errorf("Did not receive ACK (%v), instead got %v", BytelikeACK, data)
	}

	return nil
}

func StreamRecvAction[T io.Reader](stream T) (Action, error) {
	data, err := StreamRecvUint8(stream)
	if err != nil {
		return 0, err
	}

	action, err := Action(0).FromUint8(data)
	if err != nil {
		return 0, err
	}

	return action, nil
}

func StreamRecvUint8[T io.Reader](stream T) (uint8, error) {
	buf, err := StreamRecvSliceByte(stream, 1)
	if err != nil {
		return 0, err
	}
	return buf[0], nil
}
