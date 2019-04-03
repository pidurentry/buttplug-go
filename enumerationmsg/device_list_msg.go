package enumerationmsg

// Auto generated file - do not directly edit!
import "github.com/pidurentry/buttplug-go/message"

func init() {
    message.Repository["DeviceList"] = func() interface{} { return &DeviceList{} }
}

func (msg *DeviceList) Id() message.Id {
    return msg.ID
}

func (msg *DeviceList) Serilize() interface{} {
    return map[string]*DeviceList{
        "DeviceList": msg,
    }
}

