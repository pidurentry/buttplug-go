package statusmsg

// Auto generated file - do not directly edit!
import "github.com/pidurentry/buttplug-go"

func init() {
    buttplug.MessageRepository["RequestLog"] = func() interface{} { return &RequestLog{} }
}

func (msg *RequestLog) Id() buttplug.MessageId {
    return msg.ID
}

func (msg *RequestLog) Serilize() interface{} {
    return map[string]*RequestLog{
        "RequestLog": msg,
    }
}

