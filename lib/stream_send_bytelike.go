package lib

import "io"

func StreamSendOkOrNot[T io.Writer](stream T, ok bool, reasonIfNotOk string) error {
	var bytelike uint8
	if ok {
		bytelike = BytelikeOk
	} else {
		bytelike = BytelikeNotOk
	}

	err := StreamSendUint8(stream, bytelike)
	if err != nil {
		return err
	}

	if !ok {
		err := StreamSendDatalenString(stream, reasonIfNotOk)
		if err != nil {
			return err
		}
	}

	return nil
}

func StreamSendACK[T io.Writer](stream T) error {
	return StreamSendUint8(stream, BytelikeACK)
}

func StreamSendAction[T io.Writer](stream T, action Action) error {
	data := action.ToUint8()

	err := StreamSendUint8(stream, data)
	if err != nil {
		return err
	}

	return nil
}

func StreamSendUint8[T io.Writer](stream T, data uint8) error {
	buf := []byte{data}
	return StreamSendSliceByte(stream, buf)
}
