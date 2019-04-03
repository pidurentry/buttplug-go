package handshakemsg

// Auto generated file - do not directly edit!
import "github.com/pidurentry/buttplug-go/message"

func init() {
    message.Repository["ServerInfo"] = func() interface{} { return &ServerInfo{} }
}

func (msg *ServerInfo) Id() message.Id {
    return msg.ID
}

func (msg *ServerInfo) Serilize() interface{} {
    return map[string]*ServerInfo{
        "ServerInfo": msg,
    }
}

