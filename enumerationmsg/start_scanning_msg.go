package enumerationmsg

// Auto generated file - do not directly edit!
import "github.com/pidurentry/buttplug-go/message"

func init() {
    message.Repository["StartScanning"] = func() interface{} { return &StartScanning{} }
}

func (msg *StartScanning) Id() message.Id {
    return msg.ID
}

func (msg *StartScanning) Serilize() interface{} {
    return map[string]*StartScanning{
        "StartScanning": msg,
    }
}

