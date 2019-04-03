package enumerationmsg

import "github.com/pidurentry/buttplug-go"

//go:generate go run ../message/generator.go -- $GOFILE

type DeviceList struct {
	ID      buttplug.MessageId `json:"Id"`
	Devices []*buttplug.Device `json:"Devices"`
}
