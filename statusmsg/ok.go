package statusmsg

import "github.com/pidurentry/buttplug-go"

//go:generate go run ../message/generator.go -- $GOFILE

type Ok struct {
	ID buttplug.MessageId `json:"Id"`
}
