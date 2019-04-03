package genericdevicemsg

// Auto generated file - do not directly edit!
import "github.com/pidurentry/buttplug-go/message"

func init() {
    message.Repository["RotateCmd"] = func() interface{} { return &RotateCmd{} }
}

func (msg *RotateCmd) Id() message.Id {
    return msg.ID
}

func (msg *RotateCmd) Serilize() interface{} {
    return map[string]*RotateCmd{
        "RotateCmd": msg,
    }
}

