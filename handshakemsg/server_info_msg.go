package handshakemsg

// Auto generated file - do not directly edit!
import "github.com/pidurentry/buttplug-go"

func init() {
    buttplug.MessageRepository["ServerInfo"] = func() interface{} { return &ServerInfo{} }
}

func (msg *ServerInfo) Serilize() interface{} {
    return map[string]*ServerInfo{
        "ServerInfo": msg,
    }
}

