package enumerationmsg

// Auto generated file - do not directly edit!
import "github.com/pidurentry/buttplug-go"

func init() {
    buttplug.MessageRepository["DeviceRemoved"] = func() interface{} { return &DeviceRemoved{} }
}

func (msg *DeviceRemoved) Id() buttplug.MessageId {
    return msg.ID
}

func (msg *DeviceRemoved) Serilize() interface{} {
    return map[string]*DeviceRemoved{
        "DeviceRemoved": msg,
    }
}

