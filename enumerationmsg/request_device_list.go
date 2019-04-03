package enumerationmsg

import "github.com/pidurentry/buttplug-go"

//go:generate go run ../message/generator.go -- $GOFILE

type RequestDeviceList struct {
	ID buttplug.MessageId `json:"Id"`
}
