package buttplug

import (
	"github.com/pidurentry/buttplug-go/device"
	"github.com/pidurentry/buttplug-go/genericdevicemsg"
	"github.com/pidurentry/buttplug-go/message"
)

type Raw interface {
	DeviceName() device.Name
	DeviceIndex() device.Index
	Cmd([]byte) error
	Stop() error
}

type raw struct {
	*managedDevice
}

func newRaw(deviceName device.Name, deviceIndex device.Index, handler Handler) Raw {
	return &raw{
		managedDevice: &managedDevice{
			deviceName:  deviceName,
			deviceIndex: deviceIndex,
			handler:     handler,
		},
	}
}

func (raw *raw) Cmd(cmd []byte) error {
	return raw.send(&genericdevicemsg.RawCmd{
		ID:          message.NewId(),
		DeviceIndex: raw.deviceIndex,
		Command:     cmd,
	})
}
