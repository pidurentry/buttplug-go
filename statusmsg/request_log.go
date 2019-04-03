package statusmsg

import "github.com/pidurentry/buttplug-go/message"

//go:generate go run ../tools/message_generator.go -- $GOFILE

type RequestLog struct {
	ID       message.Id `json:"Id"`
	LogLevel LogLevel   `json:"LogLevel"`
}
