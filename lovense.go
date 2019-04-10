package buttplug

import (
	"github.com/pidurentry/buttplug-go/device"
	"github.com/pidurentry/buttplug-go/message"
	"github.com/pidurentry/buttplug-go/specificdevicemsg"
)

type Lovense interface {
	DeviceName() device.Name
	DeviceIndex() device.Index
	LovenseCmd(string) error
	Stop() error
}

type lovense struct {
	*managedDevice
}

func newLovense(deviceName device.Name, deviceIndex device.Index, handler Handler) Lovense {
	return &lovense{
		managedDevice: &managedDevice{
			deviceName:  deviceName,
			deviceIndex: deviceIndex,
			handler:     handler,
		},
	}
}

func (lovense *lovense) LovenseCmd(cmd string) error {
	return lovense.send(&specificdevicemsg.LovenseCmd{
		ID:          message.NewId(),
		DeviceIndex: lovense.deviceIndex,
		Command:     cmd,
	})
}
