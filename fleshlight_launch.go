package buttplug

import (
	"github.com/pidurentry/buttplug-go/device"
	"github.com/pidurentry/buttplug-go/message"
	"github.com/pidurentry/buttplug-go/specificdevicemsg"
)

type FleshlightLaunch interface {
	DeviceName() device.Name
	DeviceIndex() device.Index
	FleshlightCmd(position, speed int) error
	Stop() error
}

type fleshlightLaunch struct {
	*managedDevice
}

func newFleshlightLaunch(deviceName device.Name, deviceIndex device.Index, handler Handler) FleshlightLaunch {
	return &fleshlightLaunch{
		managedDevice: &managedDevice{
			deviceName:  deviceName,
			deviceIndex: deviceIndex,
			handler:     handler,
		},
	}
}

func (launch *fleshlightLaunch) FleshlightCmd(position, speed int) error {
	return launch.send(&specificdevicemsg.FleshlightLaunchFW12Cmd{
		ID:          message.NewId(),
		DeviceIndex: launch.deviceIndex,
		Position:    position,
		Speed:       speed,
	})
}
