package statusmsg

import "github.com/pidurentry/buttplug-go"

//go:generate go run ../message/generator.go -- $GOFILE

type RequestLog struct {
	ID       buttplug.MessageId `json:"Id"`
	LogLevel LogLevel           `json:"LogLevel"`
}
