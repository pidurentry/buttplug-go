package buttplug

import (
	"github.com/pidurentry/buttplug-go/device"
	"github.com/pidurentry/buttplug-go/genericdevicemsg"
	"github.com/pidurentry/buttplug-go/message"
)

type Rotate interface {
	DeviceName() device.Name
	DeviceIndex() device.Index
	RotateCount() device.FeatureCount
	Rotate(...device.Rotation) error
	Stop() error
}

type rotate struct {
	*managedDevice
	featureCount device.FeatureCount
}

func newRotate(deviceName device.Name, deviceIndex device.Index, featureCount device.FeatureCount, handler Handler) Rotate {
	return &rotate{
		managedDevice: &managedDevice{
			deviceName:  deviceName,
			deviceIndex: deviceIndex,
			handler:     handler,
		},
		featureCount: featureCount,
	}
}

func (rotate *rotate) RotateCount() device.FeatureCount {
	return rotate.featureCount
}

func (rotate *rotate) Rotate(rotations ...device.Rotation) error {
	return rotate.send(&genericdevicemsg.RotateCmd{
		ID:          message.NewId(),
		DeviceIndex: rotate.deviceIndex,
		Rotations:   rotations,
	})
}
