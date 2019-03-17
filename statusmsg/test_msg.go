package statusmsg

// Auto generated file - do not directly edit!
import "github.com/pidurentry/buttplug-go"

func init() {
    buttplug.MessageRepository["Test"] = func() interface{} { return &Test{} }
}

func (msg *Test) Id() buttplug.MessageId {
    return msg.ID
}

func (msg *Test) Serilize() interface{} {
    return map[string]*Test{
        "Test": msg,
    }
}

