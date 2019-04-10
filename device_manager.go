package buttplug

import (
	"sync"
	"time"

	"github.com/pidurentry/buttplug-go/enumerationmsg"
	"github.com/pidurentry/buttplug-go/genericdevicemsg"
	"github.com/pidurentry/buttplug-go/logging"
	"github.com/pidurentry/buttplug-go/message"
	"github.com/pidurentry/buttplug-go/statusmsg"
)

type DeviceManager interface {
	Scan(duration time.Duration) *sync.WaitGroup
	StopAll() bool
	Devices() []Device
	Raws() []Raw
	Vibrators() []Vibrate
	Linears() []Linear
	Rotators() []Rotate
	Kiiroos() []Kiiroo
	FleshlightLaunches() []FleshlightLaunch
	Lovenses() []Lovense
	VorzeA10Cyclones() []VorzeA10Cyclone
}

type deviceManager struct {
	mux      *sync.RWMutex
	handler  Handler
	scanning *sync.WaitGroup
	devices  []Device
}

func NewDeviceManager(handler Handler) DeviceManager {
	return &deviceManager{
		mux:     &sync.RWMutex{},
		handler: handler,
	}
}

func (deviceManager *deviceManager) Scan(duration time.Duration) *sync.WaitGroup {
	deviceManager.mux.Lock()
	defer deviceManager.mux.Unlock()

	if deviceManager.scanning != nil {
		return deviceManager.scanning
	}

	deviceManager.scanning = &sync.WaitGroup{}
	deviceManager.scanning.Add(1)

	go deviceManager.scan(duration, deviceManager.scanning)
	return deviceManager.scanning
}

func (deviceManager *deviceManager) scan(duration time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()

	channel := deviceManager.startScanning()
	if channel == nil {
		return
	}

	timer := time.After(duration)
scan:
	for {
		var msg Message
		select {
		case msg = <-channel:
		case msg = <-deviceManager.handler.System():
		case <-timer:
			deviceManager.stopScanning()
			break scan
		}

		if _, ok := msg.(*enumerationmsg.ScanningFinished); ok {
			break scan
		}
	}

	deviceManager.requestDeviceList()
}

func (deviceManager *deviceManager) startScanning() <-chan Message {
	channel, err := deviceManager.handler.Register(&enumerationmsg.StartScanning{message.NewId()})
	if err != nil {
		logging.GetLogger().Errorf("Failed to start scanning: %v", err)
		return nil
	}
	return channel
}

func (deviceManager *deviceManager) stopScanning() bool {
	msg, err := deviceManager.handler.Call(&enumerationmsg.StopScanning{message.NewId()})
	if err != nil {
		logging.GetLogger().Errorf("Failed to stop scanning: %v", err)
		return false
	}

	if _, ok := msg.(*statusmsg.Ok); !ok {
		logging.GetLogger().Errorf("Unexpected response to stop scanning request: %#v", msg)
		return false
	}

	return true
}

func (deviceManager *deviceManager) requestDeviceList() bool {
	msg, err := deviceManager.handler.Call(&enumerationmsg.RequestDeviceList{message.NewId()})
	if err != nil {
		logging.GetLogger().Errorf("Failed to get device list: %v", err)
		return false
	}

	deviceList, ok := msg.(*enumerationmsg.DeviceList)
	if !ok {
		logging.GetLogger().Errorf("Unexpected response device list request: %#v", msg)
		return false
	}

	deviceManager.mux.Lock()
	deviceManager.devices = make([]Device, 0)
	for _, device := range deviceList.Devices {
		deviceManager.devices = append(deviceManager.devices, deviceManager.create(device)...)
	}
	deviceManager.mux.Unlock()

	return true
}

