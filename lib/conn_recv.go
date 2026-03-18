package lib

import (
	"context"
	"errors"
	"fmt"

	"github.com/quic-go/quic-go"
)

func ConnRecvEOF(conn *quic.Conn) error {
	// IMPROVE000: We can perhaps make something better using `conn.Context().Done()`
	// and `conn.Context().Err()`
	_, err := conn.AcceptStream(context.Background())
	if err != nil {

		if appErr, ok := errors.AsType[*quic.ApplicationError](err); ok {
			if (appErr.ErrorCode == 0) && (appErr.Remote) {
				return nil
			}
		}

		// var appErr *quic.ApplicationError
		// if errors.As(err, &appErr) {
		// 	if appErr.ErrorCode == 0 && appErr.Remote {
		// 		fmt.Println("Connection closed gracefully")
		// 		return nil
		// 	}
		// }

		return fmt.Errorf("Could not receive connection EOF: %v", err)

	}

	return fmt.Errorf("Expected to receive connection EOF, instead received a new stream")
}
