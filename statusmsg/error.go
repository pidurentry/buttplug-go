package statusmsg

import (
	"fmt"

	"github.com/pidurentry/buttplug-go"
)

type ErrorCode int
type ErrorMessage string

const (
	ERROR_UNKNOWN ErrorCode = iota
	ERROR_INIT
	ERROR_PING
	ERROR_MSG
	ERROR_DEVICE
)

type Error struct {
	id           buttplug.MessageId
	errorCode    ErrorCode    `json:"ErrorCode"`
	errorMessage ErrorMessage `json:"ErrorMessage"`
}

func (error *Error) Id() buttplug.MessageId {
	return error.id
}

func (error *Error) Error() string {
	return fmt.Sprintf("%d: %s", error.errorCode, error.errorMessage)
}

func (error *Error) ErrorCode() ErrorCode {
	return error.errorCode
}

func (error *Error) ErrorMessage() ErrorMessage {
	return error.errorMessage
}
