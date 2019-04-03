package statusmsg

// Auto generated file - do not directly edit!
import "github.com/pidurentry/buttplug-go/message"

func init() {
    message.Repository["Log"] = func() interface{} { return &Log{} }
}

func (msg *Log) Id() message.Id {
    return msg.ID
}

func (msg *Log) Serilize() interface{} {
    return map[string]*Log{
        "Log": msg,
    }
}

