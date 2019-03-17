package statusmsg

import "github.com/pidurentry/buttplug-go"

//go:generate go run ../message/generator.go -- $GOFILE

type TestString string

type Test struct {
	ID         buttplug.MessageId `json:"Id"`
	TestString TestString         `json:"TestString"`
}
