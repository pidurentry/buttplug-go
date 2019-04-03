package buttplug

import (
	"sync"
	"time"

	"github.com/pidurentry/buttplug-go/enumerationmsg"
	"github.com/pidurentry/buttplug-go/logging"
	"github.com/pidurentry/buttplug-go/message"
	"github.com/pidurentry/buttplug-go/statusmsg"
)

type DeviceManager interface {
	Scan(duration time.Duration) *sync.WaitGroup
	Devices() []*enumerationmsg.Device
}

type deviceManager struct {
	mux      *sync.Mutex
	handler  Handler
	scanning *sync.WaitGroup
	devices  []*enumerationmsg.Device
}

func NewDeviceManager(handler Handler) DeviceManager {
	return &deviceManager{
		mux:     &sync.Mutex{},
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
	deviceManager.devices = deviceList.Devices
	deviceManager.mux.Unlock()

	return true
}

func (deviceManager *deviceManager) Devices() []*enumerationmsg.Device {
	deviceManager.mux.Lock()
	defer deviceManager.mux.Unlock()
	return deviceManager.devices
}
