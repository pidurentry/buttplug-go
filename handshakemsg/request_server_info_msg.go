package handshakemsg

// Auto generated file - do not directly edit!
import "github.com/pidurentry/buttplug-go/message"

func init() {
    message.Repository["RequestServerInfo"] = func() interface{} { return &RequestServerInfo{} }
}

func (msg *RequestServerInfo) Id() message.Id {
    return msg.ID
}

func (msg *RequestServerInfo) Serilize() interface{} {
    return map[string]*RequestServerInfo{
        "RequestServerInfo": msg,
    }
}

