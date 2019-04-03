package genericdevicemsg

import (
	"github.com/pidurentry/buttplug-go/device"
	"github.com/pidurentry/buttplug-go/message"
)

//go:generate go run ../tools/message_generator.go -- $GOFILE

type RawCmd struct {
	ID          message.Id   `json:"Id"`
	DeviceIndex device.Index `json:"DeviceIndex"`
	Command     []byte       `json:"Command"`
}
