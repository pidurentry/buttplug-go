package enumerationmsg

import "github.com/pidurentry/buttplug-go/device"

type Device struct {
	DeviceName     device.Name     `json:"DeviceName"`
	DeviceIndex    device.Index    `json:"DeviceIndex"`
	DeviceMessages device.Messages `json:"DeviceMessages"`
}
