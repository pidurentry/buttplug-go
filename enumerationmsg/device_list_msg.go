package enumerationmsg

// Auto generated file - do not directly edit!
import "github.com/pidurentry/buttplug-go"

func init() {
    buttplug.MessageRepository["DeviceList"] = func() interface{} { return &DeviceList{} }
}

func (msg *DeviceList) Id() buttplug.MessageId {
    return msg.ID
}

func (msg *DeviceList) Serilize() interface{} {
    return map[string]*DeviceList{
        "DeviceList": msg,
    }
}

