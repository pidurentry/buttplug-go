package genericdevicemsg

// Auto generated file - do not directly edit!
import "github.com/pidurentry/buttplug-go/message"

func init() {
    message.Repository["StopAllDevices"] = func() interface{} { return &StopAllDevices{} }
}

func (msg *StopAllDevices) Id() message.Id {
    return msg.ID
}

func (msg *StopAllDevices) Serilize() interface{} {
    return map[string]*StopAllDevices{
        "StopAllDevices": msg,
    }
}

