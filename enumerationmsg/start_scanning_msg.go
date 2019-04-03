package enumerationmsg

// Auto generated file - do not directly edit!
import "github.com/pidurentry/buttplug-go"

func init() {
    buttplug.MessageRepository["StartScanning"] = func() interface{} { return &StartScanning{} }
}

func (msg *StartScanning) Id() buttplug.MessageId {
    return msg.ID
}

func (msg *StartScanning) Serilize() interface{} {
    return map[string]*StartScanning{
        "StartScanning": msg,
    }
}

