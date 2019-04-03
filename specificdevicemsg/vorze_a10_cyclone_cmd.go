package specificdevicemsg

import (
	"github.com/pidurentry/buttplug-go/device"
	"github.com/pidurentry/buttplug-go/message"
)

//go:generate go run ../tools/message_generator.go -- $GOFILE

type VorzeA10CycloneCmd struct {
	ID          message.Id       `json:"Id"`
	DeviceIndex device.Index     `json:"DeviceIndex"`
	Speed       int              `json:"Speed"`
	Direction   device.Direction `json:"Clockwise"`
}
