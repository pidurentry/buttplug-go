package statusmsg

// Auto generated file - do not directly edit!
import "github.com/pidurentry/buttplug-go/message"

func init() {
    message.Repository["RequestLog"] = func() interface{} { return &RequestLog{} }
}

func (msg *RequestLog) Id() message.Id {
    return msg.ID
}

func (msg *RequestLog) Serilize() interface{} {
    return map[string]*RequestLog{
        "RequestLog": msg,
    }
}

