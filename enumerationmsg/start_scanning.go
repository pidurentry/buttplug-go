package enumerationmsg

import "github.com/pidurentry/buttplug-go"

//go:generate go run ../message/generator.go -- $GOFILE

type StartScanning struct {
	ID buttplug.MessageId `json:"Id"`
}
