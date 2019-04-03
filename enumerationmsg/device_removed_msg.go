package enumerationmsg

// Auto generated file - do not directly edit!
import "github.com/pidurentry/buttplug-go/message"

func init() {
    message.Repository["DeviceRemoved"] = func() interface{} { return &DeviceRemoved{} }
}

func (msg *DeviceRemoved) Id() message.Id {
    return msg.ID
}

func (msg *DeviceRemoved) Serilize() interface{} {
    return map[string]*DeviceRemoved{
        "DeviceRemoved": msg,
    }
}

