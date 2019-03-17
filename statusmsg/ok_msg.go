package statusmsg

// Auto generated file - do not directly edit!
import "github.com/pidurentry/buttplug-go"

func init() {
    buttplug.MessageRepository["Ok"] = func() interface{} { return &Ok{} }
}

func (msg *Ok) Serilize() interface{} {
    return map[string]*Ok{
        "Ok": msg,
    }
}

