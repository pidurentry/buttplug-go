package enumerationmsg

import "github.com/pidurentry/buttplug-go"

//go:generate go run ../message/generator.go -- $GOFILE

type DeviceRemoved struct {
	ID          buttplug.MessageId   `json:"Id"`
	DeviceIndex buttplug.DeviceIndex `json:"DeviceIndex"`
}
