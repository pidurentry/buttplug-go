package enumerationmsg

// Auto generated file - do not directly edit!
import "github.com/pidurentry/buttplug-go/message"

func init() {
    message.Repository["DeviceAdded"] = func() interface{} { return &DeviceAdded{} }
}

func (msg *DeviceAdded) Id() message.Id {
    return msg.ID
}

func (msg *DeviceAdded) Serilize() interface{} {
    return map[string]*DeviceAdded{
        "DeviceAdded": msg,
    }
}

