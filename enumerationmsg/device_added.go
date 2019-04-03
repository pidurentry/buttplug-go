package enumerationmsg

import "github.com/pidurentry/buttplug-go"

//go:generate go run ../message/generator.go -- $GOFILE

type DeviceAdded struct {
	ID buttplug.MessageId `json:"Id"`
	*buttplug.Device
}
