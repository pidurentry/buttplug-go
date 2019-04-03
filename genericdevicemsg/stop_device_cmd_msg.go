package genericdevicemsg

// Auto generated file - do not directly edit!
import "github.com/pidurentry/buttplug-go/message"

func init() {
    message.Repository["StopDeviceCmd"] = func() interface{} { return &StopDeviceCmd{} }
}

func (msg *StopDeviceCmd) Id() message.Id {
    return msg.ID
}

func (msg *StopDeviceCmd) Serilize() interface{} {
    return map[string]*StopDeviceCmd{
        "StopDeviceCmd": msg,
    }
}

