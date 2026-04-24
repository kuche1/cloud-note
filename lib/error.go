package lib

import (
	"errors"

	"github.com/quic-go/quic-go"
)

const ErrorCodeTimeoutDuringActionRead = 1

// TODO: this is actually a client-only function, it shouldnt be here
func ErrorIsTimeout(err error) bool {
	_, ok := errors.AsType[*quic.IdleTimeoutError](err)
	if ok {
		return true
	}

	appErr, ok := errors.AsType[*quic.ApplicationError](err)
	if ok {
		if appErr.ErrorCode == ErrorCodeTimeoutDuringActionRead {
			return true
		}
		return false
	}

	return false
}
