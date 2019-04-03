package statusmsg

// Auto generated file - do not directly edit!
import "github.com/pidurentry/buttplug-go/message"

func init() {
    message.Repository["Ok"] = func() interface{} { return &Ok{} }
}

func (msg *Ok) Id() message.Id {
    return msg.ID
}

func (msg *Ok) Serilize() interface{} {
    return map[string]*Ok{
        "Ok": msg,
    }
}

