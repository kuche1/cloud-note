package lib

// // IMPROVE: Dynamically fill the note as bytes are received
// // TODO: Add a timeout
// func ReadUntilEOF(stream *quic.Stream) ([]byte, error) {
// 	bigBuf := make([]byte, 0)
// 	smallBuf := make([]byte, 1024)

// 	for {
// 		bytesRead, err := stream.Read(smallBuf)
// 		if err != nil {
// 			if err == io.EOF {
// 				fmt.Printf("DBG: got EOF\n")
// 				break
// 			}
// 			// if err.Error() == "Application error 0x0 (remote)" {
// 			// 	// TODO: See if we can make this better
// 			// 	fmt.Printf("DBG: connection closed\n")
// 			// 	break
// 			// }
// 			fmt.Printf("DBG: got some other error: %v\n", err)
// 			return nil, err
// 		}
// 		fmt.Printf("DBG: got some data: %v\n", smallBuf[:bytesRead])
// 		bigBuf = append(bigBuf, smallBuf[:bytesRead]...)
// 		fmt.Printf("DBG: Got %v bytes\n", bytesRead)
// 	}

// 	fmt.Printf("DBG: read all: %v\n", bigBuf)
// 	return bigBuf, nil
// }
