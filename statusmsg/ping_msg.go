package statusmsg

// Auto generated file - do not directly edit!
import "github.com/pidurentry/buttplug-go/message"

func init() {
    message.Repository["Ping"] = func() interface{} { return &Ping{} }
}

func (msg *Ping) Id() message.Id {
    return msg.ID
}

func (msg *Ping) Serilize() interface{} {
    return map[string]*Ping{
        "Ping": msg,
    }
}

