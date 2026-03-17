package lib

import (
	"fmt"
	"io"

	"github.com/quic-go/quic-go"
)

// IMPROVE: Dynamically fill the note as bytes are received
// TODO: Add a timeout
func ReadUntilEOF(stream *quic.Stream) ([]byte, error) {
	bigBuf := make([]byte, 0)
	smallBuf := make([]byte, 1024)

	for {
		bytesRead, err := stream.Read(smallBuf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		bigBuf = append(bigBuf, smallBuf[:bytesRead]...)
		fmt.Printf("DBG: Got %v bytes\n", bytesRead)
	}

	fmt.Printf("DBG: read all\n")
	return bigBuf, nil
}
