package enumerationmsg

// Auto generated file - do not directly edit!
import "github.com/pidurentry/buttplug-go"

func init() {
    buttplug.MessageRepository["DeviceAdded"] = func() interface{} { return &DeviceAdded{} }
}

func (msg *DeviceAdded) Id() buttplug.MessageId {
    return msg.ID
}

func (msg *DeviceAdded) Serilize() interface{} {
    return map[string]*DeviceAdded{
        "DeviceAdded": msg,
    }
}

