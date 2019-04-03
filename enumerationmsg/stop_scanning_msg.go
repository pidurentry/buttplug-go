package enumerationmsg

// Auto generated file - do not directly edit!
import "github.com/pidurentry/buttplug-go/message"

func init() {
    message.Repository["StopScanning"] = func() interface{} { return &StopScanning{} }
}

func (msg *StopScanning) Id() message.Id {
    return msg.ID
}

func (msg *StopScanning) Serilize() interface{} {
    return map[string]*StopScanning{
        "StopScanning": msg,
    }
}

