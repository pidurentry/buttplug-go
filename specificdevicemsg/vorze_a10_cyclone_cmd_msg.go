package specificdevicemsg

// Auto generated file - do not directly edit!
import "github.com/pidurentry/buttplug-go/message"

func init() {
    message.Repository["VorzeA10CycloneCmd"] = func() interface{} { return &VorzeA10CycloneCmd{} }
}

func (msg *VorzeA10CycloneCmd) Id() message.Id {
    return msg.ID
}

func (msg *VorzeA10CycloneCmd) Serilize() interface{} {
    return map[string]*VorzeA10CycloneCmd{
        "VorzeA10CycloneCmd": msg,
    }
}

