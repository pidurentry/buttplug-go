package statusmsg

// Auto generated file - do not directly edit!
import "github.com/pidurentry/buttplug-go"

func init() {
    buttplug.MessageRepository["Ping"] = func() interface{} { return &Ping{} }
}

func (msg *Ping) Serilize() interface{} {
    return map[string]*Ping{
        "Ping": msg,
    }
}

