package lib

// func DeadlineSet(stream *quic.Stream, timeoutFromNow time.Duration) error {
// 	return stream.SetReadDeadline(time.Now().Add(timeoutFromNow))
// }

// func DeadlineClear(stream *quic.Stream) error {
// 	// UPDATE:
// 	// actually, when writing this I didn't know of QUIC's
// 	// own timeout mechanism, so in reality this might actually
// 	// remove the deadline altogether (which is not necessarily
// 	// a problem, it's just that actually removes it unlike
// 	// what the comment below suggests)
// 	//
// 	// my tests show that this sets the deadline to the default
// 	// (rather than "no deadline")
// 	return stream.SetReadDeadline(time.Time{})
// }
