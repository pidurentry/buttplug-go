package statusmsg

// Auto generated file - do not directly edit!
import "github.com/pidurentry/buttplug-go/message"

func init() {
    message.Repository["Test"] = func() interface{} { return &Test{} }
}

func (msg *Test) Id() message.Id {
    return msg.ID
}

func (msg *Test) Serilize() interface{} {
    return map[string]*Test{
        "Test": msg,
    }
}

