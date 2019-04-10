package buttplug

import (
	"github.com/pidurentry/buttplug-go/device"
	"github.com/pidurentry/buttplug-go/genericdevicemsg"
	"github.com/pidurentry/buttplug-go/message"
	"github.com/pidurentry/buttplug-go/statusmsg"
)

type Device interface {
	DeviceName() device.Name
	DeviceIndex() device.Index
	Stop() error
}

type managedDevice struct {
	deviceName  device.Name
	deviceIndex device.Index
	handler     Handler
}

func (managedDevice *managedDevice) DeviceName() device.Name {
	return managedDevice.deviceName
}

func (managedDevice *managedDevice) DeviceIndex() device.Index {
	return managedDevice.deviceIndex
}

func (managedDevice *managedDevice) Stop() error {
	return managedDevice.send(&genericdevicemsg.StopDeviceCmd{
		ID:          message.NewId(),
		DeviceIndex: managedDevice.deviceIndex,
	})
}

func (managedDevice *managedDevice) send(msg Message) error {
	response, err := managedDevice.handler.Call(msg)

	if err != nil {
		switch err := err.(type) {
		case *statusmsg.Error:
			if err.ErrorCode == statusmsg.ERROR_DEVICE {
				return &DeviceError{string(err.ErrorMessage)}
			}
		}
		return err
	}

	if _, ok := response.(*statusmsg.Ok); !ok {
		return &CommandFailure{}
	}

	return nil
}
