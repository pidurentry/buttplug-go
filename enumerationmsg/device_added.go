package enumerationmsg

import (
	"github.com/pidurentry/buttplug-go/message"
)

//go:generate go run ../tools/message_generator.go -- $GOFILE

type DeviceAdded struct {
	ID message.Id `json:"Id"`
	*Device
}
