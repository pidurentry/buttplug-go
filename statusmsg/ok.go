package statusmsg

import "github.com/pidurentry/buttplug-go/message"

//go:generate go run ../tools/message_generator.go -- $GOFILE

type Ok struct {
	ID message.Id `json:"Id"`
}
