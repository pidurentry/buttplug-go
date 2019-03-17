package statusmsg

import "github.com/pidurentry/buttplug-go"

//go:generate go run ../message/generator.go -- $GOFILE

type Ping struct {
	ID buttplug.MessageId `json:"Id"`
}
