package buttplug

import (
	"github.com/pidurentry/buttplug-go/device"
	"github.com/pidurentry/buttplug-go/message"
	"github.com/pidurentry/buttplug-go/specificdevicemsg"
)

type VorzeA10Cyclone interface {
	DeviceName() device.Name
	DeviceIndex() device.Index
	VorzeA10CycloneCmd(speed int, direction device.Direction) error
	Stop() error
}

type vorzeA10Cyclone struct {
	*managedDevice
}

func newVorzeA10Cyclone(deviceName device.Name, deviceIndex device.Index, handler Handler) VorzeA10Cyclone {
	return &vorzeA10Cyclone{
		managedDevice: &managedDevice{
			deviceName:  deviceName,
			deviceIndex: deviceIndex,
			handler:     handler,
		},
	}
}

func (cyclone *vorzeA10Cyclone) VorzeA10CycloneCmd(speed int, direction device.Direction) error {
	return cyclone.send(&specificdevicemsg.VorzeA10CycloneCmd{
		ID:          message.NewId(),
		DeviceIndex: cyclone.deviceIndex,
		Speed:       speed,
		Direction:   direction,
	})
}
