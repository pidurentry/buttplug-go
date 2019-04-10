package buttplug

import (
	"github.com/pidurentry/buttplug-go/device"
	"github.com/pidurentry/buttplug-go/message"
	"github.com/pidurentry/buttplug-go/specificdevicemsg"
)

type Kiiroo interface {
	DeviceName() device.Name
	DeviceIndex() device.Index
	KiirooCmd(string) error
	Stop() error
}

type kiiroo struct {
	*managedDevice
}

func newKiiroo(deviceName device.Name, deviceIndex device.Index, handler Handler) Kiiroo {
	return &kiiroo{
		managedDevice: &managedDevice{
			deviceName:  deviceName,
			deviceIndex: deviceIndex,
			handler:     handler,
		},
	}
}

func (kiiroo *kiiroo) KiirooCmd(cmd string) error {
	return kiiroo.send(&specificdevicemsg.KiirooCmd{
		ID:          message.NewId(),
		DeviceIndex: kiiroo.deviceIndex,
		Command:     cmd,
	})
}
