package specificdevicemsg

// Auto generated file - do not directly edit!
import "github.com/pidurentry/buttplug-go/message"

func init() {
    message.Repository["LovenseCmd"] = func() interface{} { return &LovenseCmd{} }
}

func (msg *LovenseCmd) Id() message.Id {
    return msg.ID
}

func (msg *LovenseCmd) Serilize() interface{} {
    return map[string]*LovenseCmd{
        "LovenseCmd": msg,
    }
}