func (deviceManager *deviceManager) create(device *enumerationmsg.Device) []Device {
	devices := make([]Device, 0)
	for msg, attributes := range device.DeviceMessages {
		switch msg {
		case "RawCmd":
			devices = append(devices, newRaw(
				device.DeviceName,
				device.DeviceIndex,
				deviceManager.handler,
			))
		case "VibrateCmd":
			devices = append(devices, newVibrate(
				device.DeviceName,
				device.DeviceIndex,
				attributes.FeatureCount,
				deviceManager.handler,
			))
		case "LinearCmd":
			devices = append(devices, newLinear(
				device.DeviceName,
				device.DeviceIndex,
				attributes.FeatureCount,
				deviceManager.handler,
			))
		case "RotateCmd":
			devices = append(devices, newRotate(
				device.DeviceName,
				device.DeviceIndex,
				attributes.FeatureCount,
				deviceManager.handler,
			))
		case "KiirooCmd":
			devices = append(devices, newKiiroo(
				device.DeviceName,
				device.DeviceIndex,
				deviceManager.handler,
			))
		case "FleshlightLaunchFW12Cmd":
			devices = append(devices, newFleshlightLaunch(
				device.DeviceName,
				device.DeviceIndex,
				deviceManager.handler,
			))
		case "LovenseCmd":
			devices = append(devices, newLovense(
				device.DeviceName,
				device.DeviceIndex,
				deviceManager.handler,
			))
		case "VorzeA10CycloneCmd":
			devices = append(devices, newVorzeA10Cyclone(
				device.DeviceName,
				device.DeviceIndex,
				deviceManager.handler,
			))
		}
	}
	return devices
}

func (deviceManager *deviceManager) StopAll() bool {
	msg, err := deviceManager.handler.Call(&genericdevicemsg.StopAllDevices{message.NewId()})
	if err != nil {
		logging.GetLogger().Errorf("Failed to stop all devices: %v", err)
		return false
	}

	if _, ok := msg.(*statusmsg.Ok); !ok {
		logging.GetLogger().Errorf("Unexpected response to stop all devices request: %#v", msg)
		return false
	}

	return true
}

func (deviceManager *deviceManager) Devices() []Device {
	deviceManager.mux.RLock()
	defer deviceManager.mux.RUnlock()
	return deviceManager.devices
}

func (deviceManager *deviceManager) Raws() []Raw {
	raws := make([]Raw, 0)
	deviceManager.deviceMap(func(device Device) {
		if raw, ok := device.(Raw); ok {
			raws = append(raws, raw)
		}
	})
	return raws
}

func (deviceManager *deviceManager) Vibrators() []Vibrate {
	vibrators := make([]Vibrate, 0)
	deviceManager.deviceMap(func(device Device) {
		if vibrate, ok := device.(Vibrate); ok {
			vibrators = append(vibrators, vibrate)
		}
	})
	return vibrators
}

func (deviceManager *deviceManager) Linears() []Linear {
	linears := make([]Linear, 0)
	deviceManager.deviceMap(func(device Device) {
		if linear, ok := device.(Linear); ok {
			linears = append(linears, linear)
		}
	})
	return linears
}

func (deviceManager *deviceManager) Rotators() []Rotate {
	rotators := make([]Rotate, 0)
	deviceManager.deviceMap(func(device Device) {
		if rotate, ok := device.(Rotate); ok {
			rotators = append(rotators, rotate)
		}
	})
	return rotators
}

func (deviceManager *deviceManager) Kiiroos() []Kiiroo {
	kiiroos := make([]Kiiroo, 0)
	deviceManager.deviceMap(func(device Device) {
		if kiiroo, ok := device.(Kiiroo); ok {
			kiiroos = append(kiiroos, kiiroo)
		}
	})
	return kiiroos
}

func (deviceManager *deviceManager) FleshlightLaunches() []FleshlightLaunch {
	launches := make([]FleshlightLaunch, 0)
	deviceManager.deviceMap(func(device Device) {
		if launche, ok := device.(FleshlightLaunch); ok {
			launches = append(launches, launche)
		}
	})
	return launches
}

func (deviceManager *deviceManager) Lovenses() []Lovense {
	lovenses := make([]Lovense, 0)
	deviceManager.deviceMap(func(device Device) {
		if lovense, ok := device.(Lovense); ok {
			lovenses = append(lovenses, lovense)
		}
	})
	return lovenses
}

func (deviceManager *deviceManager) VorzeA10Cyclones() []VorzeA10Cyclone {
	cyclones := make([]VorzeA10Cyclone, 0)
	deviceManager.deviceMap(func(device Device) {
		if cyclone, ok := device.(VorzeA10Cyclone); ok {
			cyclones = append(cyclones, cyclone)
		}
	})
	return cyclones
}

func (deviceManager *deviceManager) deviceMap(deviceMap func(Device)) {
	deviceManager.mux.RLock()
	defer deviceManager.mux.RUnlock()

	for _, device := range deviceManager.devices {
		deviceMap(device)
	}
}
