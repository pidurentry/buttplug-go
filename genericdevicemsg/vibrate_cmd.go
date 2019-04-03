package genericdevicemsg

import (
	"github.com/pidurentry/buttplug-go/device"
	"github.com/pidurentry/buttplug-go/message"
)

//go:generate go run ../tools/message_generator.go -- $GOFILE

type VibrateCmd struct {
	ID          message.Id     `json:"Id"`
	DeviceIndex device.Index   `json:"DeviceIndex"`
	Speeds      []device.Speed `json:"Speeds"`
}
