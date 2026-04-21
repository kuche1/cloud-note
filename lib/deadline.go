package lib

import (
	"time"

	"github.com/quic-go/quic-go"
)

func DeadlineSet(stream *quic.Stream, timeoutFromNow time.Duration) error {
	return stream.SetReadDeadline(time.Now().Add(timeoutFromNow))
}

func DeadlineClear(stream *quic.Stream) error {
	// my tests show that this sets the deadline to the default
	// (rather than "no deadline")
	return stream.SetReadDeadline(time.Time{})
}
