package statusmsg

// Auto generated file - do not directly edit!
import "github.com/pidurentry/buttplug-go"

func init() {
    buttplug.MessageRepository["Error"] = func() interface{} { return &Error{} }
}

func (msg *Error) Serilize() interface{} {
    return map[string]*Error{
        "Error": msg,
    }
}

