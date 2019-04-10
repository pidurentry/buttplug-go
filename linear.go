package buttplug

import (
	"github.com/pidurentry/buttplug-go/device"
	"github.com/pidurentry/buttplug-go/genericdevicemsg"
	"github.com/pidurentry/buttplug-go/message"
)

type Linear interface {
	DeviceName() device.Name
	DeviceIndex() device.Index
	VectorCount() device.FeatureCount
	Vector(...device.Vector) error
	Stop() error
}

type linear struct {
	*managedDevice
	featureCount device.FeatureCount
}

func newLinear(deviceName device.Name, deviceIndex device.Index, featureCount device.FeatureCount, handler Handler) Linear {
	return &linear{
		managedDevice: &managedDevice{
			deviceName:  deviceName,
			deviceIndex: deviceIndex,
			handler:     handler,
		},
		featureCount: featureCount,
	}
}

func (linear *linear) VectorCount() device.FeatureCount {
	return linear.featureCount
}

func (linear *linear) Vector(vectors ...device.Vector) error {
	return linear.send(&genericdevicemsg.LinearCmd{
		ID:          message.NewId(),
		DeviceIndex: linear.deviceIndex,
		Vectors:     vectors,
	})
}
