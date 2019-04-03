package genericdevicemsg

// Auto generated file - do not directly edit!
import "github.com/pidurentry/buttplug-go/message"

func init() {
    message.Repository["RawCmd"] = func() interface{} { return &RawCmd{} }
}

func (msg *RawCmd) Id() message.Id {
    return msg.ID
}

func (msg *RawCmd) Serilize() interface{} {
    return map[string]*RawCmd{
        "RawCmd": msg,
    }
}

