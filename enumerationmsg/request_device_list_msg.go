package enumerationmsg

// Auto generated file - do not directly edit!
import "github.com/pidurentry/buttplug-go"

func init() {
    buttplug.MessageRepository["RequestDeviceList"] = func() interface{} { return &RequestDeviceList{} }
}

func (msg *RequestDeviceList) Id() buttplug.MessageId {
    return msg.ID
}

func (msg *RequestDeviceList) Serilize() interface{} {
    return map[string]*RequestDeviceList{
        "RequestDeviceList": msg,
    }
}

