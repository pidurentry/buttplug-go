package statusmsg

// Auto generated file - do not directly edit!
import "github.com/pidurentry/buttplug-go"

func init() {
    buttplug.MessageRepository["Log"] = func() interface{} { return &Log{} }
}

func (msg *Log) Id() buttplug.MessageId {
    return msg.ID
}

func (msg *Log) Serilize() interface{} {
    return map[string]*Log{
        "Log": msg,
    }
}

