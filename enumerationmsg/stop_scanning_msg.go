package enumerationmsg

// Auto generated file - do not directly edit!
import "github.com/pidurentry/buttplug-go"

func init() {
    buttplug.MessageRepository["StopScanning"] = func() interface{} { return &StopScanning{} }
}

func (msg *StopScanning) Id() buttplug.MessageId {
    return msg.ID
}

func (msg *StopScanning) Serilize() interface{} {
    return map[string]*StopScanning{
        "StopScanning": msg,
    }
}

