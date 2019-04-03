package genericdevicemsg

// Auto generated file - do not directly edit!
import "github.com/pidurentry/buttplug-go/message"

func init() {
    message.Repository["LinearCmd"] = func() interface{} { return &LinearCmd{} }
}

func (msg *LinearCmd) Id() message.Id {
    return msg.ID
}

func (msg *LinearCmd) Serilize() interface{} {
    return map[string]*LinearCmd{
        "LinearCmd": msg,
    }
}

