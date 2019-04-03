package statusmsg

// Auto generated file - do not directly edit!
import "github.com/pidurentry/buttplug-go/message"

func init() {
    message.Repository["Error"] = func() interface{} { return &Error{} }
}

func (msg *Error) Id() message.Id {
    return msg.ID
}

func (msg *Error) Serilize() interface{} {
    return map[string]*Error{
        "Error": msg,
    }
}

