package statusmsg

import (
	"fmt"

	"github.com/pidurentry/buttplug-go"
)

//go:generate go run ../message/generator.go -- $GOFILE

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
	ID           buttplug.MessageId `json:"Id"`
	ErrorCode    ErrorCode          `json:"ErrorCode"`
	ErrorMessage ErrorMessage       `json:"ErrorMessage"`
}

func (error *Error) Error() string {
	return fmt.Sprintf("%d: %s", error.ErrorCode, error.ErrorMessage)
}
