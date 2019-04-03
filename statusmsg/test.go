package statusmsg

import "github.com/pidurentry/buttplug-go/message"

//go:generate go run ../tools/message_generator.go -- $GOFILE

type TestString string

type Test struct {
	ID         message.Id `json:"Id"`
	TestString TestString `json:"TestString"`
}
