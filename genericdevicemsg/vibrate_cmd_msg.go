package genericdevicemsg

// Auto generated file - do not directly edit!
import "github.com/pidurentry/buttplug-go/message"

func init() {
    message.Repository["VibrateCmd"] = func() interface{} { return &VibrateCmd{} }
}

func (msg *VibrateCmd) Id() message.Id {
    return msg.ID
}

func (msg *VibrateCmd) Serilize() interface{} {
    return map[string]*VibrateCmd{
        "VibrateCmd": msg,
    }
}

