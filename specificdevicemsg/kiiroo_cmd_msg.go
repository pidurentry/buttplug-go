package specificdevicemsg

// Auto generated file - do not directly edit!
import "github.com/pidurentry/buttplug-go/message"

func init() {
    message.Repository["KiirooCmd"] = func() interface{} { return &KiirooCmd{} }
}

func (msg *KiirooCmd) Id() message.Id {
    return msg.ID
}

func (msg *KiirooCmd) Serilize() interface{} {
    return map[string]*KiirooCmd{
        "KiirooCmd": msg,
    }
}

