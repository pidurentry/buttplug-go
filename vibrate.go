package buttplug

import (
	"github.com/pidurentry/buttplug-go/device"
	"github.com/pidurentry/buttplug-go/genericdevicemsg"
	"github.com/pidurentry/buttplug-go/message"
)

type Vibrate interface {
	DeviceName() device.Name
	DeviceIndex() device.Index
	VibrateCount() device.FeatureCount
	Vibrate(...device.Speed) error
	Stop() error
}

type vibrate struct {
	*managedDevice
	featureCount device.FeatureCount
}

func newVibrate(deviceName device.Name, deviceIndex device.Index, featureCount device.FeatureCount, handler Handler) Vibrate {
	return &vibrate{
		managedDevice: &managedDevice{
			deviceName:  deviceName,
			deviceIndex: deviceIndex,
			handler:     handler,
		},
		featureCount: featureCount,
	}
}

func (vibrate *vibrate) VibrateCount() device.FeatureCount {
	return vibrate.featureCount
}

func (vibrate *vibrate) Vibrate(speeds ...device.Speed) error {
	return vibrate.send(&genericdevicemsg.VibrateCmd{
		ID:          message.NewId(),
		DeviceIndex: vibrate.deviceIndex,
		Speeds:      speeds,
	})
}
