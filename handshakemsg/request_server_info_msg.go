package handshakemsg

// Auto generated file - do not directly edit!
import "github.com/pidurentry/buttplug-go"

func init() {
    buttplug.MessageRepository["RequestServerInfo"] = func() interface{} { return &RequestServerInfo{} }
}

func (msg *RequestServerInfo) Serilize() interface{} {
    return map[string]*RequestServerInfo{
        "RequestServerInfo": msg,
    }
}

