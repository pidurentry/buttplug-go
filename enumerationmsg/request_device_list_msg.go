package enumerationmsg

// Auto generated file - do not directly edit!
import "github.com/pidurentry/buttplug-go/message"

func init() {
    message.Repository["RequestDeviceList"] = func() interface{} { return &RequestDeviceList{} }
}

func (msg *RequestDeviceList) Id() message.Id {
    return msg.ID
}

func (msg *RequestDeviceList) Serilize() interface{} {
    return map[string]*RequestDeviceList{
        "RequestDeviceList": msg,
    }
}

